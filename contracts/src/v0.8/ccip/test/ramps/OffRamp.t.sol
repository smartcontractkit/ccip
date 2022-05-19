pragma solidity ^0.8.13;

import "../mocks/MockERC20.sol";
import "../mocks/MockAFN.sol";
import "../mocks/MockPool.sol";
import "../../../tests/MockV3Aggregator.sol";
import "../../utils/CCIP.sol";
import "../../ramps/OffRampRouter.sol";
import "../helpers/OffRampHelper.sol";
import "forge-std/Test.sol";

contract OffRampTest is Test {
  uint256 _sourceChainId = 1;
  uint256 _destChainId = 2;
  address public _owner;

  bytes32 constant merkleRoot = 0xdc65a6f19eb104e2f359d9b61535e2caaf0cc269aed55247a8c816ee65b8a4d1;

  IERC20[] _sourceTokens;
  PoolInterface[] _pools;
  AggregatorV2V3Interface[] _feeds;
  MockAFN _afn;
  OffRampRouter _router;
  OffRampHelper _offRamp;

  function setUp() public {
    _owner = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
    // Set the sender to _owner
    vm.startPrank(_owner);

    _sourceTokens.push(new MockERC20("LINK", "LNK", _owner, 2**256 - 1));
    _afn = new MockAFN();
    _pools.push(new MockPool(5));
    _feeds.push(new MockV3Aggregator(0, 1));

    _offRamp = new OffRampHelper(_sourceChainId, _destChainId, _sourceTokens, _pools, _feeds, _afn, 1e18, 0, 5);

    OffRampInterface[] memory _offRamps = new OffRampInterface[](1);
    _offRamps[0] = _offRamp;
    _router = new OffRampRouter(_offRamps);
    _offRamp.setRouter(_router);
  }

  function loadReports(uint64 sequenceNumber) private returns (CCIP.Message memory) {
    IERC20[] memory tokens;
    uint256[] memory amounts;
    bytes memory data = abi.encode(0);
    CCIP.MessagePayload memory payload = CCIP.MessagePayload(
      tokens,
      amounts,
      2,
      0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC,
      0x0000000000000000000000000000000000000000,
      data
    );
    return CCIP.Message(1, sequenceNumber, 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC, payload);
  }

  function testStaticReport() public {
    bytes memory report = abi.encode(CCIP.RelayReport("testing a normal reporting phase", 824, 931));
    _offRamp.report(report);
  }

  function testReportFuzzing(
    bytes32 hash,
    uint64 min,
    uint64 max
  ) public {
    vm.assume(min <= max);
    bytes memory report = abi.encode(CCIP.RelayReport(hash, min, max));
    _offRamp.report(report);
  }

  function testMerkleMulti1of4() public {
    CCIP.Message[] memory messages = new CCIP.Message[](1);
    messages[0] = loadReports(1);
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = 0x6fc9d4a4dd6d794ab623fd553e241c545ae9e7612468e7c0979edb6916906a2b;
    proofs[1] = 0x68ac8101a381a1e54aed97a512bfd2ec5f7e8252f709cded851d9715110998cb;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 0);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti2of4() public {
    CCIP.Message[] memory messages = new CCIP.Message[](2);
    messages[0] = loadReports(1);
    messages[1] = loadReports(2);
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = 0x6fc9d4a4dd6d794ab623fd553e241c545ae9e7612468e7c0979edb6916906a2b;
    proofs[1] = 0xda8397089e9ef02255d607848984b83f3e89b726d8461ea18ccf9de3e3a14171;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 4);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti3of4() public {
    CCIP.Message[] memory messages = new CCIP.Message[](3);
    messages[0] = loadReports(1);
    messages[1] = loadReports(3);
    messages[2] = loadReports(2);
    bytes32[] memory proofs = new bytes32[](1);
    proofs[0] = 0xda8397089e9ef02255d607848984b83f3e89b726d8461ea18ccf9de3e3a14171;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 5);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti4of4() public {
    CCIP.Message[] memory messages = new CCIP.Message[](4);
    messages[0] = loadReports(1);
    messages[1] = loadReports(3);
    messages[2] = loadReports(4);
    messages[3] = loadReports(2);
    bytes32[] memory proofs = new bytes32[](0);

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 7);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }
}
