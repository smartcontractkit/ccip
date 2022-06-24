// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../access/OwnerIsCreator.sol";
import "../../vendor/IERC20.sol";
import "../../interfaces/AggregatorV2V3Interface.sol";

contract PriceFeedRegistry is OwnerIsCreator {
  error FeedAlreadyAdded();
  error FeedDoesNotExist();
  error NoFeeds();
  error TokenFeedMismatch();
  error InvalidPriceFeedConfig();

  event FeedAdded(IERC20 token, AggregatorV2V3Interface feed);
  event FeedRemoved(IERC20 token, AggregatorV2V3Interface feed);

  struct FeedConfig {
    AggregatorV2V3Interface feed;
    uint96 listIndex;
  }

  // token => price feed
  mapping(IERC20 => FeedConfig) private s_feeds;
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
    if (tokens.length != feeds.length) revert InvalidPriceFeedConfig();

    s_tokenList = tokens;
    for (uint256 i = 0; i < tokens.length; i++) {
      AggregatorV2V3Interface feed = feeds[i];
      s_feeds[tokens[i]] = FeedConfig({feed: feed, listIndex: uint96(i)});
    }
  }

  function addFeed(IERC20 token, AggregatorV2V3Interface feed) public onlyOwner {
    if (address(token) == address(0) || address(feed) == address(0)) revert InvalidPriceFeedConfig();
    FeedConfig memory config = s_feeds[token];
    // Check if the feed is already set
    if (address(config.feed) != address(0)) revert FeedAlreadyAdded();

    // Set the s_feeds with new config values
    config.feed = feed;
    config.listIndex = uint96(s_tokenList.length);
    s_feeds[token] = config;

    // Add to the s_tokenList
    s_tokenList.push(token);

    emit FeedAdded(token, feed);
  }

  function removeFeed(IERC20 token, AggregatorV2V3Interface feed) public onlyOwner {
    // Check that there are any feeds to remove
    uint256 listLength = s_tokenList.length;
    if (listLength == 0) revert NoFeeds();

    FeedConfig memory oldConfig = s_feeds[token];
    // Check if the feed exists
    if (address(oldConfig.feed) == address(0)) revert FeedDoesNotExist();
    // Sanity check
    if (address(oldConfig.feed) != address(feed)) revert TokenFeedMismatch();

    // In the list, swap the feed token in question with the last item,
    // Update the index of the item swapped, then pop from the list to remove.

    IERC20 lastItem = s_tokenList[listLength - 1];
    // Perform swap
    s_tokenList[listLength - 1] = s_tokenList[oldConfig.listIndex];
    s_tokenList[oldConfig.listIndex] = lastItem;
    // Update listIndex on moved item
    s_feeds[lastItem].listIndex = oldConfig.listIndex;
    // Pop, and delete from mapping
    s_tokenList.pop();
    delete s_feeds[token];

    emit FeedRemoved(token, feed);
  }

  /**
   * @notice Get a price feed by its token
   * @param token token
   * @return Price feed
   */
  function getFeed(IERC20 token) public view returns (AggregatorV2V3Interface) {
    return s_feeds[token].feed;
  }

  /**
   * @notice Get all configured tokens
   * @return Array of configured tokens
   */
  function getFeedTokens() public view returns (IERC20[] memory) {
    return s_tokenList;
  }
}
