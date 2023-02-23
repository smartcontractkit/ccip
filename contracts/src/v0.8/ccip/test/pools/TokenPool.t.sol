// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {TokenPoolHelper} from "../helpers/TokenPoolHelper.sol";

contract TokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  TokenPoolHelper internal s_tokenPool;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new MockERC20("LINK", "LNK", OWNER, 2**256 - 1);
    s_tokenPool = new TokenPoolHelper(s_token);
  }
}

contract TokenPool_constructor is TokenPoolSetup {
  // Reverts
  function testNullAddressNotAllowedReverts() public {
    vm.expectRevert(IPool.NullAddressNotAllowed.selector);

    s_tokenPool = new TokenPoolHelper(IERC20(address(0)));
  }
}

contract TokenPool_setOnRamp is TokenPoolSetup {
  // Success
  function testSetOnRampTrueSuccess() public {
    s_tokenPool.setOnRamp(USER_1, true);
    assertTrue(s_tokenPool.isOnRamp(USER_1));
  }

  function testSetOnRampFalseSuccess() public {
    s_tokenPool.setOnRamp(USER_1, true);
    s_tokenPool.setOnRamp(USER_1, false);
    assertFalse(s_tokenPool.isOnRamp(USER_1));
  }

  // Reverts
  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_tokenPool.setOnRamp(USER_1, true);
  }
}

contract TokenPool_setOffRamp is TokenPoolSetup {
  // Success
  function testSetOffRampTrueSuccess() public {
    s_tokenPool.setOffRamp(USER_1, true);
    assertTrue(s_tokenPool.isOffRamp(USER_1));
  }

  function testSetOffRampFalseSuccess() public {
    s_tokenPool.setOffRamp(USER_1, true);
    s_tokenPool.setOffRamp(USER_1, false);
    assertFalse(s_tokenPool.isOffRamp(USER_1));
  }

  // Reverts
  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_tokenPool.setOffRamp(USER_1, true);
  }
}

contract TokenPool_pause is TokenPoolSetup {
  // Success
  function testPauseSuccess() public {
    s_tokenPool.pause();
    assertTrue(s_tokenPool.paused());
  }

  // Reverts
  function testPauseReverts() public {
    s_tokenPool.pause();
    vm.expectRevert("Pausable: paused");
    s_tokenPool.pause();
  }

  function testNonOwnerRevets() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_tokenPool.pause();
  }
}

contract TokenPool_unpause is TokenPoolSetup {
  // Success
  function testUnpauseSuccess() public {
    s_tokenPool.pause();
    s_tokenPool.unpause();
    assertFalse(s_tokenPool.paused());
  }

  // Reverts
  function testUnpauseReverts() public {
    vm.expectRevert("Pausable: not paused");
    s_tokenPool.unpause();
  }

  function testNonOwnerRevets() public {
    s_tokenPool.pause();
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_tokenPool.unpause();
  }
}
