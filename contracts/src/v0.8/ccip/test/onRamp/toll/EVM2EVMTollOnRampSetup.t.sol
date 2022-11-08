// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../interfaces/onRamp/EVM2EVMTollOnRampInterface.sol";
import "../../../onRamp/toll/EVM2EVMTollOnRamp.sol";
import "../../../onRamp/toll/EVM2AnyTollOnRampRouter.sol";
import "../../models/ExtraArgs.t.sol";

contract EVM2EVMTollOnRampSetup is TokenSetup {
  using CCIP for CCIP.EVMExtraArgsV1;

  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMTollMessage message);

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
      getCastedSourceTokens(),
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN,
      s_onRampRouter
    );
    uint256[] memory fees = new uint256[](1);
    fees[0] = uint256(RELAYING_FEE_JUELS);
    IERC20[] memory feeTokens = new IERC20[](1);
    feeTokens[0] = IERC20(s_sourceTokens[0]);
    s_onRamp.setFeeConfig(EVM2EVMTollOnRampInterface.FeeConfig({feeTokens: feeTokens, fees: fees}));

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_onRampRouter), 2**128);
  }

  function _generateTokenMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = i_tokenAmount0;
    amounts[1] = i_tokenAmount1;
    address[] memory tokens = s_sourceTokens;
    return
      CCIP.EVM2AnyTollMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: RELAYING_FEE_JUELS,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT})._toBytes()
      });
  }

  function _generateEmptyMessage() public view returns (CCIP.EVM2AnyTollMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    address[] memory tokens = new address[](0);
    return
      CCIP.EVM2AnyTollMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokens: tokens,
        amounts: amounts,
        feeToken: s_sourceTokens[0],
        feeTokenAmount: RELAYING_FEE_JUELS,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT})._toBytes()
      });
  }

  function _messageToEvent(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    return
      CCIP.EVM2EVMTollMessage({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: abi.decode(message.receiver, (address)),
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        feeToken: message.feeToken,
        feeTokenAmount: message.feeTokenAmount - RELAYING_FEE_JUELS,
        gasLimit: this.fromBytesHelper(message.extraArgs).gasLimit
      });
  }

  function _messageToEventNoFee(CCIP.EVM2AnyTollMessage memory message, uint64 seqNum)
    public
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    return
      CCIP.EVM2EVMTollMessage({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: abi.decode(message.receiver, (address)),
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        feeToken: message.feeToken,
        feeTokenAmount: message.feeTokenAmount,
        gasLimit: this.fromBytesHelper(message.extraArgs).gasLimit
      });
  }
}
