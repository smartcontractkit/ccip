// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../interfaces/pools/PoolInterface.sol";
import "../access/OwnerIsCreator.sol";

contract TokenPoolRegistry is OwnerIsCreator {
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error NoPools();
  error PoolDoesNotExist();
  error TokenPoolMismatch();

  event PoolAdded(IERC20 token, PoolInterface pool);
  event PoolRemoved(IERC20 token, PoolInterface pool);

  struct PoolConfig {
    PoolInterface pool;
    uint96 listIndex;
  }

  // token => token pool
  mapping(IERC20 => PoolConfig) private s_pools;
  // List of tokens
  IERC20[] private s_tokenList;

  /**
   * @notice The `tokens` and `pools` passed to this constructor depend on which chain this contract
   * is being deployed to. Mappings of source token => destination pool is maintained on the destination
   * chain. Therefore, when being deployed as an inheriting OffRamp, `tokens` should represent source chain tokens,
   * `pools` destinations chain pools. When being deployed as an inheriting OnRamp, `tokens` and `pools`
   * should both be source chain.
   */
  constructor(IERC20[] memory tokens, PoolInterface[] memory pools) {
    if (tokens.length != pools.length) revert InvalidTokenPoolConfig();

    // Set new tokens and pools
    s_tokenList = tokens;
    for (uint256 i = 0; i < tokens.length; ++i) {
      PoolInterface pool = pools[i];
      s_pools[tokens[i]] = PoolConfig({pool: pool, listIndex: uint96(i)});
    }
  }

  function addPool(IERC20 token, PoolInterface pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    PoolConfig memory config = s_pools[token];
    // Check if the pool is already set
    if (address(config.pool) != address(0)) revert PoolAlreadyAdded();

    // Set the s_pools with new config values
    config.pool = pool;
    config.listIndex = uint96(s_tokenList.length);
    s_pools[token] = config;

    // Add to the s_tokenList
    s_tokenList.push(token);

    emit PoolAdded(token, pool);
  }

  function removePool(IERC20 token, PoolInterface pool) public onlyOwner {
    // Check that there are any pools to remove
    uint256 listLength = s_tokenList.length;
    if (listLength == 0) revert NoPools();

    PoolConfig memory oldConfig = s_pools[token];
    // Check if the pool exists
    if (address(oldConfig.pool) == address(0)) revert PoolDoesNotExist();
    // Sanity check
    if (address(oldConfig.pool) != address(pool)) revert TokenPoolMismatch();

    // In the list, swap the pool token in question with the last item,
    // Update the index of the item swapped, then pop from the list to remove.

    IERC20 lastItem = s_tokenList[listLength - 1];
    // Perform swap
    s_tokenList[listLength - 1] = s_tokenList[oldConfig.listIndex];
    s_tokenList[oldConfig.listIndex] = lastItem;
    // Update listIndex on moved item
    s_pools[lastItem].listIndex = oldConfig.listIndex;
    // Pop, and delete from mapping
    s_tokenList.pop();
    delete s_pools[token];

    emit PoolRemoved(token, pool);
  }

  /**
   * @notice Get a token pool by its token
   * @param sourceToken token
   * @return Token Pool
   */
  function getPool(IERC20 sourceToken) public view returns (PoolInterface) {
    return s_pools[sourceToken].pool;
  }

  /**
   * @notice Get all configured source tokens
   * @return Array of configured source tokens
   */
  function getPoolTokens() public view returns (IERC20[] memory) {
    return s_tokenList;
  }

  /**
   * @notice Get the destination token from the pool based on a given source token.
   * @param sourceToken The source token
   * @return the destination token
   */
  function getDestinationToken(IERC20 sourceToken) public view returns (IERC20) {
    PoolInterface pool = s_pools[sourceToken].pool;
    if (address(pool) == address(0)) revert PoolDoesNotExist();
    return s_pools[sourceToken].pool.getToken();
  }

  /**
   * @notice Get all configured destination tokens
   * @return tokens Array of configured destination tokens
   */
  function getDestinationTokens() external view returns (IERC20[] memory tokens) {
    tokens = new IERC20[](s_tokenList.length);
    for (uint256 i = 0; i < s_tokenList.length; ++i) {
      tokens[i] = getDestinationToken(s_tokenList[i]);
    }
  }
}
