// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Vm} from "forge-std/Vm.sol";

import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {IPool} from "../../interfaces/pools/IPool.sol";

import {CallWithExactGas} from "../../../shared/call/CallWithExactGas.sol";

import {ARM} from "../../ARM.sol";
import {AggregateRateLimiter} from "../../AggregateRateLimiter.sol";
import {Router} from "../../Router.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {EVM2EVMMultiOffRamp} from "../../offRamp/EVM2EVMMultiOffRamp.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {EVM2EVMMultiOffRampHelper} from "../helpers/EVM2EVMMultiOffRampHelper.sol";
import {MaybeRevertingBurnMintTokenPool} from "../helpers/MaybeRevertingBurnMintTokenPool.sol";
import {ConformingReceiver} from "../helpers/receivers/ConformingReceiver.sol";
import {MaybeRevertMessageReceiver} from "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import {MaybeRevertMessageReceiverNo165} from "../helpers/receivers/MaybeRevertMessageReceiverNo165.sol";
import {ReentrancyAbuser} from "../helpers/receivers/ReentrancyAbuser.sol";
import {MockCommitStore} from "../mocks/MockCommitStore.sol";
import {OCR2Base} from "../ocr/OCR2Base.t.sol";
import {OCR2BaseNoChecks} from "../ocr/OCR2BaseNoChecks.t.sol";
import {EVM2EVMMultiOffRampSetup} from "./EVM2EVMMultiOffRampSetup.t.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

// TODO: re-add tests:
//       - constructor
//       - ccipReceive
//       - execute
//       - execute_upgrade
//       - executeSingleMessage
//       - report
//       - manuallyExecute
//       - getExecutionState
//       - trialExecute
//       - getAllRateLimitTokens
//       - updateRateLimitTokens

contract EVM2EVMMultiOffRamp_setDynamicConfig is EVM2EVMMultiOffRampSetup {
  // OffRamp event
  event ConfigSet(EVM2EVMMultiOffRamp.StaticConfig staticConfig, EVM2EVMMultiOffRamp.DynamicConfig dynamicConfig);

  function test_SetDynamicConfig_Success() public {
    EVM2EVMMultiOffRamp.StaticConfig memory staticConfig = s_offRamp.getStaticConfig();
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig =
      generateDynamicMultiOffRampConfig(USER_3, address(s_priceRegistry));
    bytes memory onchainConfig = abi.encode(dynamicConfig);

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    vm.expectEmit();
    uint32 configCount = 1;
    emit ConfigSet(
      uint32(block.number),
      getBasicConfigDigest(address(s_offRamp), s_f, configCount, onchainConfig),
      configCount + 1,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, onchainConfig, s_offchainConfigVersion, abi.encode("")
    );

    EVM2EVMMultiOffRamp.DynamicConfig memory newConfig = s_offRamp.getDynamicConfig();
    _assertSameConfig(dynamicConfig, newConfig);
  }

  function test_NonOwner_Revert() public {
    vm.startPrank(STRANGER);
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig =
      generateDynamicMultiOffRampConfig(USER_3, address(s_priceRegistry));

    vm.expectRevert("Only callable by owner");

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, abi.encode(dynamicConfig), s_offchainConfigVersion, abi.encode("")
    );
  }

  function test_RouterZeroAddress_Revert() public {
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig =
      generateDynamicMultiOffRampConfig(ZERO_ADDRESS, ZERO_ADDRESS);

    vm.expectRevert(EVM2EVMMultiOffRamp.ZeroAddressNotAllowed.selector);

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, abi.encode(dynamicConfig), s_offchainConfigVersion, abi.encode("")
    );
  }
}

contract EVM2EVMMultiOffRamp_metadataHash is EVM2EVMMultiOffRampSetup {
  function test_MetadataHash_Success() public view {
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
      )
    );
  }

  function test_MetadataHashChangesOnSourceChain_Success() public view {
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR + 1, ON_RAMP_ADDRESS);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR + 1, DEST_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
      )
    );
    assertTrue(h != s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS));
  }

  function test_MetadataHashChangesOnOnRampAddress_Success() public view {
    address mockOnRampAddress = address(uint160(ON_RAMP_ADDRESS) + 1);
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, mockOnRampAddress);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, mockOnRampAddress)
      )
    );
    assertTrue(h != s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS));
  }
}

