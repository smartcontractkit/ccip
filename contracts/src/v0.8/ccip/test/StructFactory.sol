// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ARM} from "../ARM.sol";
import "../offRamp/EVM2EVMOffRamp.sol";
import "../onRamp/EVM2EVMOnRamp.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {Internal} from "../libraries/Internal.sol";

contract StructFactory {
  // addresses
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant STRANGER = address(999999);
  address internal constant DUMMY_CONTRACT_ADDRESS = 0x1111111111111111111111111111111111111112;
  address internal constant ON_RAMP_ADDRESS = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;
  address internal constant ZERO_ADDRESS = address(0);
  address internal constant BLESS_VOTER_1 = address(1);
  address internal constant CURSE_VOTER_1 = address(10);
  address internal constant CURSE_UNVOTER_1 = address(110);
  address internal constant BLESS_VOTER_2 = address(2);
  address internal constant CURSE_VOTER_2 = address(12);
  address internal constant CURSE_UNVOTER_2 = address(112);
  address internal constant BLESS_VOTER_3 = address(3);
  address internal constant CURSE_VOTER_3 = address(13);
  address internal constant CURSE_UNVOTER_3 = address(113);
  address internal constant BLESS_VOTER_4 = address(4);
  address internal constant CURSE_VOTER_4 = address(14);
  address internal constant CURSE_UNVOTER_4 = address(114);

  address internal constant USER_1 = address(1);
  address internal constant USER_2 = address(2);
  address internal constant USER_3 = address(3);
  address internal constant USER_4 = address(4);

  // arm

  function armConstructorArgs() internal pure returns (ARM.Config memory) {
    ARM.Voter[] memory voters = new ARM.Voter[](4);
    voters[0] = ARM.Voter({
      blessVoteAddr: BLESS_VOTER_1,
      curseVoteAddr: CURSE_VOTER_1,
      curseUnvoteAddr: CURSE_UNVOTER_1,
      blessWeight: WEIGHT_1,
      curseWeight: WEIGHT_1
    });
    voters[1] = ARM.Voter({
      blessVoteAddr: BLESS_VOTER_2,
      curseVoteAddr: CURSE_VOTER_2,
      curseUnvoteAddr: CURSE_UNVOTER_2,
      blessWeight: WEIGHT_10,
      curseWeight: WEIGHT_10
    });
    voters[2] = ARM.Voter({
      blessVoteAddr: BLESS_VOTER_3,
      curseVoteAddr: CURSE_VOTER_3,
      curseUnvoteAddr: CURSE_UNVOTER_3,
      blessWeight: WEIGHT_20,
      curseWeight: WEIGHT_20
    });
    voters[3] = ARM.Voter({
      blessVoteAddr: BLESS_VOTER_4,
      curseVoteAddr: CURSE_VOTER_4,
      curseUnvoteAddr: CURSE_UNVOTER_4,
      blessWeight: WEIGHT_40,
      curseWeight: WEIGHT_40
    });
    return
      ARM.Config({
        voters: voters,
        blessWeightThreshold: WEIGHT_10 + WEIGHT_20 + WEIGHT_40,
        curseWeightThreshold: WEIGHT_1 + WEIGHT_10 + WEIGHT_20 + WEIGHT_40
      });
  }

  uint8 internal constant ZERO = 0;
  uint8 internal constant WEIGHT_1 = 1;
  uint8 internal constant WEIGHT_10 = 10;
  uint8 internal constant WEIGHT_20 = 20;
  uint8 internal constant WEIGHT_40 = 40;

  // message info
  uint64 internal constant SOURCE_CHAIN_ID = 1;
  uint64 internal constant DEST_CHAIN_ID = 2;
  uint64 internal constant GAS_LIMIT = 200_000;

  // timing
  uint256 internal constant BLOCK_TIME = 1234567890;
  uint32 internal constant TWELVE_HOURS = 60 * 60 * 12;

  // onramp
  uint96 internal constant MAX_NOP_FEES_JUELS = 1e27;
  uint32 internal constant DEST_GAS_OVERHEAD = 350_000;
  uint16 internal constant DEST_GAS_PER_PAYLOAD_BYTE = 16;

  // offRamp
  uint256 internal constant POOL_BALANCE = 1e25;
  uint32 internal constant EXECUTION_DELAY_SECONDS = 0;
  uint24 internal constant MAX_DATA_SIZE = 30_000;
  uint16 internal constant MAX_TOKENS_LENGTH = 5;
  uint16 internal constant GAS_FOR_CALL_EXACT_CHECK = 5000;
  uint32 internal constant PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS = 500;
  uint32 internal constant MAX_GAS_LIMIT = 4_000_000;

  function generateManualGasLimit(uint256 callDataLength) internal view returns (uint256) {
    return ((gasleft() - 2 * (16 * callDataLength + GAS_FOR_CALL_EXACT_CHECK)) * 62) / 64;
  }

  function generateDynamicOffRampConfig(
    address router,
    address priceRegistry
  ) internal pure returns (EVM2EVMOffRamp.DynamicConfig memory) {
    return
      EVM2EVMOffRamp.DynamicConfig({
        router: router,
        maxDataSize: MAX_DATA_SIZE,
        priceRegistry: priceRegistry,
        maxTokensLength: MAX_TOKENS_LENGTH,
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS
      });
  }

  function generateDynamicOnRampConfig(
    address router,
    address priceRegistry
  ) internal pure returns (EVM2EVMOnRamp.DynamicConfig memory) {
    return
      EVM2EVMOnRamp.DynamicConfig({
        router: router,
        maxTokensLength: MAX_TOKENS_LENGTH,
        destGasOverhead: DEST_GAS_OVERHEAD,
        destGasPerPayloadByte: DEST_GAS_PER_PAYLOAD_BYTE,
        destCalldataOverhead: 188,
        destGasPerCalldataByte: 16,
        destCalldataMultiplier: 0,
        priceRegistry: priceRegistry,
        maxDataSize: MAX_DATA_SIZE,
        maxGasLimit: MAX_GAS_LIMIT
      });
  }

  function getTokensAndPools(
    address[] memory sourceTokens,
    IPool[] memory pools
  ) internal pure returns (Internal.PoolUpdate[] memory) {
    Internal.PoolUpdate[] memory tokensAndPools = new Internal.PoolUpdate[](sourceTokens.length);
    for (uint256 i = 0; i < sourceTokens.length; ++i) {
      tokensAndPools[i] = Internal.PoolUpdate({token: sourceTokens[i], pool: address(pools[i])});
    }
    return tokensAndPools;
  }

  function getNopsAndWeights() internal pure returns (EVM2EVMOnRamp.NopAndWeight[] memory) {
    EVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights = new EVM2EVMOnRamp.NopAndWeight[](3);
    nopsAndWeights[0] = EVM2EVMOnRamp.NopAndWeight({nop: USER_1, weight: 19284});
    nopsAndWeights[1] = EVM2EVMOnRamp.NopAndWeight({nop: USER_2, weight: 52935});
    nopsAndWeights[2] = EVM2EVMOnRamp.NopAndWeight({nop: USER_3, weight: 8});
    return nopsAndWeights;
  }

  // Rate limiter
  address constant ADMIN = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;

  function rateLimiterConfig() internal pure returns (RateLimiter.Config memory) {
    return RateLimiter.Config({isEnabled: true, capacity: 100e28, rate: 1e15});
  }

  function getSinglePriceUpdateStruct(
    address token,
    uint224 price
  ) internal pure returns (Internal.PriceUpdates memory) {
    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](1);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: token, usdPerToken: price});

    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: tokenPriceUpdates,
      destChainSelector: 0,
      usdPerUnitGas: 0
    });

    return priceUpdates;
  }

  function getPriceUpdatesStruct(
    address[] memory tokens,
    uint224[] memory prices
  ) internal pure returns (Internal.PriceUpdates memory) {
    uint256 length = tokens.length;

    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](length);
    for (uint256 i = 0; i < length; ++i) {
      tokenPriceUpdates[i] = Internal.TokenPriceUpdate({sourceToken: tokens[i], usdPerToken: prices[i]});
    }
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: tokenPriceUpdates,
      destChainSelector: 0,
      usdPerUnitGas: 0
    });

    return priceUpdates;
  }

  // OffRamp
  function getEmptyPriceUpdates() internal pure returns (Internal.PriceUpdates memory priceUpdates) {
    return
      Internal.PriceUpdates({
        tokenPriceUpdates: new Internal.TokenPriceUpdate[](0),
        destChainSelector: 0,
        usdPerUnitGas: 0
      });
  }
}
