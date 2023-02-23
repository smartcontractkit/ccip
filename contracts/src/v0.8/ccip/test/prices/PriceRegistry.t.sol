// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPriceRegistry} from "../../interfaces/prices/IPriceRegistry.sol";

import {Internal} from "../../models/Internal.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {RouterSetup} from "../router/RouterSetup.t.sol";
import {PriceRegistry} from "../../prices/PriceRegistry.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

contract PriceRegistrySetup is TokenSetup, RouterSetup {
  PriceRegistry internal s_priceRegistry;
  // Cheat to store the price updates in storage since struct arrays aren't supported.
  bytes internal s_encodedInitialPriceUpdates;

  function setUp() public virtual override(TokenSetup, RouterSetup) {
    TokenSetup.setUp();
    RouterSetup.setUp();

    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](3);
    // Include USD prices
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceTokens[0], usdPerFeeToken: 5e18});
    feeTokenPriceUpdates[1] = Internal.FeeTokenPriceUpdate({
      sourceFeeToken: s_sourceTokens[1],
      usdPerFeeToken: 2000e18
    });
    feeTokenPriceUpdates[2] = Internal.FeeTokenPriceUpdate({
      sourceFeeToken: s_sourceRouter.getWrappedNative(),
      usdPerFeeToken: 2000e18
    });
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: DEST_CHAIN_ID,
      usdPerUnitGas: 1e6
    });
    s_encodedInitialPriceUpdates = abi.encode(priceUpdates);
    address[] memory priceUpdaters = new address[](0);
    s_priceRegistry = new PriceRegistry(priceUpdates, priceUpdaters, uint32(TWELVE_HOURS));
  }

  function testSetupSuccess() public {
    assertEq(s_priceRegistry.getFeeTokenPrice(s_sourceTokens[0]).value, 5e18);
    assertEq(s_priceRegistry.getFeeTokenPrice(s_sourceTokens[1]).value, 2000e18);
    assertEq(s_priceRegistry.getDestinationChainGasPrice(DEST_CHAIN_ID).value, 1e6);
    assertEq(s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID), 2e5);
    assertEq(s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[1], DEST_CHAIN_ID), 500);
  }
}

contract PriceRegistry_constructor is PriceRegistrySetup {
  function testInvalidStalenessThresholdReverts() public {
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: new Internal.FeeTokenPriceUpdate[](0),
      destChainId: DEST_CHAIN_ID,
      usdPerUnitGas: 1e6
    });

    vm.expectRevert(IPriceRegistry.InvalidStalenessThreshold.selector);
    s_priceRegistry = new PriceRegistry(priceUpdates, new address[](0), 0);
  }
}

contract PriceRegistry_addPriceUpdaters is PriceRegistrySetup {
  function testSuccess() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    s_priceRegistry.addPriceUpdaters(priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 1);
    assertEq(s_priceRegistry.getPriceUpdaters()[0], STRANGER);
  }

  function testOnlyCallableByOwnerReverts() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_priceRegistry.addPriceUpdaters(priceUpdaters);
  }
}

contract PriceRegistry_removePriceUpdaters is PriceRegistrySetup {
  function testSuccess() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    s_priceRegistry.addPriceUpdaters(priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 1);
    assertEq(s_priceRegistry.getPriceUpdaters()[0], STRANGER);
    s_priceRegistry.removePriceUpdaters(priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 0);
  }

  function testOnlyCallableByOwnerReverts() public {
    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = STRANGER;
    s_priceRegistry.addPriceUpdaters(priceUpdaters);
    assertEq(s_priceRegistry.getPriceUpdaters().length, 1);
    assertEq(s_priceRegistry.getPriceUpdaters()[0], STRANGER);
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_priceRegistry.removePriceUpdaters(priceUpdaters);
  }
}

