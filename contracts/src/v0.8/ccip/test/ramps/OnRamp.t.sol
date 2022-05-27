// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../mocks/MockERC20.sol";
import "../mocks/MockAFN.sol";
import "../../ramps/OnRamp.sol";
import "../mocks/MockPool.sol";
import "../../../tests/MockV3Aggregator.sol";
import "../mocks/MockOnRampRouter.sol";
import "../../interfaces/OnRampInterface.sol";
import "../../utils/CCIP.sol";
import "../../ramps/OnRampRouter.sol";
import "forge-std/Test.sol";

contract OnRampTest is Test {
  uint256 public s_chainID = 1;
  uint256 public s_destinationChainId = 2;
  address public s_owner;
  address[] public s_allowList;

  IERC20[] public s_tokens;
  MockAFN public s_afn;
  PoolInterface[] public s_pools;
  AggregatorV2V3Interface[] public s_feeds;
  OnRampRouter public s_router;
  OnRamp public s_onRamp;

  function setUp() public {
    s_owner = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
    // Set the sender to s_owner
    vm.startPrank(s_owner);

    s_router = new OnRampRouter();
    s_tokens.push(new MockERC20("LINK", "LNK", s_owner, 2**256 - 1));
    s_afn = new MockAFN();
    s_pools.push(new MockPool(5));
    s_feeds.push(new MockV3Aggregator(0, 1));
    OnRampInterface.OnRampConfig memory Config = OnRampInterface.OnRampConfig(address(s_router), 0, 2e6, 5);

    s_onRamp = new OnRamp(
      s_chainID,
      s_destinationChainId,
      s_tokens,
      s_pools,
      s_feeds,
      s_allowList,
      s_afn,
      1e18,
      Config
    );

    s_router.setOnRamp(s_destinationChainId, s_onRamp);
    s_tokens[0].approve(address(s_router), 2**128);
  }

  function testGetRequiredFee() public view {
    s_onRamp.getRequiredFee(s_tokens[0]);
  }

  function testRequestXChainSendsExactApprove() public {
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**128;
    requestCrossChainSend(s_tokens, amounts, "");
  }

  function testRequestXChainSends() public {
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**64;
    requestCrossChainSend(s_tokens, amounts, "");
  }

  function requestCrossChainSend(
    IERC20[] memory tokens,
    uint256[] memory amounts,
    bytes memory data
  ) public {
    s_router.requestCrossChainSend(CCIP.MessagePayload(tokens, amounts, s_destinationChainId, s_owner, s_owner, data));
  }
}
