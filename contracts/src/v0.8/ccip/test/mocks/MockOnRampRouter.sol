// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../utils/CCIP.sol";

contract MockOnRampRouter {
  CCIP.EVM2AnyTollMessage public mp;

  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage calldata payload) external returns (uint64) {
    mp = payload;
    return 0;
  }

  function getMessagePayload()
    external
    view
    returns (
      address receiver,
      bytes memory data,
      IERC20[] memory tokens,
      uint256[] memory amounts
    )
  {
    receiver = mp.receiver;
    data = mp.data;
    tokens = mp.tokens;
    amounts = mp.amounts;
  }
}
