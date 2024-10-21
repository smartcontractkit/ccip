// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {CallWithExactGas} from "../../../shared/call/CallWithExactGas.sol";
import {Router} from "../../Router.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {EVM2EVMOffRamp} from "../../offRamp/EVM2EVMOffRamp.sol";
import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";

import {EnumerableSet} from "../../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

import {console2} from "forge-std/Console2.sol";
import {StdStorage, stdStorage} from "forge-std/StdStorage.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

library Constants {
  function _resolveChainSelector(
    uint64 chainSelector
  ) internal pure returns (string memory) {
    if (chainSelector == 3478487238524512106) {
      return "Arbitrum Sepolia";
    } else if (chainSelector == 8871595565390010547) {
      return "Gnosis Chiado";
    } else if (chainSelector == 16281711391670634445) {
      return "Polygon Amoy";
    } else if (chainSelector == 13264668187771770619) {
      return "BNB Testnet";
    } else if (chainSelector == 10344971235874465080) {
      return "Base Testnet";
    } else if (chainSelector == 829525985033418733) {
      return "Mode Sepolia";
    } else if (chainSelector == 5224473277236331295) {
      return "Optimism Sepolia";
    } else if (chainSelector == 16015286601757825753) {
      return "Sepolia";
    } else if (chainSelector == 6898391096552792247) {
      return "ZKSync Testnet";
    } else if (chainSelector == 14767482510784806043) {
      return "Avax Fuji";
    }

    return "Unknown";
  }
}

