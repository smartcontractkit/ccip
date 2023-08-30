# Changelog CCIP

All notable changes to the CCIP project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- unreleased -->
## [dev]

...
## 1.1.0 - Unreleased

### Added

### Changed
- Changed OnRamp fee calculation logic and corresponding configuration fields.
  - `destGasOverhead` and `destGasPerPayloadByte` are moved from **FeeTokenConfig** to **DynamicConfig**. These values are same on a given lane regardless of fee token.
  - `networkFeeAmountUSD` is renamed to `networkFeeUSD`. It is now multiples of 0.01 USD, as opposed to 1 wei before.
  - `minFee`, `maxFee` are moved from **TokenTransferFeeConfig** to `minTokenTransferFeeUSD`, `maxTokenTransferFeeUSD` in **FeeTokenConfig**.
  - New fields `destGasOverhead` and `destCalldataOverhead` are added to **TokenTransferFeeConfig**.
    - `destGasOverhead` is the amount of destination token transfer gas, to be billed as part of exec gas fee.
    - `destCalldataOverhead` is the size of additional calldata being passed to destination for token transfers. For example, USDC transfers require additional attestation data.
  - new fields `destCalldataOverhead`, `destGasPerCalldataByte`, `destCalldataMultiplier` are added to **DynamicConfig**.
    - `destCalldataOverhead` is the additional L1 calldata gas on top of EVM2EVMMessage.
    - `destGasPerCalldataByte` is the number of L1 calldata gas to charge per byte of total call data
    - `destCalldataMultiplier` is the multiplier for L1 calldata gas. It is in multiples of 1e-4, or 0.0001. It can represent calldata compression factor.

### Removed
