// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {OCR3Base} from "../../ccip/ocr/OCR3Base.sol";

// GraphOCR3 is an OCR3 contract that implements a multi-chain
// connected graph.
contract GraphOCR3 is OCR3Base {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GraphOCR3 1.0.0";

  struct Neighbor {
    uint256 chainId;
    address contractAddress;
  }

  Neighbor[] internal s_neighbors;
  mapping(bytes32 => bool) internal s_seen;

  event NeighborAdded(uint256 chainId, address contractAddress);

  constructor() OCR3Base() {}

  /// @notice adds a neighbor to the graph
  function addNeighbor(uint256 chainId, address contractAddress) external {
    bytes32 h = keccak256(abi.encode(chainId, contractAddress));
    require(!s_seen[h], "already added");
    s_neighbors.push(Neighbor(chainId, contractAddress));
    s_seen[h] = true;
    emit NeighborAdded(chainId, contractAddress);
  }

  function getNeighbors() external view returns (Neighbor[] memory) {
    return s_neighbors;
  }

  function _report(bytes calldata report, uint64 sequenceNumber) internal override {
    // do nothing
  }
}
