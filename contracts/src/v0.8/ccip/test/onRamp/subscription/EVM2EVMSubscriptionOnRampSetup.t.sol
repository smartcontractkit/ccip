// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../onRamp/subscription/EVM2EVMSubscriptionOnRamp.sol";

contract EVM2EVMSubscriptionOnRampSetup is TokenSetup {
  using CCIP for CCIP.EVMExtraArgsV1;
  using CCIP for bytes;

  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMSubscriptionMessage message);
  event OnRampSet(uint256 indexed chainId, EVM2EVMSubscriptionOnRampInterface indexed onRamp);
  event OnRampRemoved(uint256 indexed chainId, EVM2EVMSubscriptionOnRampInterface indexed onRamp);
  event FeeSet(uint96);
  event SubscriptionFunded(address indexed sender, uint256 amount);
  event SubscriptionUnfunded(address indexed sender, uint256 amount);

  EVM2AnySubscriptionOnRampRouter internal s_onRampRouter;
  EVM2EVMSubscriptionOnRamp internal s_onRamp;

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  IERC20 internal s_sourceFeeToken;
  address[] internal s_allowList;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_sourceFeeToken = IERC20(s_sourceTokens[0]);

    s_onRampRouter = new EVM2AnySubscriptionOnRampRouter(
      EVM2AnySubscriptionOnRampRouterInterface.RouterConfig(0, s_sourceFeeToken, OWNER)
    );

    s_onRamp = new EVM2EVMSubscriptionOnRamp(
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

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the fee token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceFeeToken.approve(address(s_onRampRouter), 2**64);
  }

  function _generateTokenMessage() public view returns (CCIP.EVM2AnySubscriptionMessage memory) {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = i_tokenAmount0;
    amounts[1] = i_tokenAmount1;
    address[] memory tokens = s_sourceTokens;
    return
      CCIP.EVM2AnySubscriptionMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokens: tokens,
        amounts: amounts,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT})._toBytes()
      });
  }

  function _generateEmptyMessage() public pure returns (CCIP.EVM2AnySubscriptionMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    address[] memory tokens = new address[](0);
    return
      CCIP.EVM2AnySubscriptionMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokens: tokens,
        amounts: amounts,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: GAS_LIMIT})._toBytes()
      });
  }

  function _messageToEvent(
    CCIP.EVM2AnySubscriptionMessage calldata message,
    uint64 seqNum,
    uint64 nonce
  ) external pure returns (CCIP.EVM2EVMSubscriptionMessage memory) {
    return
      CCIP.EVM2EVMSubscriptionMessage({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: abi.decode(message.receiver, (address)),
        nonce: nonce,
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        gasLimit: message.extraArgs._fromBytes().gasLimit
      });
  }
}