contract PriceRegistry_updatePrices is PriceRegistrySetup {
  // Cheat to store the price updates in storage since struct arrays aren't supported.
  bytes internal s_encodedNewPriceUpdates;

  function setUp() public virtual override {
    PriceRegistrySetup.setUp();
    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](2);
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceTokens[0], usdPerFeeToken: 4e18});
    feeTokenPriceUpdates[1] = Internal.FeeTokenPriceUpdate({
      sourceFeeToken: s_sourceTokens[1],
      usdPerFeeToken: 1800e18
    });
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: DEST_CHAIN_ID,
      usdPerUnitGas: 2e6
    });
    s_encodedNewPriceUpdates = abi.encode(priceUpdates);
  }

  function testSuccess() public {
    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedNewPriceUpdates, (Internal.PriceUpdates));
    s_priceRegistry.updatePrices(priceUpdates);

    assertEq(
      s_priceRegistry.getFeeTokenPrice(s_sourceTokens[0]).value,
      priceUpdates.feeTokenPriceUpdates[0].usdPerFeeToken
    );
    assertEq(
      s_priceRegistry.getFeeTokenPrice(s_sourceTokens[1]).value,
      priceUpdates.feeTokenPriceUpdates[1].usdPerFeeToken
    );
    assertEq(s_priceRegistry.getDestinationChainGasPrice(DEST_CHAIN_ID).value, priceUpdates.usdPerUnitGas);
  }

  function testOnlyCallableByUpdaterOrOwnerReverts() public {
    Internal.PriceUpdates memory priceUpdates = abi.decode(s_encodedNewPriceUpdates, (Internal.PriceUpdates));
    changePrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.OnlyCallableByUpdaterOrOwner.selector));
    s_priceRegistry.updatePrices(priceUpdates);
  }
}

contract PriceRegistry_convertFeeTokenAmountToLinkAmount is PriceRegistrySetup {
  function testSuccess() public {
    Internal.PriceUpdates memory initialPriceUpdates = abi.decode(
      s_encodedInitialPriceUpdates,
      (Internal.PriceUpdates)
    );
    uint256 amount = 3e16;
    uint256 conversionRate = (uint256(initialPriceUpdates.feeTokenPriceUpdates[1].usdPerFeeToken) * 1e18) /
      uint256(initialPriceUpdates.feeTokenPriceUpdates[0].usdPerFeeToken);
    uint256 expected = (amount * conversionRate) / 1e18;
    assertEq(s_priceRegistry.convertFeeTokenAmountToLinkAmount(s_sourceTokens[0], s_sourceTokens[1], amount), expected);
  }

  function testFeeTokenNotSupportedReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertFeeTokenAmountToLinkAmount(DUMMY_CONTRACT_ADDRESS, s_sourceTokens[1], 3e16);
  }

  function testStaleFeeTokenReverts() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](1);
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceTokens[0], usdPerFeeToken: 4e18});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(
        IPriceRegistry.StaleTokenPrice.selector,
        s_sourceTokens[1],
        uint128(TWELVE_HOURS),
        uint128(TWELVE_HOURS + 1)
      )
    );
    s_priceRegistry.convertFeeTokenAmountToLinkAmount(s_sourceTokens[0], s_sourceTokens[1], 3e16);
  }

  function testLinkTokenNotSupportedReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.convertFeeTokenAmountToLinkAmount(s_sourceTokens[0], DUMMY_CONTRACT_ADDRESS, 3e16);
  }

  function testStaleLinkTokenReverts() public {
    vm.warp(block.timestamp + TWELVE_HOURS + 1);

    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](1);
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceTokens[1], usdPerFeeToken: 18e17});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(
        IPriceRegistry.StaleTokenPrice.selector,
        s_sourceTokens[0],
        uint128(TWELVE_HOURS),
        uint128(TWELVE_HOURS + 1)
      )
    );
    s_priceRegistry.convertFeeTokenAmountToLinkAmount(s_sourceTokens[0], s_sourceTokens[1], 3e16);
  }
}

contract PriceRegistry_getFeeTokenBaseUnitsPerUnitGas is PriceRegistrySetup {
  function testGetFeeSuccess() public {
    // 1 unit of gas costs 0.000001 USD -> 1e6
    // 1 LINK costs 5 USD -> 5e18
    // gasPrice / linkPrice = (1e6 * 1e18) / 5e18 = 2e5
    assertEq(s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID), 2e5);
  }

  function testUnsupportedTokenReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.TokenNotSupported.selector, DUMMY_CONTRACT_ADDRESS));
    s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(DUMMY_CONTRACT_ADDRESS, DEST_CHAIN_ID);
  }

  function testUnsupportedChainReverts() public {
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.ChainNotSupported.selector, DEST_CHAIN_ID + 1));
    s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID + 1);
  }

  function testStaleGasPriceReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(abi.encodeWithSelector(IPriceRegistry.StaleGasPrice.selector, DEST_CHAIN_ID, TWELVE_HOURS, diff));
    s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID);
  }

  function testStaleTokenPriceReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);

    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: new Internal.FeeTokenPriceUpdate[](0),
      destChainId: DEST_CHAIN_ID,
      usdPerUnitGas: 1e6
    });
    s_priceRegistry.updatePrices(priceUpdates);

    vm.expectRevert(
      abi.encodeWithSelector(IPriceRegistry.StaleTokenPrice.selector, s_sourceTokens[0], TWELVE_HOURS, diff)
    );
    s_priceRegistry.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID);
  }
}
