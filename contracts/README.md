# Chainlink CCIP Smart Contracts

## Installation

```sh
# via pnpm
$ pnpm add @chainlink/contracts-ccip
# via npm
$ npm install @chainlink/contracts-ccip --save
```

### Directory Structure

```sh
@chainlink/contracts-ccip
├── src # Solidity contracts
│   ├── v0.4
│   ├── v0.5
│   ├── v0.6
│   ├── v0.7
│   └── v0.8
└── abi # ABI json output
    ├── v0.4
    ├── v0.5
    ├── v0.6
    ├── v0.7
    └── v0.8
```

### Usage

The solidity smart contracts themselves can be imported via the `src` directory of `@chainlink/contracts-ccip`:

```solidity
import '@chainlink-ccip/contracts/src/v0.8/ccip/applications/CCIPReceiver.sol';

```

## License

[BUSL-1.1](https://spdx.org/licenses/BUSL-1.1.html)
