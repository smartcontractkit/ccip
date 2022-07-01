pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../utils/CCIP.sol";
import "../../../onRamp/subscription/EVM2EVMSubscriptionOnRamp.sol";
import "../../../onRamp/subscription/EVM2AnySubscriptionOnRampRouter.sol";

contract EVM2EVMSubscriptionOnRampSetup is TokenSetup {
  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMSubscriptionEvent message);
  event OnRampSet(uint256 indexed chainId, Any2EVMSubscriptionOnRampInterface indexed onRamp);
  event FeeSet(uint96);
  event SubscriptionFunded(address indexed sender, uint256 amount);
  event SubscriptionUnfunded(address indexed sender, uint256 amount);

  EVM2AnySubscriptionOnRampRouter public s_onRampRouter;
  EVM2EVMSubscriptionOnRamp public s_onRamp;
  BaseOnRampInterface.OnRampConfig public s_onRampConfig;

  uint256 immutable TOKEN_AMOUNT_0 = 9;
  uint256 immutable TOKEN_AMOUNT_1 = 7;

  IERC20 s_sourceFeeToken;
  address[] public s_allowList;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_sourceFeeToken = s_sourceTokens[0];

    s_onRampRouter = new EVM2AnySubscriptionOnRampRouter(
      Any2EVMSubscriptionOnRampRouterInterface.RouterConfig(0, s_sourceFeeToken, OWNER)
    );

    s_onRampConfig = BaseOnRampInterface.OnRampConfig({relayingFeeJuels: 0, maxDataSize: 50, maxTokensLength: 3});

    s_onRamp = new EVM2EVMSubscriptionOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      1e18,
      s_onRampConfig,
      s_onRampRouter
    );

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the fee token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceFeeToken.approve(address(s_onRampRouter), 2**128);
  }

  function getTokenMessage() public view returns (CCIP.EVM2AnySubscriptionMessage memory) {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = TOKEN_AMOUNT_0;
    amounts[1] = TOKEN_AMOUNT_1;
    IERC20[] memory tokens = s_sourceTokens;
    return CCIP.EVM2AnySubscriptionMessage({receiver: OWNER, data: "", tokens: tokens, amounts: amounts, gasLimit: 0});
  }

  function getEmptyMessage() public pure returns (CCIP.EVM2AnySubscriptionMessage memory) {
    uint256[] memory amounts = new uint256[](0);
    IERC20[] memory tokens = new IERC20[](0);
    return CCIP.EVM2AnySubscriptionMessage({receiver: OWNER, data: "", tokens: tokens, amounts: amounts, gasLimit: 0});
  }

  function messageToEvent(
    CCIP.EVM2AnySubscriptionMessage memory message,
    uint64 seqNum,
    uint64 nonce
  ) public pure returns (CCIP.EVM2EVMSubscriptionEvent memory) {
    return
      CCIP.EVM2EVMSubscriptionEvent({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: message.receiver,
        nonce: nonce,
        data: message.data,
        tokens: message.tokens,
        amounts: message.amounts,
        gasLimit: message.gasLimit
      });
  }
}
