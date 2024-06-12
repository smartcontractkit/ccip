// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IPool} from "../../interfaces/IPool.sol";

import {PriceRegistry} from "../../PriceRegistry.sol";
import {Router} from "../../Router.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {EVM2EVMMultiOnRamp} from "../../onRamp/EVM2EVMMultiOnRamp.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {EVM2EVMMultiOnRampHelper} from "../helpers/EVM2EVMMultiOnRampHelper.sol";
import {PriceRegistrySetup} from "../priceRegistry/PriceRegistry.t.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract EVM2EVMMultiOnRampSetup is TokenSetup, PriceRegistrySetup {
  // Duplicate event of the CCIPSendRequested in the IOnRamp
  event CCIPSendRequested(Internal.EVM2EVMMessage message);

  address internal constant CUSTOM_TOKEN = address(12345);
  uint224 internal constant CUSTOM_TOKEN_PRICE = 1e17; // $0.1 CUSTOM

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  bytes32 internal s_metadataHash;

  EVM2EVMMultiOnRampHelper internal s_onRamp;
  address[] internal s_offRamps;

  address internal s_destTokenPool = makeAddr("destTokenPool");
  address internal s_destToken = makeAddr("destToken");

  EVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[] internal s_premiumMultiplierWeiPerEthArgs;
  EVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[] internal s_tokenTransferFeeConfigArgs;

  function setUp() public virtual override(TokenSetup, PriceRegistrySetup) {
    TokenSetup.setUp();
    PriceRegistrySetup.setUp();

    s_priceRegistry.updatePrices(getSingleTokenPriceUpdateStruct(CUSTOM_TOKEN, CUSTOM_TOKEN_PRICE));

    s_premiumMultiplierWeiPerEthArgs.push(
      EVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs({
        token: s_sourceFeeToken,
        premiumMultiplierWeiPerEth: 5e17 // 0.5x
      })
    );
    s_premiumMultiplierWeiPerEthArgs.push(
      EVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs({
        token: s_sourceRouter.getWrappedNative(),
        premiumMultiplierWeiPerEth: 2e18 // 2x
      })
    );

    s_tokenTransferFeeConfigArgs.push();
    s_tokenTransferFeeConfigArgs[0].destChainSelector = DEST_CHAIN_SELECTOR;
    s_tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      EVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs({
        token: s_sourceFeeToken,
        tokenTransferFeeConfig: EVM2EVMMultiOnRamp.TokenTransferFeeConfig({
          minFeeUSDCents: 1_00, // 1 USD
          maxFeeUSDCents: 1000_00, // 1,000 USD
          deciBps: 2_5, // 2.5 bps, or 0.025%
          destGasOverhead: 40_000,
          destBytesOverhead: 32,
          aggregateRateLimitEnabled: true,
          isEnabled: true
        })
      })
    );
    s_tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      EVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs({
        token: s_sourceRouter.getWrappedNative(),
        tokenTransferFeeConfig: EVM2EVMMultiOnRamp.TokenTransferFeeConfig({
          minFeeUSDCents: 50, // 0.5 USD
          maxFeeUSDCents: 500_00, // 500 USD
          deciBps: 5_0, // 5 bps, or 0.05%
          destGasOverhead: 10_000,
          destBytesOverhead: 100,
          aggregateRateLimitEnabled: true,
          isEnabled: true
        })
      })
    );
    s_tokenTransferFeeConfigArgs[0].tokenTransferFeeConfigs.push(
      EVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs({
        token: CUSTOM_TOKEN,
        tokenTransferFeeConfig: EVM2EVMMultiOnRamp.TokenTransferFeeConfig({
          minFeeUSDCents: 2_00, // 1 USD
          maxFeeUSDCents: 2000_00, // 1,000 USD
          deciBps: 10_0, // 10 bps, or 0.1%
          destGasOverhead: 1,
          destBytesOverhead: 200,
          aggregateRateLimitEnabled: true,
          isEnabled: true
        })
      })
    );

    (s_onRamp, s_metadataHash) =
      _deployOnRamp(SOURCE_CHAIN_SELECTOR, address(s_sourceRouter), address(s_tokenAdminRegistry));

    s_offRamps = new address[](2);
    s_offRamps[0] = address(10);
    s_offRamps[1] = address(11);
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](2);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: address(s_onRamp)});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: s_offRamps[0]});
    offRampUpdates[1] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: s_offRamps[1]});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2 ** 128);
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2 ** 128);
  }

  function _generateTokenMessage() public view returns (Client.EVM2AnyMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
    tokenAmounts[0].amount = i_tokenAmount0;
    tokenAmounts[1].amount = i_tokenAmount1;
    return Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
  }

  function _generateSingleTokenMessage(
    address token,
    uint256 amount
  ) public view returns (Client.EVM2AnyMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});

    return Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
  }

  function _generateEmptyMessage() public view returns (Client.EVM2AnyMessage memory) {
    return Client.EVM2AnyMessage({
      receiver: abi.encode(OWNER),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](0),
      feeToken: s_sourceFeeToken,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT}))
    });
  }

  function _messageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount,
    address originalSender
  ) public view returns (Internal.EVM2EVMMessage memory) {
    return _messageToEvent(
      message,
      SOURCE_CHAIN_SELECTOR,
      seqNum,
      nonce,
      feeTokenAmount,
      originalSender,
      s_metadataHash,
      s_tokenAdminRegistry
    );
  }

  function _messageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 sourChainSelector,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount,
    address originalSender,
    bytes32 metadaHash,
    TokenAdminRegistry tokenAdminRegistry
  ) internal view returns (Internal.EVM2EVMMessage memory) {
    // Slicing is only available for calldata. So we have to build a new bytes array.
    bytes memory args = new bytes(message.extraArgs.length - 4);
    for (uint256 i = 4; i < message.extraArgs.length; ++i) {
      args[i - 4] = message.extraArgs[i];
    }
    Internal.EVM2EVMMessage memory messageEvent = Internal.EVM2EVMMessage({
      sequenceNumber: seqNum,
      feeTokenAmount: feeTokenAmount,
      sender: originalSender,
      nonce: nonce,
      gasLimit: abi.decode(args, (Client.EVMExtraArgsV1)).gasLimit,
      strict: false,
      sourceChainSelector: sourChainSelector,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokenAmounts: message.tokenAmounts,
      sourceTokenData: new bytes[](message.tokenAmounts.length),
      feeToken: message.feeToken,
      messageId: ""
    });

    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      address destToken = s_destTokenBySourceToken[message.tokenAmounts[i].token];

      messageEvent.sourceTokenData[i] = abi.encode(
        Internal.SourceTokenData({
          sourcePoolAddress: abi.encode(tokenAdminRegistry.getTokenConfig(message.tokenAmounts[i].token).tokenPool),
          destTokenAddress: abi.encode(destToken),
          extraData: ""
        })
      );
    }

    messageEvent.messageId = Internal._hash(messageEvent, metadaHash);
    return messageEvent;
  }

  function _generateDynamicMultiOnRampConfig(
    address router,
    address priceRegistry,
    address tokenAdminRegistry
  ) internal pure returns (EVM2EVMMultiOnRamp.DynamicConfig memory) {
    return EVM2EVMMultiOnRamp.DynamicConfig({
      router: router,
      priceRegistry: priceRegistry,
      tokenAdminRegistry: tokenAdminRegistry
    });
  }

  function _generateDestChainConfigArgs() internal pure returns (EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory) {
    EVM2EVMMultiOnRamp.DestChainConfigArgs[] memory destChainConfigs = new EVM2EVMMultiOnRamp.DestChainConfigArgs[](1);
    destChainConfigs[0] = EVM2EVMMultiOnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      dynamicConfig: EVM2EVMMultiOnRamp.DestChainDynamicConfig({
        isEnabled: true,
        maxNumberOfTokensPerMsg: MAX_TOKENS_LENGTH,
        destGasOverhead: DEST_GAS_OVERHEAD,
        destGasPerPayloadByte: DEST_GAS_PER_PAYLOAD_BYTE,
        destDataAvailabilityOverheadGas: DEST_DATA_AVAILABILITY_OVERHEAD_GAS,
        destGasPerDataAvailabilityByte: DEST_GAS_PER_DATA_AVAILABILITY_BYTE,
        destDataAvailabilityMultiplierBps: DEST_GAS_DATA_AVAILABILITY_MULTIPLIER_BPS,
        maxDataBytes: MAX_DATA_SIZE,
        maxPerMsgGasLimit: MAX_GAS_LIMIT,
        defaultTokenFeeUSDCents: DEFAULT_TOKEN_FEE_USD_CENTS,
        defaultTokenDestGasOverhead: DEFAULT_TOKEN_DEST_GAS_OVERHEAD,
        defaultTokenDestBytesOverhead: DEFAULT_TOKEN_BYTES_OVERHEAD,
        defaultTxGasLimit: GAS_LIMIT,
        gasMultiplierWeiPerEth: 5e17,
        networkFeeUSDCents: 1_00
      }),
      prevOnRamp: address(0)
    });
    return destChainConfigs;
  }

  function _generateTokenTransferFeeConfigArgs(
    uint256 destChainSelectorLength,
    uint256 tokenLength
  ) internal pure returns (EVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[] memory) {
    EVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs =
      new EVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[](destChainSelectorLength);
    for (uint256 i = 0; i < destChainSelectorLength; ++i) {
      tokenTransferFeeConfigArgs[i].tokenTransferFeeConfigs =
        new EVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[](tokenLength);
    }
    return tokenTransferFeeConfigArgs;
  }

  function _getMultiOnRampNopsAndWeights() internal pure returns (EVM2EVMMultiOnRamp.NopAndWeight[] memory) {
    EVM2EVMMultiOnRamp.NopAndWeight[] memory nopsAndWeights = new EVM2EVMMultiOnRamp.NopAndWeight[](3);
    nopsAndWeights[0] = EVM2EVMMultiOnRamp.NopAndWeight({nop: USER_1, weight: 19284});
    nopsAndWeights[1] = EVM2EVMMultiOnRamp.NopAndWeight({nop: USER_2, weight: 52935});
    nopsAndWeights[2] = EVM2EVMMultiOnRamp.NopAndWeight({nop: USER_3, weight: 8});
    return nopsAndWeights;
  }

  function _deployOnRamp(
    uint64 sourceChainSelector,
    address sourceRouter,
    address tokenAdminRegistry
  ) internal returns (EVM2EVMMultiOnRampHelper, bytes32 metadataHash) {
    EVM2EVMMultiOnRampHelper onRamp = new EVM2EVMMultiOnRampHelper(
      EVM2EVMMultiOnRamp.StaticConfig({
        linkToken: s_sourceTokens[0],
        chainSelector: sourceChainSelector,
        maxNopFeesJuels: MAX_NOP_FEES_JUELS,
        rmnProxy: address(s_mockRMN)
      }),
      _generateDynamicMultiOnRampConfig(sourceRouter, address(s_priceRegistry), tokenAdminRegistry),
      _generateDestChainConfigArgs(),
      getOutboundRateLimiterConfig(),
      s_premiumMultiplierWeiPerEthArgs,
      s_tokenTransferFeeConfigArgs,
      _getMultiOnRampNopsAndWeights()
    );
    onRamp.setAdmin(ADMIN);

    return (
      onRamp,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, sourceChainSelector, DEST_CHAIN_SELECTOR, address(onRamp)))
    );
  }

  function _assertDestChainConfigsEqual(
    EVM2EVMMultiOnRamp.DestChainConfig memory a,
    EVM2EVMMultiOnRamp.DestChainConfig memory b
  ) internal pure {
    assertEq(a.dynamicConfig.isEnabled, b.dynamicConfig.isEnabled);
    assertEq(a.dynamicConfig.maxNumberOfTokensPerMsg, b.dynamicConfig.maxNumberOfTokensPerMsg);
    assertEq(a.dynamicConfig.maxDataBytes, b.dynamicConfig.maxDataBytes);
    assertEq(a.dynamicConfig.maxPerMsgGasLimit, b.dynamicConfig.maxPerMsgGasLimit);
    assertEq(a.dynamicConfig.destGasOverhead, b.dynamicConfig.destGasOverhead);
    assertEq(a.dynamicConfig.destGasPerPayloadByte, b.dynamicConfig.destGasPerPayloadByte);
    assertEq(a.dynamicConfig.destDataAvailabilityOverheadGas, b.dynamicConfig.destDataAvailabilityOverheadGas);
    assertEq(a.dynamicConfig.destGasPerDataAvailabilityByte, b.dynamicConfig.destGasPerDataAvailabilityByte);
    assertEq(a.dynamicConfig.destDataAvailabilityMultiplierBps, b.dynamicConfig.destDataAvailabilityMultiplierBps);
    assertEq(a.dynamicConfig.defaultTokenFeeUSDCents, b.dynamicConfig.defaultTokenFeeUSDCents);
    assertEq(a.dynamicConfig.defaultTokenDestGasOverhead, b.dynamicConfig.defaultTokenDestGasOverhead);
    assertEq(a.dynamicConfig.defaultTokenDestBytesOverhead, b.dynamicConfig.defaultTokenDestBytesOverhead);
    assertEq(a.dynamicConfig.defaultTxGasLimit, b.dynamicConfig.defaultTxGasLimit);
    assertEq(a.prevOnRamp, b.prevOnRamp);
    assertEq(a.sequenceNumber, b.sequenceNumber);
    assertEq(a.metadataHash, b.metadataHash);
  }

  function _assertStaticConfigsEqual(
    EVM2EVMMultiOnRamp.StaticConfig memory a,
    EVM2EVMMultiOnRamp.StaticConfig memory b
  ) internal pure {
    assertEq(a.linkToken, b.linkToken);
    assertEq(a.chainSelector, b.chainSelector);
    assertEq(a.maxNopFeesJuels, b.maxNopFeesJuels);
    assertEq(a.rmnProxy, b.rmnProxy);
  }

  function _assertDynamicConfigsEqual(
    EVM2EVMMultiOnRamp.DynamicConfig memory a,
    EVM2EVMMultiOnRamp.DynamicConfig memory b
  ) internal pure {
    assertEq(a.router, b.router);
    assertEq(a.priceRegistry, b.priceRegistry);
    assertEq(a.tokenAdminRegistry, b.tokenAdminRegistry);
  }

  function _assertTokenTransferFeeConfigEqual(
    EVM2EVMMultiOnRamp.TokenTransferFeeConfig memory a,
    EVM2EVMMultiOnRamp.TokenTransferFeeConfig memory b
  ) internal pure {
    assertEq(a.minFeeUSDCents, b.minFeeUSDCents);
    assertEq(a.maxFeeUSDCents, b.maxFeeUSDCents);
    assertEq(a.deciBps, b.deciBps);
    assertEq(a.destGasOverhead, b.destGasOverhead);
    assertEq(a.destBytesOverhead, b.destBytesOverhead);
    assertEq(a.aggregateRateLimitEnabled, b.aggregateRateLimitEnabled);
    assertEq(a.isEnabled, b.isEnabled);
  }
}
