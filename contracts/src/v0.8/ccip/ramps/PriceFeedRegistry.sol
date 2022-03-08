// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../access/OwnerIsCreator.sol";
import "../../vendor/IERC20.sol";
import "../../interfaces/AggregatorV2V3Interface.sol";

contract PriceFeedRegistry is OwnerIsCreator {
  error InvalidPriceFeedConfig();

  event FeedsSet(IERC20[] tokens, AggregatorV2V3Interface[] feeds);

  // token => price feed
  mapping(IERC20 => AggregatorV2V3Interface) private s_feeds;
  // List of tokens
  IERC20[] private s_tokenList;

  /**
   * @notice The `tokens` and `feeds` passed to this constructor depend on which chain this contract
   * is being deployed to. Mappings of source token => destination feed is maintained on the destination
   * chain. Therefore, when being deployed as an inheriting OffRamp, `tokens` should represent source chain tokens,
   * `feeds` destinations chain feeds. When being deployed as an inheriting OnRamp, `tokens` and `feeds`
   * should both be source chain.
   */
  constructor(IERC20[] memory tokens, AggregatorV2V3Interface[] memory feeds) {
    setFeeds(tokens, feeds);
  }

  /**
   * @notice Remove existing tokens and feeds and set them using the parameters
   * @param tokens token array
   * @param feeds price feed array
   */
  function setFeeds(IERC20[] memory tokens, AggregatorV2V3Interface[] memory feeds) public onlyOwner {
    if (tokens.length != feeds.length || tokens.length == 0) revert InvalidPriceFeedConfig();

    // Unset existing tokens and pools
    IERC20[] memory existingTokens = s_tokenList;
    for (uint256 i = 0; i < existingTokens.length; i++) {
      // Unset feed
      delete s_feeds[existingTokens[i]];
    }

    // Set new tokens and pools
    s_tokenList = tokens;
    for (uint256 i = 0; i < tokens.length; i++) {
      AggregatorV2V3Interface feed = feeds[i];
      s_feeds[tokens[i]] = feed;
    }

    emit FeedsSet(tokens, feeds);
  }

  /**
   * @notice Get a price feed by its token
   * @param token token
   * @return Price feed
   */
  function getFeed(IERC20 token) public view returns (AggregatorV2V3Interface) {
    return s_feeds[token];
  }

  /**
   * @notice Get all configured tokens
   * @return Array of configured tokens
   */
  function getFeedTokens() public view returns (IERC20[] memory) {
    return s_tokenList;
  }
}
