// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import {EVM2EVMGEOnRampInterface} from "../../../interfaces/onRamp/EVM2EVMGEOnRampInterface.sol";
import {EVM2EVMGEOnRamp} from "../../../onRamp/ge/EVM2EVMGEOnRamp.sol";
import {DynamicFeeCalculatorInterface} from "../../../interfaces/dynamicFeeCalculator/DynamicFeeCalculatorInterface.sol";
import {GasFeeCache, GasFeeCacheInterface} from "../../../dynamicFeeCalculator/GasFeeCache.sol";
import {GERouter} from "../../../router/GERouter.sol";
import {GESRouterSetup} from "../../router/GERouterSetup.t.sol";

contract EVM2EVMGEOnRampSetup is TokenSetup, GESRouterSetup {
  using CCIP for CCIP.EVMExtraArgsV1;

  // Duplicate event of the CCIPSendRequested in the GEOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMGEMessage message);

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  address[] internal s_allowList;

  EVM2EVMGEOnRamp internal s_onRamp;

  function setUp() public virtual override(TokenSetup, GESRouterSetup) {
    TokenSetup.setUp();
    GESRouterSetup.setUp();

    CCIP.FeeUpdate[] memory fees = new CCIP.FeeUpdate[](1);
    fees[0] = CCIP.FeeUpdate({chainId: DEST_CHAIN_ID, gasPrice: 100});
    address[] memory feeUpdaters = new address[](0);
    GasFeeCacheInterface gasFeeCache = new GasFeeCache(fees, feeUpdaters);

    s_onRamp = new EVM2EVMGEOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      getCastedSourceTokens(),
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN,
      s_sourceRouter,
      dynamicFeeCalculatorConfig(address(gasFeeCache))
    );
    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_sourceRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2**128);
  }

  function _generateTokenMessage() public view returns (CCIP.EVM2AnyGEMessage memory) {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = i_tokenAmount0;
    tokensAndAmounts[1].amount = i_tokenAmount1;
    return
      CCIP.EVM2AnyGEMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: tokensAndAmounts,
        feeToken: tokensAndAmounts[0].token,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false})._toBytes()
      });
  }

  function _generateEmptyMessage() public view returns (CCIP.EVM2AnyGEMessage memory) {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](0);
    return
      CCIP.EVM2AnyGEMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: tokensAndAmounts,
        feeToken: getCastedSourceEVMTokenAndAmountsWithZeroAmounts()[0].token,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false})._toBytes()
      });
  }

  function _messageToEvent(
    CCIP.EVM2AnyGEMessage memory message,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount
  ) public view returns (CCIP.EVM2EVMGEMessage memory) {
    CCIP.EVMExtraArgsV1 memory extraArgs = this.fromBytesHelper(message.extraArgs);
    return
      CCIP.EVM2EVMGEMessage({
        sequenceNumber: seqNum,
        feeTokenAmount: feeTokenAmount,
        sender: OWNER,
        nonce: nonce,
        gasLimit: extraArgs.gasLimit,
        strict: extraArgs.strict,
        sourceChainId: SOURCE_CHAIN_ID,
        receiver: abi.decode(message.receiver, (address)),
        data: message.data,
        tokensAndAmounts: message.tokensAndAmounts,
        feeToken: message.feeToken
      });
  }

  // DynamicFeeCalculator
  function dynamicFeeCalculatorConfig(address gasFeeCacheAddress)
    internal
    view
    returns (DynamicFeeCalculatorInterface.DynamicFeeConfig memory feeConfig)
  {
    return
      DynamicFeeCalculatorInterface.DynamicFeeConfig({
        feeToken: s_sourceTokens[0],
        feeAmount: 1,
        destGasOverhead: 1,
        multiplier: 108e16,
        gasFeeCache: gasFeeCacheAddress,
        destChainId: DEST_CHAIN_ID
      });
  }
}
