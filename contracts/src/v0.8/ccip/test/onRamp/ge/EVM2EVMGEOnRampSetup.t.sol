// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IFeeManager} from "../../../interfaces/fees/IFeeManager.sol";
import {IEVM2EVMGEOnRamp} from "../../../interfaces/onRamp/IEVM2EVMGEOnRamp.sol";

import {EVM2EVMGEOnRamp} from "../../../onRamp/ge/EVM2EVMGEOnRamp.sol";
import {FeeManager} from "../../../fees/FeeManager.sol";
import {GERouter} from "../../../router/GERouter.sol";
import {GERouterSetup} from "../../router/GERouterSetup.t.sol";
import {GE} from "../../../models/GE.sol";
import {GEConsumer} from "../../../models/GEConsumer.sol";
import "../../../offRamp/ge/EVM2EVMGEOffRamp.sol";
import "../../TokenSetup.t.sol";

contract EVM2EVMGEOnRampSetup is TokenSetup, GERouterSetup {
  // Duplicate event of the CCIPSendRequested in the IGEOnRamp
  event CCIPSendRequested(GE.EVM2EVMGEMessage message);

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  bytes32 internal s_metadataHash;

  address[] internal s_allowList;

  EVM2EVMGEOnRamp internal s_onRamp;
  address[] s_offRamps;
  // Naming chosen to not collide with s_feeManager in the offRampSetup since both
  // are imported into the e2e test.
  IFeeManager internal s_IFeeManager;

  function setUp() public virtual override(TokenSetup, GERouterSetup) {
    TokenSetup.setUp();
    GERouterSetup.setUp();

    GE.FeeUpdate[] memory feeUpdates = new GE.FeeUpdate[](1);
    feeUpdates[0] = GE.FeeUpdate({sourceFeeToken: s_sourceTokens[0], destChainId: DEST_CHAIN_ID, linkPerUnitGas: 100});
    address[] memory feeUpdaters = new address[](0);
    s_IFeeManager = new FeeManager(feeUpdates, feeUpdaters, uint128(TWELVE_HOURS));

    s_onRamp = new EVM2EVMGEOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN,
      s_sourceRouter,
      feeManagerConfig(address(s_IFeeManager))
    );

    s_metadataHash = keccak256(
      abi.encode(GE.EVM_2_EVM_GE_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, address(s_onRamp))
    );

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    LockReleaseTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    LockReleaseTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_sourceRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    s_offRamps = new address[](2);
    s_offRamps[0] = address(10);
    s_offRamps[1] = address(11);
    s_sourceRouter.addOffRamp(s_offRamps[0]);
    s_sourceRouter.addOffRamp(s_offRamps[1]);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2**128);
  }

  function _generateTokenMessage() public view returns (GEConsumer.EVM2AnyGEMessage memory) {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = i_tokenAmount0;
    tokensAndAmounts[1].amount = i_tokenAmount1;
    return
      GEConsumer.EVM2AnyGEMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: tokensAndAmounts,
        feeToken: s_sourceFeeToken,
        extraArgs: GEConsumer._argsToBytes(GEConsumer.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _generateEmptyMessage() public view returns (GEConsumer.EVM2AnyGEMessage memory) {
    return
      GEConsumer.EVM2AnyGEMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokensAndAmounts: new Common.EVMTokenAndAmount[](0),
        feeToken: s_sourceFeeToken,
        extraArgs: GEConsumer._argsToBytes(GEConsumer.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _messageToEvent(
    GEConsumer.EVM2AnyGEMessage memory message,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount
  ) public view returns (GE.EVM2EVMGEMessage memory) {
    GE.EVM2EVMGEMessage memory messageEvent = GE.EVM2EVMGEMessage({
      sequenceNumber: seqNum,
      feeTokenAmount: feeTokenAmount,
      sender: OWNER,
      nonce: nonce,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainId: SOURCE_CHAIN_ID,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokensAndAmounts: message.tokensAndAmounts,
      feeToken: message.feeToken,
      messageId: ""
    });

    messageEvent.messageId = GE._hash(messageEvent, s_metadataHash);
    return messageEvent;
  }

  function feeManagerConfig(address feeManagerAddress)
    internal
    view
    returns (IEVM2EVMGEOnRamp.DynamicFeeConfig memory feeConfig)
  {
    return
      IEVM2EVMGEOnRamp.DynamicFeeConfig({
        linkToken: s_sourceTokens[0],
        feeAmount: 1,
        destGasOverhead: 1,
        multiplier: 108e16,
        feeManager: feeManagerAddress,
        destChainId: DEST_CHAIN_ID
      });
  }
}
