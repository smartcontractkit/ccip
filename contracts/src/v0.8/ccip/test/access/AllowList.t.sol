// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../mocks/MockERC20.sol";
import "../../access/AllowList.sol";

contract AllowListSetup is BaseTest {
  event AllowListSet(address[] allowlist);

  AllowList s_allowList;

  function setUp() public virtual override {
    BaseTest.setUp();
    address[] memory allowedAddresses = new address[](1);
    allowedAddresses[0] = OWNER;
    s_allowList = new AllowList(allowedAddresses);
  }
}

/// @notice #setAllowlistEnabled
contract AllowList_setAllowlistEnabled is AllowListSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_allowList.getAllowlistEnabled());
    s_allowList.setAllowlistEnabled(false);
    assertFalse(s_allowList.getAllowlistEnabled());
    s_allowList.setAllowlistEnabled(true);
    assertTrue(s_allowList.getAllowlistEnabled());
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_allowList.setAllowlistEnabled(true);
  }
}

/// @notice #getAllowlistEnabled
contract AllowList_getAllowlistEnabled is AllowListSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_allowList.getAllowlistEnabled());
    s_allowList.setAllowlistEnabled(false);
    assertFalse(s_allowList.getAllowlistEnabled());
    s_allowList.setAllowlistEnabled(true);
    assertTrue(s_allowList.getAllowlistEnabled());
  }
}

/// @notice #setAllowlist
contract AllowList_setAllowlist is AllowListSetup {
  // Success
  function testSuccess() public {
    address[] memory newAddresses = new address[](2);
    newAddresses[0] = address(1);
    newAddresses[1] = address(2);

    vm.expectEmit(false, false, false, true);
    emit AllowListSet(newAddresses);

    s_allowList.setAllowlist(newAddresses);
    address[] memory setAddresses = s_allowList.getAllowlist();

    assertEq(newAddresses[0], setAddresses[0]);
    assertEq(newAddresses[1], setAddresses[1]);
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    address[] memory newAddresses = new address[](2);
    s_allowList.setAllowlist(newAddresses);
  }
}

/// @notice #getAllowlist
contract AllowList_getAllowlist is AllowListSetup {
  // Success
  function testSuccess() public {
    address[] memory setAddresses = s_allowList.getAllowlist();
    assertEq(OWNER, setAddresses[0]);
  }
}
