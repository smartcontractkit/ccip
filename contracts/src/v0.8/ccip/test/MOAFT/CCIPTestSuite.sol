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

  bytes32 internal constant TypeAndVersion1_5_OnRamp = keccak256("EVM2EVMOnRamp 1.5.0");

  uint256 internal constant TOKENS_TO_SEND = 10;
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

  constructor(
    address router
  ) {
    i_router = Router(router);
    _loadDestChainInfo();

    vm.deal(address(this), 1e18 ether);
  }

  function sendAllTokens() external {
    uint256[] memory remoteChains = s_remoteChainSelectors.values();

    for (uint256 i = 0; i < remoteChains.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChains[i]);
      RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[remoteChainSelector];
      for (uint256 j = 0; j < remoteChainConfig.tokens.length; ++j) {
        address token = remoteChainConfig.tokens[j];

        bytes memory payload = abi.encodeWithSelector(this.SendTokenMsg.selector, token, remoteChainSelector);

        (bool success, bytes memory retData,) = CallWithExactGas._callWithExactGasSafeReturnData(
          payload, address(this), TX_GAS_LIMIT, GAS_FOR_CALL_WITH_EXACT_GAS, MAX_RETURN_BYTES
        );

        if (success) {
          console2.log("SUCCESS: sent token", token);
          s_remoteChainConfigs[remoteChainSelector].oldSuccessfulTokens.push(token);
        } else {
          console2.log("FAILURE: sent token", token);
          console2.logBytes(retData);
        }
      }
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
        tokens: i_router.getSupportedTokens(offRamp.sourceChainSelector)
      });
    }
  }

  function _getRandomAddress() internal returns (address) {
    return address(uint160(++s_pseudoRandomSeed));
  }
}
