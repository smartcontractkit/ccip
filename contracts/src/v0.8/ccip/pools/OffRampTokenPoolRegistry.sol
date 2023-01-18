// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IPool} from "../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

import {IERC20} from "../../vendor/IERC20.sol";

contract OffRampTokenPoolRegistry is OwnerIsCreator {
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error NoPools();
  error PoolDoesNotExist();
  error TokenPoolMismatch();

  event PoolAdded(IERC20 token, IPool pool);
  event PoolRemoved(IERC20 token, IPool pool);

  struct PoolConfig {
    IPool pool;
    uint96 listIndex;
  }

  // source token => token pool
  mapping(IERC20 => PoolConfig) private s_poolsBySourceToken;
  // dest token => token pool
  mapping(IERC20 => IPool) private s_poolsByDestToken;
  // List of tokens
  IERC20[] private s_sourceTokenList;

  /**
   * @notice The `tokens` and `pools` passed to this constructor depend on which chain this contract
   * is being deployed to. Mappings of source token => destination pool is maintained on the destination
   * chain. Therefore, when being deployed as an inheriting OffRamp, `tokens` should represent source chain tokens,
   * `pools` destinations chain pools. When being deployed as an inheriting OnRamp, `tokens` and `pools`
   * should both be source chain.
   */
  constructor(IERC20[] memory tokens, IPool[] memory pools) {
    if (tokens.length != pools.length) revert InvalidTokenPoolConfig();

    // Set new tokens and pools
    s_sourceTokenList = tokens;
    for (uint256 i = 0; i < tokens.length; ++i) {
      PoolConfig memory poolConfig = PoolConfig({pool: pools[i], listIndex: uint96(i)});
      s_poolsBySourceToken[tokens[i]] = poolConfig;
      s_poolsByDestToken[pools[i].getToken()] = poolConfig.pool;
    }
  }

  function addPool(IERC20 token, IPool pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    PoolConfig memory config = s_poolsBySourceToken[token];
    // Check if the pool is already set
    if (address(config.pool) != address(0)) revert PoolAlreadyAdded();

    // Set the s_pools with new config values
    config.pool = pool;
    config.listIndex = uint96(s_sourceTokenList.length);
    s_poolsBySourceToken[token] = config;
    s_poolsByDestToken[pool.getToken()] = pool;

    // Add to the s_tokenList
    s_sourceTokenList.push(token);

    emit PoolAdded(token, pool);
  }

  function removePool(IERC20 token, IPool pool) public onlyOwner {
    // Check that there are any pools to remove
    uint256 listLength = s_sourceTokenList.length;
    if (listLength == 0) revert NoPools();

    PoolConfig memory oldConfig = s_poolsBySourceToken[token];
    // Check if the pool exists
    if (address(oldConfig.pool) == address(0)) revert PoolDoesNotExist();
    // Sanity check
    if (address(oldConfig.pool) != address(pool)) revert TokenPoolMismatch();

    // In the list, swap the pool token in question with the last item,
    // Update the index of the item swapped, then pop from the list to remove.

    IERC20 lastItem = s_sourceTokenList[listLength - 1];
    // Perform swap
    s_sourceTokenList[listLength - 1] = s_sourceTokenList[oldConfig.listIndex];
    s_sourceTokenList[oldConfig.listIndex] = lastItem;
    // Update listIndex on moved item
    s_poolsBySourceToken[lastItem].listIndex = oldConfig.listIndex;
    // Pop, and delete from mapping
    s_sourceTokenList.pop();
    delete s_poolsByDestToken[pool.getToken()];
    delete s_poolsBySourceToken[token];

    emit PoolRemoved(token, pool);
  }

  /**
   * @notice Get a token pool by its source token
   * @param sourceToken token
   * @return Token Pool
   */
  function getPoolBySourceToken(IERC20 sourceToken) public view returns (IPool) {
    return s_poolsBySourceToken[sourceToken].pool;
  }

  /**
   * @notice Get a token pool by its dest token
   * @param destToken token
   * @return Token Pool
   */
  function getPoolByDestToken(IERC20 destToken) public view returns (IPool) {
    return s_poolsByDestToken[destToken];
  }

  /**
   * @notice Get all supported source tokens
   * @return Array of supported source tokens
   */
  function getSupportedTokens() public view returns (IERC20[] memory) {
    return s_sourceTokenList;
  }

  /**
   * @notice Get the destination token from the pool based on a given source token.
   * @param sourceToken The source token
   * @return the destination token
   */
  function getDestinationToken(IERC20 sourceToken) public view returns (IERC20) {
    IPool pool = s_poolsBySourceToken[sourceToken].pool;
    if (address(pool) == address(0)) revert PoolDoesNotExist();
    return s_poolsBySourceToken[sourceToken].pool.getToken();
  }

  /**
   * @notice Get all configured destination tokens
   * @return tokens Array of configured destination tokens
   */
  function getDestinationTokens() external view returns (IERC20[] memory tokens) {
    tokens = new IERC20[](s_sourceTokenList.length);
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      tokens[i] = getDestinationToken(s_sourceTokenList[i]);
    }
  }
}
