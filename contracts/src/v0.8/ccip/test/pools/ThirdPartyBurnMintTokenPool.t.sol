// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IBurnMintERC20} from "../../interfaces/pools/IBurnMintERC20.sol";

import "../BaseTest.t.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {ThirdPartyBurnMintTokenPool} from "../../pools/ThirdPartyBurnMintTokenPool.sol";
import {Router} from "../../Router.sol";

contract ThirdPartyBurnMintTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  ThirdPartyBurnMintTokenPool internal s_thirdPartyPool;
  address s_routerAllowedOffRamp = address(234);
  Router s_router;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new MockERC20("LINK", "LNK", OWNER, 2**256 - 1);
    s_router = new Router(address(s_token));

    Router.OnRampUpdate[] memory onRampUpdates = new Router.OnRampUpdate[](0);
    Router.OffRampUpdate[] memory offRampUpdates = new Router.OffRampUpdate[](1);
    address[] memory offRamps = new address[](1);
    offRamps[0] = s_routerAllowedOffRamp;
    offRampUpdates[0] = Router.OffRampUpdate({sourceChainId: SOURCE_CHAIN_ID, offRamps: offRamps});

    s_router.applyRampUpdates(onRampUpdates, offRampUpdates);

    s_thirdPartyPool = new ThirdPartyBurnMintTokenPool(
      IBurnMintERC20(address(s_token)),
      rateLimiterConfig(),
      address(s_router)
    );
  }
}

contract ThirdPartyBurnMintTokenPool_applyRampUpdates is ThirdPartyBurnMintTokenPoolSetup {
  event OnRampAllowanceSet(address onRamp, bool allowed);
  event OffRampAllowanceSet(address onRamp, bool allowed);

  function testApplyRampUpdatesSuccess() public {
    IPool.RampUpdate[] memory onRamps = new IPool.RampUpdate[](1);
    onRamps[0] = IPool.RampUpdate({ramp: address(1), allowed: true});

    IPool.RampUpdate[] memory offRamps = new IPool.RampUpdate[](1);
    offRamps[0] = IPool.RampUpdate({ramp: s_routerAllowedOffRamp, allowed: true});

    vm.expectEmit();
    emit OnRampAllowanceSet(onRamps[0].ramp, true);

    vm.expectEmit();
    emit OffRampAllowanceSet(offRamps[0].ramp, true);

    s_thirdPartyPool.applyRampUpdates(onRamps, offRamps);

    offRamps[0] = IPool.RampUpdate({ramp: s_routerAllowedOffRamp, allowed: false});

    vm.expectEmit();
    emit OffRampAllowanceSet(offRamps[0].ramp, false);

    s_thirdPartyPool.applyRampUpdates(onRamps, offRamps);
  }

  // Reverts

  function testInvalidOffRampReverts() public {
    address invalidOffRamp = address(23456787654321);
    IPool.RampUpdate[] memory offRamps = new IPool.RampUpdate[](1);
    offRamps[0] = IPool.RampUpdate({ramp: invalidOffRamp, allowed: true});

    vm.expectRevert(abi.encodeWithSelector(ThirdPartyBurnMintTokenPool.InvalidOffRamp.selector, invalidOffRamp));

    s_thirdPartyPool.applyRampUpdates(new IPool.RampUpdate[](0), offRamps);
  }
}
