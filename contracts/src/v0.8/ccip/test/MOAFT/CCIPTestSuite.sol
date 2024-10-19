// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {CallWithExactGas} from "../../../shared/call/CallWithExactGas.sol";
import {Router} from "../../Router.sol";
import {Client} from "../../libraries/Client.sol";
import {EVM2EVMOffRamp} from "../../offRamp/EVM2EVMOffRamp.sol";
import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {EnumerableSet} from "../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

import {console2} from "forge-std/Console2.sol";
import {StdStorage, stdStorage} from "forge-std/StdStorage.sol";
import {Test} from "forge-std/Test.sol";

contract CCIPTestSuite is Test {
  using stdStorage for StdStorage;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.AddressSet;

  bytes32 internal constant TypeAndVersion1_5_OnRamp = keccak256("EVM2EVMOnRamp 1.5.0");

  uint256 internal constant TOKENS_TO_SEND = 1;
  uint16 internal constant MAX_RETURN_BYTES = 4 + 8 * 32;
  uint16 internal constant GAS_FOR_CALL_WITH_EXACT_GAS = 2500;
  uint256 internal constant TX_GAS_LIMIT = 1e6;

  uint256 internal s_pseudoRandomSeed = uint256(keccak256(abi.encodePacked(block.timestamp, block.number)));

  Router internal immutable i_router;

  struct RemoteChainConfig {
    EVM2EVMOnRamp OldOnRamp;
    EVM2EVMOnRamp NewOnRamp;
    EVM2EVMOffRamp OldOffRamp;
    EVM2EVMOffRamp NewOffRamp;
    address[] tokens;
    address[] oldSuccessfulTokens;
  }

  EnumerableSet.UintSet internal s_remoteChainSelectors;
  mapping(uint64 remoteChainSelector => RemoteChainConfig) public s_remoteChainConfigs;

  EnumerableSet.AddressSet internal s_failedTokensInitially;
  EnumerableSet.AddressSet internal s_failedTokensAfterMigration;

  constructor(
    address router
  ) {
    i_router = Router(router);
    _loadDestChainInfo();

    vm.deal(address(this), 1e18 ether);
  }

  function sendAllTokens(
    bool onlyPreviousSuccess
  ) external {
    uint256[] memory remoteChains = s_remoteChainSelectors.values();

    uint256 totalSuccessfulTokens = 0;
    uint256 totalFailedTokens = 0;
    for (uint256 i = 0; i < remoteChains.length; ++i) {
      uint256 successfulTokens = 0;
      uint64 remoteChainSelector = uint64(remoteChains[i]);
      console2.log("Sending tokens to chain: ", remoteChainSelector);
      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[remoteChainSelector];
      address[] memory tokens = onlyPreviousSuccess ? remoteChainConfig.oldSuccessfulTokens : remoteChainConfig.tokens;
      for (uint256 j = 0; j < tokens.length; ++j) {
        address token = tokens[j];

        bytes memory payload = abi.encodeWithSelector(this.SendTokenMsg.selector, token, remoteChainSelector);

        (bool success, bytes memory retData,) = CallWithExactGas._callWithExactGasSafeReturnData(
          payload, address(this), TX_GAS_LIMIT, GAS_FOR_CALL_WITH_EXACT_GAS, MAX_RETURN_BYTES
        );

        if (success) {
          console2.log("[SUCCESS] token", token);
          if (!onlyPreviousSuccess) {
            s_remoteChainConfigs[remoteChainSelector].oldSuccessfulTokens.push(token);
          }
          successfulTokens++;
        } else {
          if (!onlyPreviousSuccess) {
            s_failedTokensInitially.add(token);
          } else {
            s_failedTokensAfterMigration.add(token);
          }
          console2.log("[FAILURE] token", token);
          console2.logBytes(retData);
        }
      }
      console2.log("Tokens sent: success", successfulTokens, "failed", tokens.length - successfulTokens);
      console2.log("");

      totalSuccessfulTokens += successfulTokens;
      totalFailedTokens += tokens.length - successfulTokens;
    }
    console2.log("--------------------------------------- +");
    console2.log("Total sent: success", totalSuccessfulTokens, "failed", totalFailedTokens);
    console2.log("");
    console2.log("Failed tokens initially:");
    for (uint256 i = 0; i < s_failedTokensInitially.length(); ++i) {
      console2.logAddress(s_failedTokensInitially.at(i));
    }
    console2.log("");
    console2.log("Newly failing tokens after migration:");
    for (uint256 i = 0; s_failedTokensAfterMigration.length() > 0;) {
      console2.logAddress(s_failedTokensAfterMigration.at(i));
      s_failedTokensAfterMigration.remove(s_failedTokensAfterMigration.at(i));
    }
  }

  function SendTokenMsg(address token, uint64 destChainSelector) public {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: TOKENS_TO_SEND});

    deal(token, address(this), TOKENS_TO_SEND * 10);

    IERC20(token).approve(address(i_router), TOKENS_TO_SEND);
    i_router.ccipSend{value: 1 ether}(
      destChainSelector,
      Client.EVM2AnyMessage({
        receiver: abi.encode(_getRandomAddress()),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: address(0),
        extraArgs: ""
      })
    );
  }

  function _loadDestChainInfo() internal {
    Router.OffRamp[] memory offRamps = i_router.getOffRamps();
    for (uint256 i = 0; i < offRamps.length; ++i) {
      Router.OffRamp memory offRamp = offRamps[i];

      EVM2EVMOnRamp currentOnRamp = EVM2EVMOnRamp(i_router.getOnRamp(offRamp.sourceChainSelector));
      // Skip 1.5 lanes for now
      if (keccak256(bytes(currentOnRamp.typeAndVersion())) == TypeAndVersion1_5_OnRamp) {
        continue;
      }

      s_remoteChainSelectors.add(offRamp.sourceChainSelector);

      s_remoteChainConfigs[offRamp.sourceChainSelector] = RemoteChainConfig({
        OldOnRamp: EVM2EVMOnRamp(currentOnRamp),
        NewOnRamp: EVM2EVMOnRamp(address(0)),
        OldOffRamp: EVM2EVMOffRamp(offRamp.offRamp),
        NewOffRamp: EVM2EVMOffRamp(address(0)),
        tokens: i_router.getSupportedTokens(offRamp.sourceChainSelector),
        oldSuccessfulTokens: new address[](0)
      });
    }
  }

  function _getRandomAddress() internal returns (address) {
    return address(uint160(++s_pseudoRandomSeed));
  }
}
