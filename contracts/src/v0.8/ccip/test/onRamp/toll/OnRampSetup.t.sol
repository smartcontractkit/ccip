pragma solidity ^0.8.13;

import "../../TokenSetup.t.sol";
import "../../../utils/CCIP.sol";

contract OnRampSetup is TokenSetup {
  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMTollEvent message);

  function setUp() public virtual override {
    TokenSetup.setUp();
  }

  function getEmptyMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    IERC20[] memory tokens = new IERC20[](0);
    return
      CCIP.EVM2AnyTollMessage({
        receiver: s_owner,
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: 0,
        gasLimit: 0
      });
  }

  function messageToEvent(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    pure
    returns (CCIP.EVM2EVMTollEvent memory)
  {
    return
      CCIP.EVM2EVMTollEvent({
        sequenceNumber: seqNum,
        sourceChainId: s_sourceChainId,
        sender: s_owner,
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
