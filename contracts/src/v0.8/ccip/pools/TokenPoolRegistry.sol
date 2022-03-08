// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../interfaces/PoolInterface.sol";
import "../access/OwnerIsCreator.sol";

contract TokenPoolRegistry is OwnerIsCreator {
  error InvalidTokenPoolConfig();

  event PoolsSet(IERC20[] tokens, PoolInterface[] pools);

  // token => token pool
  mapping(IERC20 => PoolInterface) private s_pools;
  // List of tokens
  IERC20[] private s_tokenList;
  // Mapping of whether token pools have been configured here
  // Checked when executing messages - make sure the receiver of the message is not a configured pool
  mapping(PoolInterface => bool) private s_tokenPoolConfigured;

  /**
   * @notice The `tokens` and `pools` passed to this constructor depend on which chain this contract
   * is being deployed to. Mappings of source token => destination pool is maintained on the destination
   * chain. Therefore, when being deployed as an inheriting OffRamp, `tokens` should represent source chain tokens,
   * `pools` destinations chain pools. When being deployed as an inheriting OnRamp, `tokens` and `pools`
   * should both be source chain.
   */
  constructor(IERC20[] memory tokens, PoolInterface[] memory pools) {
    setPools(tokens, pools);
  }

  /**
   * @notice Removes the existing tokens and pool and sets using the parameters
   * @param tokens token array
   * @param pools Token Pool array
   */
  function setPools(IERC20[] memory tokens, PoolInterface[] memory pools) public onlyOwner {
    if (tokens.length != pools.length || tokens.length == 0) revert InvalidTokenPoolConfig();

    // Unset existing tokens and pools
    IERC20[] memory existingTokens = s_tokenList;
    for (uint256 i = 0; i < existingTokens.length; i++) {
      IERC20 existingToken = existingTokens[i];
      // Unset s_tokenPoolConfigured
      PoolInterface existingPool = s_pools[existingToken];
      s_tokenPoolConfigured[existingPool] = false;
      // Unset s_pools
      delete s_pools[existingToken];
    }

    // Set new tokens and pools
    s_tokenList = tokens;
    for (uint256 i = 0; i < tokens.length; i++) {
      PoolInterface pool = pools[i];
      s_pools[tokens[i]] = pool;
      s_tokenPoolConfigured[pool] = true;
    }
    emit PoolsSet(tokens, pools);
  }

  /**
   * @notice Get a token pool by its token
   * @param sourceToken token
   * @return Token Pool
   */
  function getPool(IERC20 sourceToken) public view returns (PoolInterface) {
    return s_pools[sourceToken];
  }

  /**
   * @notice Check if a token pool has been configured
   */
  function isPool(address addr) public view returns (bool) {
    return s_tokenPoolConfigured[PoolInterface(addr)];
  }

  /**
   * @notice Get all configured source tokens
   * @return Array of configured source tokens
   */
  function getPoolTokens() public view returns (IERC20[] memory) {
    return s_tokenList;
  }
}
