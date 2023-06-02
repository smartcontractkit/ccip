// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../../BaseTest.t.sol";
import {BurnMintERC677} from "../../../token/ERC677/BurnMintERC677.sol";

import {Strings} from "../../../../vendor/openzeppelin-solidity/v4.8.0/utils/Strings.sol";

contract BurnMintERC677Setup is BaseTest {
  event Transfer(address indexed from, address indexed to, uint256 value);
  event MintAccessGranted(address indexed minter);
  event BurnAccessGranted(address indexed burner);
  event MintAccessRevoked(address indexed minter);
  event BurnAccessRevoked(address indexed burner);

  BurnMintERC677 internal s_burnMintERC20;

  address internal s_mockPool = address(6243783892);
  uint256 internal s_amount = 1e18;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_burnMintERC20 = new BurnMintERC677("Chainlink Token", "LINK", 18, 1e27);

    // Set s_mockPool to be a burner and minter
    s_burnMintERC20.grantMintAndBurnRoles(s_mockPool);
    deal(address(s_burnMintERC20), OWNER, s_amount);
  }
}

contract BurnMintERC677_constructor is BurnMintERC677Setup {
  function testConstructorSuccess() public {
    string memory name = "Chainlink token v2";
    string memory symbol = "LINK2";
    uint8 decimals = 19;
    uint256 maxSupply = 1e33;
    s_burnMintERC20 = new BurnMintERC677(name, symbol, decimals, maxSupply);

    assertEq(name, s_burnMintERC20.name());
    assertEq(symbol, s_burnMintERC20.symbol());
    assertEq(decimals, s_burnMintERC20.decimals());
    assertEq(maxSupply, s_burnMintERC20.maxSupply());
  }
}

contract BurnMintERC677_approve is BurnMintERC677Setup {
  function testApproveSuccess() public {
    uint256 balancePre = s_burnMintERC20.balanceOf(STRANGER);
    uint256 sendingAmount = s_amount / 2;

    s_burnMintERC20.approve(STRANGER, sendingAmount);

    changePrank(STRANGER);

    s_burnMintERC20.transferFrom(OWNER, STRANGER, sendingAmount);

    assertEq(sendingAmount + balancePre, s_burnMintERC20.balanceOf(STRANGER));
  }

  // Reverts

  function testInvalidAddressReverts() public {
    vm.expectRevert();

    s_burnMintERC20.approve(address(s_burnMintERC20), s_amount);
  }
}

contract BurnMintERC677_transfer is BurnMintERC677Setup {
  function testTransferSuccess() public {
    uint256 balancePre = s_burnMintERC20.balanceOf(STRANGER);
    uint256 sendingAmount = s_amount / 2;

    s_burnMintERC20.transfer(STRANGER, sendingAmount);

    assertEq(sendingAmount + balancePre, s_burnMintERC20.balanceOf(STRANGER));
  }

  // Reverts

  function testInvalidAddressReverts() public {
    vm.expectRevert();

    s_burnMintERC20.transfer(address(s_burnMintERC20), s_amount);
  }
}

contract BurnMintERC677_mint is BurnMintERC677Setup {
  function testBasicMintSuccess() public {
    uint256 balancePre = s_burnMintERC20.balanceOf(OWNER);

    s_burnMintERC20.grantMintAndBurnRoles(OWNER);

    vm.expectEmit();
    emit Transfer(address(0), OWNER, s_amount);

    s_burnMintERC20.mint(OWNER, s_amount);

    assertEq(balancePre + s_amount, s_burnMintERC20.balanceOf(OWNER));
  }

  // Revert

  function testSenderNotMinterReverts() public {
    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.SenderNotMinter.selector, OWNER));
    s_burnMintERC20.mint(STRANGER, 1e18);
  }

  function testMaxSupplyExceededReverts() public {
    changePrank(s_mockPool);

    // Mint max supply
    s_burnMintERC20.mint(OWNER, s_burnMintERC20.maxSupply());

    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.MaxSupplyExceeded.selector, s_burnMintERC20.maxSupply() + 1));

    // Attempt to mint 1 more than max supply
    s_burnMintERC20.mint(OWNER, 1);
  }
}

contract BurnMintERC677_burn is BurnMintERC677Setup {
  function testBasicBurnSuccess() public {
    s_burnMintERC20.grantBurnRole(OWNER);
    deal(address(s_burnMintERC20), OWNER, s_amount);

    vm.expectEmit();
    emit Transfer(OWNER, address(0), s_amount);

    s_burnMintERC20.burn(s_amount);

    assertEq(0, s_burnMintERC20.balanceOf(OWNER));
  }

  // Revert

  function testSenderNotBurnerReverts() public {
    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.SenderNotBurner.selector, OWNER));

    s_burnMintERC20.burnFrom(STRANGER, s_amount);
  }

  function testExceedsBalanceReverts() public {
    changePrank(s_mockPool);

    vm.expectRevert("ERC20: burn amount exceeds balance");

    s_burnMintERC20.burn(s_amount * 2);
  }

  function testBurnFromZeroAddressReverts() public {
    s_burnMintERC20.grantBurnRole(address(0));
    changePrank(address(0));

    vm.expectRevert("ERC20: burn from the zero address");

    s_burnMintERC20.burn(0);
  }
}

