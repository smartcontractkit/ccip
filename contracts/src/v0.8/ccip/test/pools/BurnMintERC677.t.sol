// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {BurnMintERC677} from "../../../shared/token/ERC677/BurnMintERC677.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {Strings} from "../../../vendor/openzeppelin-solidity/v4.8.0/utils/Strings.sol";

contract BurnMintERC677Setup is BaseTest {
  event Transfer(address indexed from, address indexed to, uint256 value);

  BurnMintERC677 internal s_burnMintERC20;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_burnMintERC20 = new BurnMintERC677("Chainlink Token", "LINK", 18, 0);
  }
}

contract BurnMintERC677_mint is BurnMintERC677Setup {
  function testPoolMintSuccess() public {
    uint256 amount = 1e19;
    address offRamp = address(238323465456);
    BurnMintTokenPool pool = new BurnMintTokenPool(s_burnMintERC20, new address[](0), rateLimiterConfig());
    s_burnMintERC20.grantMintAndBurnRoles(address(pool));

    TokenPool.RampUpdate[] memory offRamps = new TokenPool.RampUpdate[](1);
    offRamps[0].ramp = offRamp;
    offRamps[0].allowed = true;
    pool.applyRampUpdates(new TokenPool.RampUpdate[](0), offRamps);

    changePrank(offRamp);

    vm.expectEmit();
    emit Transfer(address(0), OWNER, amount);

    pool.releaseOrMint(bytes(""), OWNER, amount, 0, bytes(""));

    assertEq(s_burnMintERC20.balanceOf(OWNER), amount);
  }
}

contract BurnMintERC677_burn is BurnMintERC677Setup {
  function testPoolBurnSuccess() public {
    uint256 burnAmount = 1e19;
    address onRamp = address(238323465456);
    BurnMintTokenPool pool = new BurnMintTokenPool(s_burnMintERC20, new address[](0), rateLimiterConfig());
    s_burnMintERC20.grantMintAndBurnRoles(address(pool));

    TokenPool.RampUpdate[] memory onRamps = new TokenPool.RampUpdate[](1);
    onRamps[0].ramp = onRamp;
    onRamps[0].allowed = true;
    pool.applyRampUpdates(onRamps, new TokenPool.RampUpdate[](0));

    deal(address(s_burnMintERC20), address(pool), burnAmount);
    changePrank(onRamp);

    vm.expectEmit();
    emit Transfer(address(pool), address(0), burnAmount);

    pool.lockOrBurn(OWNER, bytes(""), burnAmount, 0, bytes(""));

    assertEq(s_burnMintERC20.balanceOf(address(pool)), 0);
  }
}
