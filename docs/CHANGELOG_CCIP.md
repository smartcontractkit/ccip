# Changelog CCIP

All notable changes to the CCIP project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- unreleased -->
## [dev]

...

## 0.4.0 - UNRELEASED

## Added

## Changed


<!-- unreleasedstop -->

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
