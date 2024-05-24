// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IMessageValidator} from "./interfaces/IMessageValidator.sol";
import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {OwnerIsCreator} from "./../shared/access/OwnerIsCreator.sol";
import {EnumerableMapAddresses} from "./../shared/enumerable/EnumerableMapAddresses.sol";
import {Client} from "./libraries/Client.sol";
import {RateLimiterNoEvents} from "./libraries/RateLimiterNoEvents.sol";
import {USDPriceWith18Decimals} from "./libraries/USDPriceWith18Decimals.sol";

/// @notice The aggregate rate limiter is a wrapper of the token bucket rate limiter
/// which permits rate limiting based on the aggregate value of a group of
/// token transfers, using a price registry to convert to a numeraire asset (e.g. USD).
contract MultiAggregateRateLimiter is IMessageValidator, OwnerIsCreator {
  using RateLimiterNoEvents for RateLimiterNoEvents.TokenBucket;
  using USDPriceWith18Decimals for uint224;
  using EnumerableMapAddresses for EnumerableMapAddresses.AddressToAddressMap;

  error UnauthorizedCaller(address caller);
  error PriceNotFoundForToken(address token);
  error UpdateLengthMismatch();
  error ZeroAddressNotAllowed();
  error ZeroChainSelectorNotAllowed();

  event RateLimiterConfigUpdated(uint64 indexed chainSelector, RateLimiterNoEvents.Config config);
  event RateLimiterTokensConsumed(uint64 indexed chainSelector, uint256 tokens);
  event AdminSet(address newAdmin);
  event PriceRegistrySet(address newPriceRegistry);
  event TokenAggregateRateLimitAdded(address sourceToken, address destToken);
  event TokenAggregateRateLimitRemoved(address sourceToken, address destToken);
  event AuthorizedCallerAdded(address caller);
  event AuthorizedCallerRemoved(address caller);

  /// @notice RateLimitToken struct containing both the source and destination token addresses
  struct RateLimitToken {
    address sourceToken;
    address destToken;
  }

  /// @notice Update args for changing the authorized callers
  struct AuthorizedCallerArgs {
    address[] addedCallers;
    address[] removedCallers;
  }

  /// @notice Update args for a single rate limiter config update
  struct RateLimiterConfigArgs {
    uint64 chainSelector; // Chain selector to set config for
    RateLimiterNoEvents.Config rateLimiterConfig; // Rate limiter config to set
  }

  /// @dev Tokens that should be included in Aggregate Rate Limiting (from dest -> source)
  EnumerableMapAddresses.AddressToAddressMap internal s_rateLimitedTokensDestToSource;

  /// @dev Set of callers that can call the validation functions (this is required since the validations modify state)
  mapping(address authorizedCaller => bool isAuthorized) internal s_authorizedCallers;

  /// @notice The address of the token limit admin that has the same permissions as the owner.
  address internal s_admin;
  /// @notice The address of the PriceRegistry used to query token values for ratelimiting
  address internal s_priceRegistry;

  /// @notice Rate limiter token bucket states per chain
  mapping(uint64 chainSelector => RateLimiterNoEvents.TokenBucket rateLimiter) s_rateLimitersByChainSelector;

  /// @param rateLimiterConfigs The RateLimiterNoEvents.Configs per chain containing the capacity and refill rate
  /// of the bucket
  /// @param admin the admin address to set
  /// @param priceRegistry the price registry to set
  /// @param authorizedCallers the authorized callers to set
  constructor(
    RateLimiterConfigArgs[] memory rateLimiterConfigs,
    address admin,
    address priceRegistry,
    address[] memory authorizedCallers
  ) {
    _applyRateLimiterConfigUpdates(rateLimiterConfigs);
    _setAdmin(admin);
    _setPriceRegistry(priceRegistry);
    _applyAuthorizedCallerUpdates(
      AuthorizedCallerArgs({addedCallers: authorizedCallers, removedCallers: new address[](0)})
    );
  }

  /// @inheritdoc IMessageValidator
  function validateIncomingMessage(Client.Any2EVMMessage memory message) external {
    if (!s_authorizedCallers[msg.sender]) {
      revert UnauthorizedCaller(msg.sender);
    }

    uint256 value;
    Client.EVMTokenAmount[] memory destTokenAmounts = message.destTokenAmounts;
    for (uint256 i; i < destTokenAmounts.length; ++i) {
      if (s_rateLimitedTokensDestToSource.contains(destTokenAmounts[i].token)) {
        value += _getTokenValue(destTokenAmounts[i]);
      }
    }

    if (value > 0) _rateLimitValue(message.sourceChainSelector, value);
  }

  /// @inheritdoc IMessageValidator
  function validateOutgoingMessage(Client.EVM2AnyMessage memory message, uint64 destChainSelector) external {
    // TODO: to be implemented (assuming the same rate limiter states are shared for incoming and outgoing messages)
  }

  /// @notice Consumes value from the rate limiter bucket based on the token value given.
  /// @param chainSelector chain selector to apply rate limit to
  /// @param value consumed value
  function _rateLimitValue(uint64 chainSelector, uint256 value) internal {
    s_rateLimitersByChainSelector[chainSelector]._consume(value, address(0));
    emit RateLimiterTokensConsumed(chainSelector, value);
  }

  /// @notice Retrieves the token value for a token using the PriceRegistry
  /// @return tokenValue USD value in 18 decimals
  function _getTokenValue(Client.EVMTokenAmount memory tokenAmount) internal view returns (uint256) {
    // not fetching validated price, as price staleness is not important for value-based rate limiting
    // we only need to verify the price is not 0
    uint224 pricePerToken = IPriceRegistry(s_priceRegistry).getTokenPrice(tokenAmount.token).value;
    if (pricePerToken == 0) revert PriceNotFoundForToken(tokenAmount.token);
    return pricePerToken._calcUSDValueFromTokenAmount(tokenAmount.amount);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @param chainSelector chain selector to retrieve state for
  /// @return The token bucket.
  function currentRateLimiterState(uint64 chainSelector) external view returns (RateLimiterNoEvents.TokenBucket memory) {
    return s_rateLimitersByChainSelector[chainSelector]._currentTokenBucketState();
  }

  /// @notice Applies the provided rate limiter config updates.
  /// @param rateLimiterUpdates Rate limiter updates
  /// @dev should only be callable by the owner or token limit admin
  function applyRateLimiterConfigUpdates(RateLimiterConfigArgs[] memory rateLimiterUpdates) external onlyAdminOrOwner {
    _applyRateLimiterConfigUpdates(rateLimiterUpdates);
  }

  /// @notice Applies the provided rate limiter config updates.
  /// @param rateLimiterUpdates Rate limiter updates
  function _applyRateLimiterConfigUpdates(RateLimiterConfigArgs[] memory rateLimiterUpdates) internal {
    for (uint256 i = 0; i < rateLimiterUpdates.length; ++i) {
      RateLimiterConfigArgs memory updateArgs = rateLimiterUpdates[i];
      RateLimiterNoEvents.Config memory configUpdate = updateArgs.rateLimiterConfig;
      uint64 chainSelector = updateArgs.chainSelector;

      if (chainSelector == 0) {
        revert ZeroChainSelectorNotAllowed();
      }

      RateLimiterNoEvents.TokenBucket storage tokenBucket = s_rateLimitersByChainSelector[chainSelector];

      if (tokenBucket.lastUpdated == 0) {
        // Token bucket needs to be newly added
        s_rateLimitersByChainSelector[chainSelector] = RateLimiterNoEvents.TokenBucket({
          rate: configUpdate.rate,
          capacity: configUpdate.capacity,
          tokens: configUpdate.capacity,
          lastUpdated: uint32(block.timestamp),
          isEnabled: configUpdate.isEnabled
        });
      } else {
        tokenBucket._setTokenBucketConfig(configUpdate);
      }
      emit RateLimiterConfigUpdated(chainSelector, configUpdate);
    }
  }

  /// @notice Get all tokens which are included in Aggregate Rate Limiting.
  /// @return sourceTokens The source representation of the tokens that are rate limited.
  /// @return destTokens The destination representation of the tokens that are rate limited.
  /// @dev the order of IDs in the list is **not guaranteed**, therefore, if ordering matters when
  /// making successive calls, one should keep the blockheight constant to ensure a consistent result.
  function getAllRateLimitTokens() external view returns (address[] memory sourceTokens, address[] memory destTokens) {
    sourceTokens = new address[](s_rateLimitedTokensDestToSource.length());
    destTokens = new address[](s_rateLimitedTokensDestToSource.length());

    for (uint256 i = 0; i < s_rateLimitedTokensDestToSource.length(); ++i) {
      (address destToken, address sourceToken) = s_rateLimitedTokensDestToSource.at(i);
      sourceTokens[i] = sourceToken;
      destTokens[i] = destToken;
    }
    return (sourceTokens, destTokens);
  }

  /// @notice Adds or removes tokens from being used in Aggregate Rate Limiting.
  /// @param removes - A list of one or more tokens to be removed.
  /// @param adds - A list of one or more tokens to be added.
  function updateRateLimitTokens(
    RateLimitToken[] memory removes,
    RateLimitToken[] memory adds
  ) external onlyAdminOrOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      if (s_rateLimitedTokensDestToSource.remove(removes[i].destToken)) {
        emit TokenAggregateRateLimitRemoved(removes[i].sourceToken, removes[i].destToken);
      }
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      address destToken = adds[i].destToken;
      address sourceToken = adds[i].sourceToken;

      if (destToken == address(0) || sourceToken == address(0)) {
        revert ZeroAddressNotAllowed();
      }

      if (s_rateLimitedTokensDestToSource.set(destToken, sourceToken)) {
        emit TokenAggregateRateLimitAdded(sourceToken, destToken);
      }
    }
  }

  /// @return priceRegistry The configured PriceRegistry address
  function getPriceRegistry() external view returns (address) {
    return s_priceRegistry;
  }

  /// @notice Sets the Price Registry address
  /// @param newPriceRegistry the address of the new PriceRegistry
  /// @dev precondition The address must be a non-zero address
  function setPriceRegistry(address newPriceRegistry) external onlyAdminOrOwner {
    _setPriceRegistry(newPriceRegistry);
  }

  /// @notice Sets the Price Registry address
  /// @param newPriceRegistry the address of the new PriceRegistry
  /// @dev precondition The address must be a non-zero address
  function _setPriceRegistry(address newPriceRegistry) internal {
    if (newPriceRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    s_priceRegistry = newPriceRegistry;
    emit PriceRegistrySet(newPriceRegistry);
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @param caller Address to check whether it is an authorized caller
  /// @return flag whether the caller is an authorized caller
  function isAuthorizedCaller(address caller) external view returns (bool) {
    return s_authorizedCallers[caller];
  }

  /// @notice Updates the callers that are authorized to call the message validation functions
  /// @param authorizedCallerArgs Callers to add and remove
  function applyAuthorizedCallerUpdates(AuthorizedCallerArgs memory authorizedCallerArgs) external onlyAdminOrOwner {
    _applyAuthorizedCallerUpdates(authorizedCallerArgs);
  }

  /// @notice Updates the callers that are authorized to call the message validation functions
  /// @param authorizedCallerArgs Callers to add and remove
  function _applyAuthorizedCallerUpdates(AuthorizedCallerArgs memory authorizedCallerArgs) internal {
    address[] memory addedCallers = authorizedCallerArgs.addedCallers;
    for (uint256 i; i < addedCallers.length; ++i) {
      address caller = addedCallers[i];

      if (caller == address(0)) {
        revert ZeroAddressNotAllowed();
      }

      s_authorizedCallers[caller] = true;
      emit AuthorizedCallerAdded(caller);
    }

    address[] memory removedCallers = authorizedCallerArgs.removedCallers;
    for (uint256 i; i < removedCallers.length; ++i) {
      address caller = removedCallers[i];

      if (s_authorizedCallers[caller]) {
        delete s_authorizedCallers[caller];
        emit AuthorizedCallerRemoved(caller);
      }
    }
  }

  /// @notice Gets the token limit admin address.
  /// @return the token limit admin address.
  function getTokenLimitAdmin() external view returns (address) {
    return s_admin;
  }

  /// @notice Sets the token limit admin address.
  /// @param newAdmin the address of the new admin.
  /// @dev setting this to address(0) indicates there is no active admin.
  function setAdmin(address newAdmin) external onlyAdminOrOwner {
    _setAdmin(newAdmin);
  }

  /// @notice Sets the token limit admin address.
  /// @param newAdmin the address of the new admin.
  /// @dev setting this to address(0) indicates there is no active admin.
  function _setAdmin(address newAdmin) internal {
    s_admin = newAdmin;
    emit AdminSet(newAdmin);
  }

  /// @notice a modifier that allows the owner or the s_tokenLimitAdmin call the functions
  /// it is applied to.
  modifier onlyAdminOrOwner() {
    if (msg.sender != owner() && msg.sender != s_admin) revert RateLimiterNoEvents.OnlyCallableByAdminOrOwner();
    _;
  }
}