contract EVM2EVMMultiOffRamp__releaseOrMintTokens is EVM2EVMMultiOffRampSetup {
  EVM2EVMMultiOffRamp.Any2EVMMessageRoute internal MESSAGE_ROUTE;

  function setUp() public virtual override {
    super.setUp();
    MESSAGE_ROUTE = EVM2EVMMultiOffRamp.Any2EVMMessageRoute({
      sender: abi.encode(OWNER),
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      receiver: OWNER
    });
  }

  function test_releaseOrMintTokens_Success() public {
    Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
    IERC20 dstToken1 = IERC20(s_destFeeToken);
    uint256 startingBalance = dstToken1.balanceOf(OWNER);
    uint256 amount1 = 100;
    srcTokenAmounts[0].amount = amount1;

    bytes[] memory offchainTokenData = new bytes[](srcTokenAmounts.length);
    offchainTokenData[0] = abi.encode(0x12345678);

    bytes[] memory sourceTokenData = _getDefaultSourceTokenData(srcTokenAmounts);

    vm.expectCall(
      s_destPoolBySourceToken[srcTokenAmounts[0].token],
      abi.encodeWithSelector(
        LockReleaseTokenPool.releaseOrMint.selector,
        MESSAGE_ROUTE.sender,
        MESSAGE_ROUTE.receiver,
        srcTokenAmounts[0].amount,
        MESSAGE_ROUTE.sourceChainSelector,
        abi.decode(sourceTokenData[0], (IPool.SourceTokenData)),
        offchainTokenData[0]
      )
    );

    s_offRamp.releaseOrMintTokens(srcTokenAmounts, MESSAGE_ROUTE, sourceTokenData, offchainTokenData);

    assertEq(startingBalance + amount1, dstToken1.balanceOf(MESSAGE_ROUTE.receiver));
  }

  // TODO: re-add after ARL changes
  // function test_OverValueWithARLOff_Success() public {
  //   // Set a high price to trip the ARL
  //   uint224 tokenPrice = 3 ** 128;
  //   Internal.PriceUpdates memory priceUpdates = getSingleTokenPriceUpdateStruct(s_destFeeToken, tokenPrice);
  //   s_priceRegistry.updatePrices(priceUpdates);

  //   Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
  //   IERC20 dstToken1 = IERC20(s_destFeeToken);
  //   uint256 amount1 = 100;
  //   srcTokenAmounts[0].amount = amount1;

  //   bytes memory originalSender = abi.encode(OWNER);

  //   bytes[] memory offchainTokenData = new bytes[](srcTokenAmounts.length);
  //   offchainTokenData[0] = abi.encode(0x12345678);

  //   bytes[] memory sourceTokenData = _getDefaultSourceTokenData(srcTokenAmounts);

  //   vm.expectRevert(
  //     abi.encodeWithSelector(
  //       RateLimiter.AggregateValueMaxCapacityExceeded.selector,
  //       getInboundRateLimiterConfig().capacity,
  //       (amount1 * tokenPrice) / 1e18
  //     )
  //   );

  //   // // Expect to fail from ARL
  //   s_offRamp.releaseOrMintTokens(srcTokenAmounts, originalSender, OWNER, sourceTokenData, offchainTokenData);

  //   // Configure ARL off for token
  //   EVM2EVMMultiOffRamp.RateLimitToken[] memory removes = new EVM2EVMMultiOffRamp.RateLimitToken[](1);
  //   removes[0] = EVM2EVMMultiOffRamp.RateLimitToken({sourceToken: s_sourceFeeToken, destToken: s_destFeeToken});
  //   s_offRamp.updateRateLimitTokens(removes, new EVM2EVMMultiOffRamp.RateLimitToken[](0));

  //   // Expect the call now succeeds
  //   s_offRamp.releaseOrMintTokens(srcTokenAmounts, originalSender, OWNER, sourceTokenData, offchainTokenData);
  // }

  // Revert

  function test_TokenHandlingError_Reverts() public {
    Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();

    bytes memory unknownError = bytes("unknown error");
    s_maybeRevertingPool.setShouldRevert(unknownError);

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOffRamp.TokenHandlingError.selector, unknownError));

    s_offRamp.releaseOrMintTokens(
      srcTokenAmounts, MESSAGE_ROUTE, _getDefaultSourceTokenData(srcTokenAmounts), new bytes[](srcTokenAmounts.length)
    );
  }

  // TODO: re-add after ARL changes
  // function test_RateLimitErrors_Reverts() public {
  //   Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();

  //   bytes[] memory rateLimitErrors = new bytes[](5);
  //   rateLimitErrors[0] = abi.encodeWithSelector(RateLimiter.BucketOverfilled.selector);
  //   rateLimitErrors[1] =
  //     abi.encodeWithSelector(RateLimiter.AggregateValueMaxCapacityExceeded.selector, uint256(100), uint256(1000));
  //   rateLimitErrors[2] =
  //     abi.encodeWithSelector(RateLimiter.AggregateValueRateLimitReached.selector, uint256(42), 1, s_sourceTokens[0]);
  //   rateLimitErrors[3] = abi.encodeWithSelector(
  //     RateLimiter.TokenMaxCapacityExceeded.selector, uint256(100), uint256(1000), s_sourceTokens[0]
  //   );
  //   rateLimitErrors[4] =
  //     abi.encodeWithSelector(RateLimiter.TokenRateLimitReached.selector, uint256(42), 1, s_sourceTokens[0]);

  //   for (uint256 i = 0; i < rateLimitErrors.length; ++i) {
  //     s_maybeRevertingPool.setShouldRevert(rateLimitErrors[i]);

  //     vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOffRamp.TokenHandlingError.selector, rateLimitErrors[i]));

  //     s_offRamp.releaseOrMintTokens(
  //       srcTokenAmounts,
  //       abi.encode(OWNER),
  //       OWNER,
  //       _getDefaultSourceTokenData(srcTokenAmounts),
  //       new bytes[](srcTokenAmounts.length)
  //     );
  //   }
  // }

  function test__releaseOrMintTokens_PoolIsNotAPool_Reverts() public {
    address fakePoolAddress = makeAddr("Doesn't exist");

    bytes[] memory sourceTokenData = new bytes[](1);
    sourceTokenData[0] = abi.encode(
      IPool.SourceTokenData({
        sourcePoolAddress: abi.encode(fakePoolAddress),
        destPoolAddress: abi.encode(s_offRamp),
        extraData: ""
      })
    );

    MESSAGE_ROUTE.sender = abi.encode(makeAddr("original_sender"));
    vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOffRamp.TokenHandlingError.selector, bytes("")));
    s_offRamp.releaseOrMintTokens(new Client.EVMTokenAmount[](1), MESSAGE_ROUTE, sourceTokenData, new bytes[](1));
  }

  function test__releaseOrMintTokens_PoolIsNotAContract_Reverts() public {
    address fakePoolAddress = makeAddr("Doesn't exist");

    bytes[] memory sourceTokenData = new bytes[](1);
    sourceTokenData[0] = abi.encode(
      IPool.SourceTokenData({
        sourcePoolAddress: abi.encode(fakePoolAddress),
        destPoolAddress: abi.encode(fakePoolAddress),
        extraData: ""
      })
    );

    MESSAGE_ROUTE.sender = abi.encode(makeAddr("original_sender"));
    vm.expectRevert(abi.encodeWithSelector(EVM2EVMMultiOffRamp.InvalidAddress.selector, abi.encode(fakePoolAddress)));
    s_offRamp.releaseOrMintTokens(new Client.EVMTokenAmount[](1), MESSAGE_ROUTE, sourceTokenData, new bytes[](1));
  }

  function test_PriceNotFoundForToken_Reverts() public {
    // Set token price to 0
    s_priceRegistry.updatePrices(getSingleTokenPriceUpdateStruct(s_destFeeToken, 0));

    Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
    IERC20 dstToken1 = IERC20(s_destFeeToken);
    uint256 amount1 = 100;
    srcTokenAmounts[0].amount = amount1;

    bytes[] memory offchainTokenData = new bytes[](srcTokenAmounts.length);
    offchainTokenData[0] = abi.encode(0x12345678);

    bytes[] memory sourceTokenData = _getDefaultSourceTokenData(srcTokenAmounts);

    vm.expectRevert(abi.encodeWithSelector(AggregateRateLimiter.PriceNotFoundForToken.selector, s_destFeeToken));

    s_offRamp.releaseOrMintTokens(srcTokenAmounts, MESSAGE_ROUTE, sourceTokenData, offchainTokenData);
  }

  /// forge-config: default.fuzz.runs = 32
  /// forge-config: ccip.fuzz.runs = 10024
  function test_fuzz__releaseOrMintTokens_AnyRevertIsCaught_Success(uint256 destPool) public {
    // TODO handle 447301751254033913445893214690834296930546521452, which is 4E59B44847B379578588920CA78FBF26C0B4956C
    // which triggers some Create2Deployer and causes it to fail
    vm.assume(destPool != 447301751254033913445893214690834296930546521452);
    bytes memory unusedVar = abi.encode(makeAddr("unused"));
    // Uint256 gives a good range of values to test, both inside and outside of the eth address space.
    bytes memory destPoolAddress = abi.encode(destPool);
    bytes[] memory sourceTokenData = new bytes[](1);
    sourceTokenData[0] = abi.encode(
      IPool.SourceTokenData({sourcePoolAddress: unusedVar, destPoolAddress: destPoolAddress, extraData: unusedVar})
    );

    try s_offRamp.releaseOrMintTokens(new Client.EVMTokenAmount[](1), MESSAGE_ROUTE, sourceTokenData, new bytes[](1)) {}
    catch (bytes memory reason) {
      // Any revert should be a TokenHandlingError or InvalidAddress as those are caught by the offramp
      assertTrue(
        bytes4(reason) == EVM2EVMMultiOffRamp.TokenHandlingError.selector
          || bytes4(reason) == EVM2EVMMultiOffRamp.InvalidAddress.selector,
        "Expected TokenHandlingError or InvalidAddress"
      );

      if (destPool > type(uint160).max) {
        // If the destPool is not a valid eth address, the inner error should be PoolDoesNotExist
        assertEq(reason, abi.encodeWithSelector(EVM2EVMMultiOffRamp.InvalidAddress.selector, destPoolAddress));
      }
    }
  }
}

