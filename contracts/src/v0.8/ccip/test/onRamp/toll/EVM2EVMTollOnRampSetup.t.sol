// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../onRamp/toll/EVM2EVMTollOnRamp.sol";
import "../../../onRamp/toll/EVM2AnyTollOnRampRouter.sol";

contract EVM2EVMTollOnRampSetup is TokenSetup {
  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMTollEvent message);

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  address[] internal s_allowList;

  EVM2AnyTollOnRampRouter internal s_onRampRouter;
  EVM2EVMTollOnRamp internal s_onRamp;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_onRampRouter = new EVM2AnyTollOnRampRouter();

    s_onRamp = new EVM2EVMTollOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      onRampConfig(),
      s_onRampRouter
    );

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceTokens[0].approve(address(s_onRampRouter), 2**128);
  }

  function _generateTokenMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = i_tokenAmount0;
    amounts[1] = i_tokenAmount1;
    IERC20[] memory tokens = s_sourceTokens;
    return
      CCIP.EVM2AnyTollMessage({
        receiver: OWNER,
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: RELAYING_FEE_JUELS,
        gasLimit: GAS_LIMIT
      });
  }

  function _generateEmptyMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    IERC20[] memory tokens = new IERC20[](0);
    return
      CCIP.EVM2AnyTollMessage({
        receiver: OWNER,
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: RELAYING_FEE_JUELS,
        gasLimit: GAS_LIMIT
      });
  }

  function _messageToEvent(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    pure
    returns (CCIP.EVM2EVMTollEvent memory)
  {
    return
      CCIP.EVM2EVMTollEvent({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: message.receiver,
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        feeToken: message.feeToken,
        feeTokenAmount: message.feeTokenAmount - RELAYING_FEE_JUELS,
        gasLimit: message.gasLimit
      });
  }

  function _messageToEventNoFee(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    pure
    returns (CCIP.EVM2EVMTollEvent memory)
  {
    return
      CCIP.EVM2EVMTollEvent({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: message.receiver,
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        feeToken: message.feeToken,
        feeTokenAmount: message.feeTokenAmount,
        gasLimit: message.gasLimit
      });
  }
}