contract BurnMintERC677_burnFrom is BurnMintERC677Setup {
  function setUp() public virtual override {
    BurnMintERC677Setup.setUp();
  }

  function testBurnFromSuccess() public {
    s_burnMintERC20.approve(s_mockPool, s_amount);

    changePrank(s_mockPool);

    s_burnMintERC20.burnFrom(OWNER, s_amount);

    assertEq(0, s_burnMintERC20.balanceOf(OWNER));
  }

  // Reverts

  function testSenderNotBurnerReverts() public {
    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.SenderNotBurner.selector, OWNER));

    s_burnMintERC20.burnFrom(OWNER, s_amount);
  }

  function testInsufficientAllowanceReverts() public {
    changePrank(s_mockPool);

    vm.expectRevert("ERC20: insufficient allowance");

    s_burnMintERC20.burnFrom(OWNER, s_amount);
  }

  function testExceedsBalanceReverts() public {
    s_burnMintERC20.approve(s_mockPool, s_amount * 2);

    changePrank(s_mockPool);

    vm.expectRevert("ERC20: burn amount exceeds balance");

    s_burnMintERC20.burnFrom(OWNER, s_amount * 2);
  }
}

contract BurnMintERC677_grantRole is BurnMintERC677Setup {
  function testGrantMintAccessSuccess() public {
    assertFalse(s_burnMintERC20.isMinter(STRANGER));

    vm.expectEmit();
    emit MintAccessGranted(STRANGER);

    s_burnMintERC20.grantMintAndBurnRoles(STRANGER);

    assertTrue(s_burnMintERC20.isMinter(STRANGER));

    vm.expectEmit();
    emit MintAccessRevoked(STRANGER);

    s_burnMintERC20.revokeMintRole(STRANGER);

    assertFalse(s_burnMintERC20.isMinter(STRANGER));
  }

  function testGrantBurnAccessSuccess() public {
    assertFalse(s_burnMintERC20.isBurner(STRANGER));

    vm.expectEmit();
    emit BurnAccessGranted(STRANGER);

    s_burnMintERC20.grantBurnRole(STRANGER);

    assertTrue(s_burnMintERC20.isBurner(STRANGER));

    vm.expectEmit();
    emit BurnAccessRevoked(STRANGER);

    s_burnMintERC20.revokeBurnRole(STRANGER);

    assertFalse(s_burnMintERC20.isBurner(STRANGER));
  }

  function testGrantManySuccess() public {
    uint256 numberOfPools = 10;
    address[] memory permissionedAddresses = new address[](numberOfPools + 1);
    permissionedAddresses[0] = s_mockPool;

    for (uint160 i = 0; i < numberOfPools; ++i) {
      permissionedAddresses[i + 1] = address(i);
      s_burnMintERC20.grantMintAndBurnRoles(address(i));
    }

    assertEq(permissionedAddresses, s_burnMintERC20.getBurners());
    assertEq(permissionedAddresses, s_burnMintERC20.getMinters());
  }
}

contract BurnMintERC677_grantMintAndBurnRoles is BurnMintERC677Setup {
  function testGrantMintAndBurnRolesSuccess() public {
    assertFalse(s_burnMintERC20.isMinter(STRANGER));
    assertFalse(s_burnMintERC20.isBurner(STRANGER));

    vm.expectEmit();
    emit MintAccessGranted(STRANGER);
    vm.expectEmit();
    emit BurnAccessGranted(STRANGER);

    s_burnMintERC20.grantMintAndBurnRoles(STRANGER);

    assertTrue(s_burnMintERC20.isMinter(STRANGER));
    assertTrue(s_burnMintERC20.isBurner(STRANGER));
  }
}

contract BurnMintERC677_decreaseApproval is BurnMintERC677Setup {
  function testDecreaseApprovalSuccess() public {
    s_burnMintERC20.approve(s_mockPool, s_amount);
    uint256 allowance = s_burnMintERC20.allowance(OWNER, s_mockPool);
    assertEq(allowance, s_amount);
    s_burnMintERC20.decreaseApproval(s_mockPool, s_amount);
    assertEq(s_burnMintERC20.allowance(OWNER, s_mockPool), allowance - s_amount);
  }
}

contract BurnMintERC677_increaseApproval is BurnMintERC677Setup {
  function testIncreaseApprovalSuccess() public {
    s_burnMintERC20.approve(s_mockPool, s_amount);
    uint256 allowance = s_burnMintERC20.allowance(OWNER, s_mockPool);
    assertEq(allowance, s_amount);
    s_burnMintERC20.increaseApproval(s_mockPool, s_amount);
    assertEq(s_burnMintERC20.allowance(OWNER, s_mockPool), allowance + s_amount);
  }
}
