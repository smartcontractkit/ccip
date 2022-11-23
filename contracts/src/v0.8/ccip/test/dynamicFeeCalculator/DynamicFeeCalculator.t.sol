// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {DynamicFeeCalculator} from "../../dynamicFeeCalculator/DynamicFeeCalculator.sol";
import {DynamicFeeCalculatorInterface} from "../../interfaces/dynamicFeeCalculator/DynamicFeeCalculatorInterface.sol";
import {GasFeeCacheSetup} from "./GasFeeCache.t.sol";
import {CCIP} from "../../models/Models.sol";

contract DynamicFeeCalculatorSetup is GasFeeCacheSetup {
  using CCIP for CCIP.EVMExtraArgsV1;

  DynamicFeeCalculator s_dynamicFeeCalculator;

  uint256 internal constant feeAmount = 7264;
  uint256 internal constant destGasOverhead = 9462825;
  uint256 internal constant multiplier = 141e16;

  function setUp() public virtual override {
    GasFeeCacheSetup.setUp();
    s_dynamicFeeCalculator = new DynamicFeeCalculator(
      SOURCE_CHAIN_ID,
      dynamicFeeCalculatorConfig(address(s_gasFeeCache))
    );
  }

  function dynamicFeeCalculatorConfig(address gasFeeCacheAddress)
    internal
    view
    returns (DynamicFeeCalculatorInterface.DynamicFeeConfig memory feeConfig)
  {
    return
      DynamicFeeCalculatorInterface.DynamicFeeConfig({
        feeToken: s_sourceTokens[0],
        feeAmount: feeAmount,
        destGasOverhead: destGasOverhead,
        multiplier: multiplier,
        gasFeeCache: gasFeeCacheAddress,
        destChainId: DEST_CHAIN_ID
      });
  }

  function getMessage() public view returns (CCIP.EVM2AnyGEMessage memory message) {
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
}

contract DynamicFeeCalculator_getFee is DynamicFeeCalculatorSetup {
  function testGetFeeSuccess() public {
    uint256 expectedFee = feeAmount +
      ((GAS_LIMIT + destGasOverhead) * s_gasFeeCache.getFee(DEST_CHAIN_ID) * multiplier) /
      1 ether;
    assertEq(expectedFee, s_dynamicFeeCalculator.getFee(getMessage()));
  }
}
