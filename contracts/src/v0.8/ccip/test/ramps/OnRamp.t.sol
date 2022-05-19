pragma solidity ^0.8.13;

import "../mocks/MockERC20.sol";
import "../mocks/MockAFN.sol";
import "../../ramps/OnRamp.sol";
import "../mocks/MockPool.sol";
import "../../../tests/MockV3Aggregator.sol";
import "../mocks/MockOnRampRouter.sol";
import "../../interfaces/OnRampInterface.sol";
import "../helpers/OnRampHelper.sol";
import "../../utils/CCIP.sol";
import "../../ramps/OnRampRouter.sol";
import "forge-std/Test.sol";

contract OnRampTest is Test {
  uint256 _chainID = 1;
  uint256[] _destinationChainIds = [2];
  address public _owner;
  address[] _allowList;

  IERC20[] _tokens;
  MockAFN _afn;
  PoolInterface[] _pools;
  AggregatorV2V3Interface[] _feeds;
  OnRampRouter _router;
  OnRampHelper _onRampObject;

  function setUp() public {
    _owner = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
    // Set the sender to _owner
    vm.startPrank(_owner);

    _router = new OnRampRouter();
    _tokens.push(new MockERC20("LINK", "LNK", _owner, 2**256 - 1));
    _afn = new MockAFN();
    _pools.push(new MockPool(5));
    _feeds.push(new MockV3Aggregator(0, 1));
    OnRampInterface.OnRampConfig memory Config = OnRampInterface.OnRampConfig(address(_router), 0, 2e6, 5);

    _onRampObject = new OnRampHelper(
      _chainID,
      _destinationChainIds,
      _tokens,
      _pools,
      _feeds,
      _allowList,
      _afn,
      1e18,
      Config
    );

    _router.setOnRamp(_destinationChainIds[0], _onRampObject);
  }

  function testCalculateFee() public {
    _onRampObject.publicCalculateFee(_tokens[0]);
  }

  function testRequestXChainSendsExactApprove() public {
    _tokens[0].approve(address(_router), 2**128);
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**128;
    requestCrossChainSend(_tokens, amounts, "");
  }

  function testRequestXChainSends() public {
    _tokens[0].approve(address(_router), 2**128);
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**64;
    requestCrossChainSend(_tokens, amounts, "");
  }

  function requestCrossChainSend(
    IERC20[] memory tokens,
    uint256[] memory amounts,
    bytes memory data
  ) public {
    _router.requestCrossChainSend(CCIP.MessagePayload(tokens, amounts, _destinationChainIds[0], _owner, _owner, data));
  }
}
