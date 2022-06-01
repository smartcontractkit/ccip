// SPDX-License-Identifier: MIT
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

  bytes32 constant merkleRoot = 0x3650fa2e6f647104d005f692c076394ce2201b9b1faf378285fa025aff94a5ce;

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

  function loadReports(uint64 sequenceNumber) private returns (CCIP.AnyToEVMTollMessage memory) {
    IERC20[] memory tokens;
    uint256[] memory amounts;
    IERC20 feeToken = IERC20(0x5FbDB2315678afecb367f032d93F642f64180aa3);
    bytes memory data = abi.encode(0);
    return
      CCIP.AnyToEVMTollMessage({
        sourceChainId: _sourceChainId,
        sequenceNumber: sequenceNumber,
        sender: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC,
        receiver: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC,
        data: data,
        tokens: tokens,
        amounts: amounts,
        feeToken: feeToken,
        feeTokenAmount: 0,
        gasLimit: 0
      });
  }

  function testMerkleMulti1of4() public {
    CCIP.AnyToEVMTollMessage[] memory messages = new CCIP.AnyToEVMTollMessage[](1);
    messages[0] = loadReports(3);
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = 0x5c67041d5ec270155627ea0409140f393b40330294f7592ca2eae2a680143858;
    proofs[1] = 0x503a3697b3f92649b4763051dbcf684d726c1b010285386e08088fc5d324969d;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 0);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti2of4() public {
    CCIP.AnyToEVMTollMessage[] memory messages = new CCIP.AnyToEVMTollMessage[](2);
    messages[0] = loadReports(3);
    messages[1] = loadReports(4);
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = 0x5c67041d5ec270155627ea0409140f393b40330294f7592ca2eae2a680143858;
    proofs[1] = 0x9a88b7f00af1dd298351870f81ba0bdbfb44fb8b5985ab57bc645fa2588f8415;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 4);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti3of4() public {
    CCIP.AnyToEVMTollMessage[] memory messages = new CCIP.AnyToEVMTollMessage[](3);
    messages[0] = loadReports(3);
    messages[1] = loadReports(1);
    messages[2] = loadReports(4);
    bytes32[] memory proofs = new bytes32[](1);
    proofs[0] = 0x9a88b7f00af1dd298351870f81ba0bdbfb44fb8b5985ab57bc645fa2588f8415;

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 5);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }

  function testMerkleMulti4of4() public {
    CCIP.AnyToEVMTollMessage[] memory messages = new CCIP.AnyToEVMTollMessage[](4);
    messages[0] = loadReports(1);
    messages[1] = loadReports(3);
    messages[2] = loadReports(4);
    messages[3] = loadReports(2);
    bytes32[] memory proofs = new bytes32[](0);

    CCIP.ExecutionReport memory report = CCIP.ExecutionReport(messages, proofs, 7);

    assertEq(_offRamp.merkleRoot(report), merkleRoot);
  }
}