contract CCIPTestSuite is Test {
  using stdStorage for StdStorage;
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.AddressSet;

  bytes32 internal constant TypeAndVersion1_5_OnRamp = keccak256("EVM2EVMOnRamp 1.5.0");
  bytes32 internal constant TypeAndVersion1_5_OffRamp = keccak256("EVM2EVMOffRamp 1.5.0");

  uint256 internal constant TOKENS_TO_SEND = 1;
  uint16 internal constant MAX_RETURN_BYTES = 4 + 8 * 32;
  uint16 internal constant GAS_FOR_CALL_WITH_EXACT_GAS = 2500;
  uint256 internal constant TX_GAS_LIMIT = 1e6;

  uint256 internal s_pseudoRandomSeed = uint256(keccak256(abi.encodePacked(block.timestamp, block.number)));

  Router internal immutable i_router;
  bool internal immutable i_fullLogging;

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
  EnumerableSet.AddressSet internal s_tokenSupportDropped;

  mapping(address token => string) public s_tokenNames;

  constructor(address router, bool fullLogging) {
    i_router = Router(router);
    i_fullLogging = fullLogging;
    _loadDestChainInfo();

    vm.deal(address(this), 1e18 ether);
  }

  function sendTokensSingleLane(
    uint64 remoteChainSelector
  ) external returns (Internal.EVM2EVMMessage[] memory msgs) {
    vm.recordLogs();

    (uint256 successful,) = _sendTokenToChain(false, remoteChainSelector);

    Vm.Log[] memory logs = vm.getRecordedLogs();
    Internal.EVM2EVMMessage[] memory messages = new Internal.EVM2EVMMessage[](successful);
    uint256 logsFound = 0;
    for (uint256 i = 0; i < logs.length; ++i) {
      if (logs[i].topics[0] == EVM2EVMOnRamp.CCIPSendRequested.selector) {
        messages[logsFound] = abi.decode(logs[i].data, (Internal.EVM2EVMMessage));
        logsFound++;
      }
    }
    return messages;
  }

  function sendAllTokens(
    bool postMigration
  ) external {
    uint256[] memory remoteChains = s_remoteChainSelectors.values();

    uint256 totalSuccessfulTokens = 0;
    uint256 totalFailedTokens = 0;
    for (uint256 i = 0; i < remoteChains.length; ++i) {
      uint64 remoteChainSelector = uint64(remoteChains[i]);
      (uint256 numberOfSuccesses, uint256 failed) = _sendTokenToChain(postMigration, remoteChainSelector);
      totalSuccessfulTokens += numberOfSuccesses;
      totalFailedTokens += failed;
    }

    if (i_fullLogging) {
      console2.log("--------------------------------------- +");
    }
    console2.log("Total sent: success", totalSuccessfulTokens, "failed", totalFailedTokens);
    console2.log("");

    if (!postMigration) {
      return;
    }

    console2.log("Failed tokens initially:");
    for (uint256 i = 0; i < s_failedTokensInitially.length(); ++i) {
      console2.log(unicode"❓", s_failedTokensInitially.at(i), s_tokenNames[s_failedTokensInitially.at(i)]);
    }

    console2.log("");
    console2.log("Tokens dropped with the migration:");
    for (uint256 i = 0; i < s_tokenSupportDropped.length(); ++i) {
      console2.log(unicode"🟠", s_tokenSupportDropped.at(i), s_tokenNames[s_tokenSupportDropped.at(i)]);
    }

    console2.log("");
    console2.log("Newly failing tokens after migration:");
    for (uint256 i = 0; s_failedTokensAfterMigration.length() > 0;) {
      console2.log(unicode"❌", s_failedTokensAfterMigration.at(i), s_tokenNames[s_failedTokensAfterMigration.at(i)]);
      s_failedTokensAfterMigration.remove(s_failedTokensAfterMigration.at(i));
    }
  }

  function _sendTokenToChain(
    bool postMigration,
    uint64 remoteChainSelector
  ) internal returns (uint256 successful, uint256 failed) {
    uint256 successfulTokens = 0;

    if (i_fullLogging) {
      console2.log(
        "Sending tokens to chain: ", remoteChainSelector, Constants._resolveChainSelector(remoteChainSelector)
      );
    }
    RemoteChainConfig storage remoteChainConfig = s_remoteChainConfigs[remoteChainSelector];
    address[] memory tokens = postMigration ? remoteChainConfig.oldSuccessfulTokens : remoteChainConfig.tokens;
    for (uint256 j = 0; j < tokens.length; ++j) {
      address token = tokens[j];

      (bool success, bytes memory retData,) = CallWithExactGas._callWithExactGasSafeReturnData(
        abi.encodeWithSelector(this.SendTokenMsg.selector, token, remoteChainSelector),
        address(this),
        TX_GAS_LIMIT,
        GAS_FOR_CALL_WITH_EXACT_GAS,
        MAX_RETURN_BYTES
      );

      if (success) {
        if (i_fullLogging) {
          console2.log(unicode"✅", token, s_tokenNames[token]);
        }
        if (!postMigration) {
          s_remoteChainConfigs[remoteChainSelector].oldSuccessfulTokens.push(token);
        }
        successfulTokens++;
        continue;
      }

      bool tokenSupportDropped =
        keccak256(retData) == keccak256(abi.encodeWithSelector(EVM2EVMOnRamp.UnsupportedToken.selector, token));

      if (!postMigration) {
        s_failedTokensInitially.add(token);
      } else {
        if (tokenSupportDropped) {
          s_tokenSupportDropped.add(token);
        } else {
          s_failedTokensAfterMigration.add(token);
        }
      }

      if (i_fullLogging) {
        if (tokenSupportDropped) {
          console2.log(unicode"🟠 DROPPED SUPPORT", token, s_tokenNames[token]);
        } else {
          console2.log(unicode"❌ broken", token, s_tokenNames[token]);
          console2.logBytes(retData);
        }
      }
    }
    if (i_fullLogging) {
      console2.log("Tokens sent: success", successfulTokens, "failed", tokens.length - successfulTokens);
      console2.log("");
    }

    return (successfulTokens, tokens.length - successfulTokens);
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

  function ExecuteMsgs(
    Internal.EVM2EVMMessage[] memory messages
  ) public {
    if (messages.length == 0) {
      return;
    }
    _loadLatestOffRampData();

    uint64 sourceChainSelector = messages[0].sourceChainSelector;
    EVM2EVMOffRamp offRamp = s_remoteChainConfigs[sourceChainSelector].NewOffRamp;

    vm.startPrank(address(offRamp));

    uint32[] memory gasOverrides = new uint32[](1);
    gasOverrides[0] = 100_000;
    //    uint32[] memory gasOverrides = new uint32[](0);
    uint256 succeeded = 0;

    for (uint256 i = 0; i < messages.length; ++i) {
      Internal.EVM2EVMMessage memory message = messages[i];
      bytes memory destTokenAddressBytes =
        abi.decode(message.sourceTokenData[0], (Internal.SourceTokenData)).destTokenAddress;
      address destTokenAddress = abi.decode(destTokenAddressBytes, (address));
      try offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), gasOverrides) {
        console2.log(
          unicode"✅ Executed message with source token", message.tokenAmounts[0].token, s_tokenNames[destTokenAddress]
        );
        succeeded++;
      } catch (bytes memory reason) {
        console2.log(
          unicode"❌ Failed to execute message with token",
          message.tokenAmounts[0].token,
          s_tokenNames[destTokenAddress]
        );
        console2.logBytes(reason);
      }
    }

    console2.log("Executed", succeeded, "out of", messages.length);
  }

  function _loadLatestOffRampData() internal {
    Router.OffRamp[] memory offRamps = i_router.getOffRamps();
    for (uint256 i = 0; i < offRamps.length; ++i) {
      Router.OffRamp memory offRamp = offRamps[i];
      EVM2EVMOffRamp currentOffRamp = EVM2EVMOffRamp(offRamp.offRamp);
      if (keccak256(bytes(currentOffRamp.typeAndVersion())) != TypeAndVersion1_5_OffRamp) {
        continue;
      }

      s_remoteChainConfigs[offRamp.sourceChainSelector].NewOffRamp = currentOffRamp;
    }
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

      address[] memory tokens = i_router.getSupportedTokens(offRamp.sourceChainSelector);
      for (uint256 j = 0; j < tokens.length; ++j) {
        address token = tokens[j];
        s_tokenNames[token] = IERC20(token).name();
      }

      s_remoteChainConfigs[offRamp.sourceChainSelector] = RemoteChainConfig({
        OldOnRamp: EVM2EVMOnRamp(currentOnRamp),
        NewOnRamp: EVM2EVMOnRamp(address(0)),
        OldOffRamp: EVM2EVMOffRamp(offRamp.offRamp),
        NewOffRamp: EVM2EVMOffRamp(address(0)),
        tokens: tokens,
        oldSuccessfulTokens: new address[](0)
      });
    }
  }

  function _getRandomAddress() internal returns (address) {
    return address(uint160(++s_pseudoRandomSeed));
  }
}