contract EVM2EVMMultiOffRamp_applySoureConfigUpdates is EVM2EVMMultiOffRampSetup {
  event SourceChainSelectorAdded(uint64 sourceChainSelector);
  event SourceChainConfigSet(uint64 indexed sourceChainSelector, EVM2EVMMultiOffRamp.SourceChainConfig sourceConfig);

  uint64 SOURCE_CHAIN_SELECTOR_1 = 16015286601757825753;

  function test_ApplyZeroUpdates_Success() public {
    uint64[] memory sourceChainSelectors = new uint64[](0);
    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](0);

    vm.recordLogs();
    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);

    // No logs emitted
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);

    assertEq(s_offRamp.getSourceChainSelectors().length, 0);
  }

  function test_AddNewChain_Success() public {
    uint64[] memory sourceChainSelectors = new uint64[](1);
    sourceChainSelectors[0] = SOURCE_CHAIN_SELECTOR_1;

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](1);
    sourceChainConfigs[0] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: ON_RAMP_ADDRESS,
      metadataHash: ""
    });

    EVM2EVMMultiOffRamp.SourceChainConfig memory expectedSourceChainConfig = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: ON_RAMP_ADDRESS,
      metadataHash: s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR_1, ON_RAMP_ADDRESS)
    });

    vm.expectEmit();
    emit SourceChainSelectorAdded(SOURCE_CHAIN_SELECTOR_1);

    vm.expectEmit();
    emit SourceChainConfigSet(SOURCE_CHAIN_SELECTOR_1, expectedSourceChainConfig);

    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);

    _assertSourceChainConfigEquality(s_offRamp.getSourceChainConfig(SOURCE_CHAIN_SELECTOR_1), expectedSourceChainConfig);

    uint64[] memory resultSourceChainSelectors = s_offRamp.getSourceChainSelectors();
    assertEq(resultSourceChainSelectors.length, 1);
    assertEq(resultSourceChainSelectors[0], SOURCE_CHAIN_SELECTOR_1);
  }

  function test_ReplaceExistingChain_Success() public {
    uint64[] memory sourceChainSelectors = new uint64[](1);
    sourceChainSelectors[0] = SOURCE_CHAIN_SELECTOR_1;

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](1);
    sourceChainConfigs[0] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: ON_RAMP_ADDRESS,
      metadataHash: ""
    });

    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);

    sourceChainConfigs[0].onRamp = address(uint160(ON_RAMP_ADDRESS) + 1);
    EVM2EVMMultiOffRamp.SourceChainConfig memory expectedSourceChainConfig = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: sourceChainConfigs[0].onRamp,
      metadataHash: s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR_1, sourceChainConfigs[0].onRamp)
    });

    vm.expectEmit();
    emit SourceChainConfigSet(SOURCE_CHAIN_SELECTOR_1, expectedSourceChainConfig);

    vm.recordLogs();
    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);

    // No log emitted for chain selector added (only for setting the config)
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 1);

    _assertSourceChainConfigEquality(s_offRamp.getSourceChainConfig(SOURCE_CHAIN_SELECTOR_1), expectedSourceChainConfig);

    uint64[] memory resultSourceChainSelectors = s_offRamp.getSourceChainSelectors();
    assertEq(resultSourceChainSelectors.length, 1);
    assertEq(resultSourceChainSelectors[0], SOURCE_CHAIN_SELECTOR_1);
  }

  function test_AddMultipleChains_Success() public {
    uint64[] memory sourceChainSelectors = new uint64[](3);
    sourceChainSelectors[0] = SOURCE_CHAIN_SELECTOR_1;
    sourceChainSelectors[1] = SOURCE_CHAIN_SELECTOR_1 + 1;
    sourceChainSelectors[2] = SOURCE_CHAIN_SELECTOR_1 + 2;

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](3);
    sourceChainConfigs[0] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: ON_RAMP_ADDRESS,
      metadataHash: ""
    });
    sourceChainConfigs[1] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: false,
      prevOffRamp: address(999),
      onRamp: address(uint160(ON_RAMP_ADDRESS) + 7),
      metadataHash: ""
    });
    sourceChainConfigs[2] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(1000),
      onRamp: address(uint160(ON_RAMP_ADDRESS) + 42),
      metadataHash: ""
    });

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory expectedSourceChainConfigs =
      new EVM2EVMMultiOffRamp.SourceChainConfig[](3);
    for (uint256 i = 0; i < 3; ++i) {
      expectedSourceChainConfigs[i] = EVM2EVMMultiOffRamp.SourceChainConfig({
        isEnabled: sourceChainConfigs[i].isEnabled,
        prevOffRamp: sourceChainConfigs[i].prevOffRamp,
        onRamp: sourceChainConfigs[i].onRamp,
        metadataHash: s_offRamp.metadataHash(sourceChainSelectors[i], sourceChainConfigs[i].onRamp)
      });

      vm.expectEmit();
      emit SourceChainSelectorAdded(sourceChainSelectors[i]);

      vm.expectEmit();
      emit SourceChainConfigSet(sourceChainSelectors[i], expectedSourceChainConfigs[i]);
    }

    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);

    uint64[] memory resultSourceChainSelectors = s_offRamp.getSourceChainSelectors();
    assertEq(resultSourceChainSelectors.length, 3);

    for (uint256 i = 0; i < 3; ++i) {
      _assertSourceChainConfigEquality(
        s_offRamp.getSourceChainConfig(sourceChainSelectors[i]), expectedSourceChainConfigs[i]
      );

      assertEq(resultSourceChainSelectors[i], sourceChainSelectors[i]);
    }
  }

  function test_MismatchingUpdateLenghts_Revert() public {
    uint64[] memory sourceChainSelectors = new uint64[](2);
    sourceChainSelectors[0] = SOURCE_CHAIN_SELECTOR_1;
    sourceChainSelectors[1] = SOURCE_CHAIN_SELECTOR_1 + 1;

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](1);
    sourceChainConfigs[0] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: ON_RAMP_ADDRESS,
      metadataHash: ""
    });

    vm.expectRevert(EVM2EVMMultiOffRamp.SourceConfigUpdateLengthMismatch.selector);
    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);
  }

  function test_ZeroOnRampAddress_Revert() public {
    uint64[] memory sourceChainSelectors = new uint64[](1);
    sourceChainSelectors[0] = SOURCE_CHAIN_SELECTOR_1;

    EVM2EVMMultiOffRamp.SourceChainConfig[] memory sourceChainConfigs = new EVM2EVMMultiOffRamp.SourceChainConfig[](1);
    sourceChainConfigs[0] = EVM2EVMMultiOffRamp.SourceChainConfig({
      isEnabled: true,
      prevOffRamp: address(0),
      onRamp: address(0),
      metadataHash: ""
    });

    vm.expectRevert(EVM2EVMMultiOffRamp.ZeroAddressNotAllowed.selector);
    s_offRamp.applySourceConfigUpdates(sourceChainSelectors, sourceChainConfigs);
  }

  function _assertSourceChainConfigEquality(
    EVM2EVMMultiOffRamp.SourceChainConfig memory config1,
    EVM2EVMMultiOffRamp.SourceChainConfig memory config2
  ) internal pure {
    assertEq(config1.isEnabled, config2.isEnabled);
    assertEq(config1.prevOffRamp, config2.prevOffRamp);
    assertEq(config1.onRamp, config2.onRamp);
    assertEq(config1.metadataHash, config2.metadataHash);
  }
}
