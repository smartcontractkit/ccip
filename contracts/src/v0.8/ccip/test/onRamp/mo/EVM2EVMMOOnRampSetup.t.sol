pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import "../../../utils/CCIP.sol";
import "../../../onRamp/interfaces/Any2EVMMOOnRampInterface.sol";
import "../../../onRamp/mo/EVM2AnyMOOnRampRouter.sol";
import "../../../onRamp/mo/EVM2EVMMOOnRamp.sol";

contract EVM2EVMMOOnRampSetup is TokenSetup {
  // Duplicate event of the CCIPSendRequested in the TollOnRampInterface
  event CCIPSendRequested(CCIP.EVM2EVMMOEvent message);
  event OnRampSet(uint256 indexed chainId, Any2EVMMOOnRampInterface indexed onRamp);
  event FeeSet(uint96);
  event SubscriptionFunded(address indexed sender, uint256 amount);
  event SubscriptionUnfunded(address indexed sender, uint256 amount);

  EVM2AnyMOOnRampRouter public s_onRampRouter;
  EVM2EVMMOOnRamp public s_onRamp;
  BaseOnRampInterface.OnRampConfig public s_onRampConfig;

  IERC20 s_sourceFeeToken;
  address[] public s_allowList;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_sourceFeeToken = s_sourceTokens[0];

    s_onRampRouter = new EVM2AnyMOOnRampRouter(Any2EVMMOOnRampRouterInterface.RouterConfig(0, s_sourceFeeToken, OWNER));

    s_onRampConfig = BaseOnRampInterface.OnRampConfig({relayingFeeJuels: 0, maxDataSize: 50, maxTokensLength: 3});

    s_onRamp = new EVM2EVMMOOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      HEARTBEAT,
      s_onRampConfig,
      s_onRampRouter
    );

    s_onRampRouter.setOnRamp(DEST_CHAIN_ID, s_onRamp);

    // Pre approve the fee token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceFeeToken.approve(address(s_onRampRouter), 2**128);
  }

  function getEmptyMessage() public pure returns (CCIP.EVM2AnyMOMessage memory) {
    return CCIP.EVM2AnyMOMessage({receiver: OWNER, data: "", gasLimit: 0});
  }

  function messageToEvent(
    CCIP.EVM2AnyMOMessage memory message,
    uint64 seqNum,
    uint64 nonce
  ) public pure returns (CCIP.EVM2EVMMOEvent memory) {
    return
      CCIP.EVM2EVMMOEvent({
        sequenceNumber: seqNum,
        sourceChainId: SOURCE_CHAIN_ID,
        sender: OWNER,
        receiver: message.receiver,
        nonce: nonce,
        data: message.data,
        gasLimit: message.gasLimit
      });
  }
}
