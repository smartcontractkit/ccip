# Changelog CCIP

All notable changes to the CCIP project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- unreleased -->
## [dev]

...


## 0.5.0 - Unreleased

## Changed

- TokenPools have changed to require an `allowList` parameter in constructor
  - If `allowList` is non-empty
    - only addresses in allowList can be `originalSender` when invoking lockOrBurn
    - addresses can be added to or removed from `allowList` by owner
    - allowList cannot be disabled later now
  - If `allowList` is empty
    - pool is constructed with allowList disabled
    - allowList cannot be enabled later on

## 0.4.0 - 2023-05-24

## Added

- BurnMintERC677 is the new default token that should be deployed whenever there is a need for a burn/mint token
  - Supports ERC677
  - OZ AccessControlEnumerable
  - OZ ERC20Burnable
  - OZ ERC20
  - Compatible with IBurnMintERC20 (CCIP interface)

## Changed

- IBurnMintERC20 interface has changed to follow OZ Burnable tokens
  - New interface
    - function burn(uint256 amount)
    - function burnFrom(address account, uint256 amount)
    - mint(address account, uint256 amount)
  - Old interface
    - function mint(address account, uint256 amount)
- Reduced rate limiting gas usage, this changes the config params to uint128
- Upgrade OZ dependencies to v4.8.0
- Bumped Solidity optimizations from 15k to 30k
- Config changes
  - CommitOffchainConfig
    - SourceIncomingConfirmations renamed to SourceFinalityDepth
    - DestIncomingConfirmations removed
  - ExecOffchainConfig
    - SourceIncomingConfirmations -> SourceFinalityDepth
    - DestIncomingConfirmations -> DestFinalityDepth
    - NEW DestOptimisticConfirmations (required, cannot be 0. Can be DestFinalityDepth)

## Removed

- wrapped token pools
  - Pools should be deployed as burn/mint together with a newly introduced token: BurnMintERC677.
  - This allows us to upgrade the pools without deploying a new token
  - The pool should be allowed to burn and/or mint by calling `grantMintAndBurnRoles(address pool)`


## 0.3.0 - 2023-05-09

### Added
- Added token bps fee to each individual token transfer
  - Fee structure is as follows:
    - bps fee, accurate to 0.1 bps
    - minFee, in US cents, the minimum fee to charge for 1 transfer
    - maxFee, in US cents, the maximum fee to charge for 1 transfer
  - Fee is in the range of [minFee, maxFee] 
  - The fee is configurable per token per lane per direction
  - Edge cases:
    - each token transfer is charged independently, we do not aggregate same-token transfers
    - transfers with 0 token amount is charged the minimum fee
    - all fee fields can be 0
  - The fee is charged in `feeTokens` and added to message execution fee; we do not take breadcrumbs of token transfers

### Changed

- Solidity version bumped to 0.8.19
- AggregateRateLimiter values are now in US dollar amounts with 18 decimals. Previously, it was 36 decimals.
- _setNops calls payNops 
- OnRamp and OffRamp contracts emit PoolAdded event from constructor 
- `EVM2EVMOnRamp.applyAllowListUpdates(address[] calldata removes, address[] calldata adds)` signature changed. Arguments order was `adds`, `removes`

## [0.2.0] - 2023-04-30
## [0.1.0] - 2023-03-14
