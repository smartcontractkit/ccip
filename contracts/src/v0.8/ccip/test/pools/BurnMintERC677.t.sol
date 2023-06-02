// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {BurnMintERC677} from "../../pools/tokens/BurnMintERC677.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {Strings} from "../../../vendor/openzeppelin-solidity/v4.8.0/utils/Strings.sol";

contract BurnMintERC677Setup is BaseTest {
  event Transfer(address indexed from, address indexed to, uint256 value);

  BurnMintERC677 internal s_burnMintERC20;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_burnMintERC20 = new BurnMintERC677("Chainlink Token", "LINK", 18);
  }

  function generateAccessControlError(address caller, bytes32 role) public pure returns (bytes memory) {
    return
      abi.encodePacked(
        "AccessControl: account ",
        Strings.toHexString(caller),
        " is missing role ",
        Strings.toHexString(uint256(role), 32)
      );
  }
}

contract BurnMintERC677_constructor is BurnMintERC677Setup {
  function testConstructorSuccess() public {
    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    s_burnMintERC20 = new BurnMintERC677(name, symbol, decimals);

    assertEq(name, s_burnMintERC20.name());
    assertEq(symbol, s_burnMintERC20.symbol());
    assertEq(decimals, s_burnMintERC20.decimals());
  }
}

contract BurnMintERC677_mint is BurnMintERC677Setup {
  function testBasicMintSuccess() public {
    s_burnMintERC20.grantMintAndBurnRoles(OWNER);
    uint256 amount = 1e18;

    vm.expectEmit();
    emit Transfer(address(0), OWNER, amount);

    s_burnMintERC20.mint(OWNER, amount);

    assertEq(s_burnMintERC20.balanceOf(OWNER), amount);
  }

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

  // Revert

  function testWrongRoleReverts() public {
    vm.expectRevert(generateAccessControlError(OWNER, s_burnMintERC20.getMinterRole()));
    s_burnMintERC20.mint(STRANGER, 1e18);
  }
}

contract BurnMintERC677_burn is BurnMintERC677Setup {
  function testBasicBurnSuccess() public {
    s_burnMintERC20.grantMintAndBurnRoles(OWNER);
    uint256 burnAmount = 1e20;
    deal(address(s_burnMintERC20), OWNER, burnAmount);

    vm.expectEmit();
    emit Transfer(OWNER, address(0), burnAmount);

    s_burnMintERC20.burn(burnAmount);

    assertEq(s_burnMintERC20.balanceOf(OWNER), 0);
  }

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

  // Revert

  function testWrongRoleReverts() public {
    vm.expectRevert(generateAccessControlError(OWNER, s_burnMintERC20.getBurnerRole()));
    s_burnMintERC20.burnFrom(STRANGER, 1e18);
  }
}

contract BurnMintERC677_grantMintAndBurnRoles is BurnMintERC677Setup {
  function testGrantMintAndBurnRolesSuccess() public {
    assertFalse(s_burnMintERC20.hasRole(s_burnMintERC20.getMinterRole(), STRANGER));
    assertFalse(s_burnMintERC20.hasRole(s_burnMintERC20.getBurnerRole(), STRANGER));

    s_burnMintERC20.grantMintAndBurnRoles(STRANGER);

    assertTrue(s_burnMintERC20.hasRole(s_burnMintERC20.getMinterRole(), STRANGER));
    assertTrue(s_burnMintERC20.hasRole(s_burnMintERC20.getBurnerRole(), STRANGER));
  }
}
