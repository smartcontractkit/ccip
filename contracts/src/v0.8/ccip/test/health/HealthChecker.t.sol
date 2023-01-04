// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../mocks/MockAFN.sol";
import "../helpers/HealthCheckerHelper.sol";

contract HealthCheckerSetup is BaseTest {
  HealthCheckerHelper internal s_healthChecker;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_healthChecker = new HealthCheckerHelper(s_afn);
  }
}

contract HealthChecker_constructor is HealthCheckerSetup {
  // Success

  function testConfigSuccess() public {
    assertEq(address(s_healthChecker.getAFN()), address(s_afn));
  }

  // Reverts

  function testBadConfigReverts() public {
    vm.expectRevert(HealthChecker.BadHealthConfig.selector);
    s_healthChecker = new HealthCheckerHelper(IAFN(address(0)));
  }
}

contract HealthChecker_pause is HealthCheckerSetup {
  event Paused(address account);

  // Success

  function testSuccess() public {
    vm.expectEmit(false, false, false, true);
    emit Paused(OWNER);
    s_healthChecker.pause();
    assertTrue(s_healthChecker.paused());
  }

  // Reverts

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_healthChecker.pause();
  }

  function testAlreadyPausedReverts() public {
    s_healthChecker.pause();
    vm.expectRevert("Pausable: paused");
    s_healthChecker.pause();
  }
}

contract HealthChecker_unpause is HealthCheckerSetup {
  event Unpaused(address account);

  function setUp() public override {
    HealthCheckerSetup.setUp();
    s_healthChecker.pause();
  }

  // Success

  function testSuccess() public {
    vm.expectEmit(false, false, false, true);
    emit Unpaused(OWNER);
    s_healthChecker.unpause();
    assertFalse(s_healthChecker.paused());
  }

  // Reverts

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_healthChecker.unpause();
  }

  function testNotPausedReverts() public {
    s_healthChecker.unpause();
    vm.expectRevert("Pausable: not paused");
    s_healthChecker.unpause();
  }
}

contract HealthChecker_setAFN is HealthCheckerSetup {
  event AFNSet(IAFN oldAFN, IAFN newAFN);

  IAFN internal constant NEW_AFN = IAFN(DUMMY_CONTRACT_ADDRESS);

  // Success

  function testSuccess() public {
    vm.expectEmit(false, false, false, true);
    emit AFNSet(s_afn, NEW_AFN);

    s_healthChecker.setAFN(NEW_AFN);
    assertEq(address(s_healthChecker.getAFN()), DUMMY_CONTRACT_ADDRESS);
  }

  // Reverts

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_healthChecker.setAFN(NEW_AFN);
  }

  function testBadConfigReverts() public {
    vm.expectRevert(HealthChecker.BadHealthConfig.selector);
    s_healthChecker.setAFN(IAFN(ZERO_ADDRESS));
  }
}

contract HealthChecker_isAFNHealthy is HealthCheckerSetup {
  function testHealthySuccess() public {
    s_afn.recoverFromBadSignal();
    assertTrue(s_healthChecker.isAFNHealthy());
  }

  function testNotHealthySuccess() public {
    s_afn.voteBad();
    assertFalse(s_healthChecker.isAFNHealthy());
  }
}

contract HealthChecker_whenHealthyModifier is HealthCheckerSetup {
  function testTrueSuccess() public {
    s_afn.recoverFromBadSignal();
    s_healthChecker.whenHealthyFunction();
  }

  function testFalseReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_healthChecker.whenHealthyFunction();
  }
}
