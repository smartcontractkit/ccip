// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {TokenAdminRegistry} from "../../pools/TokenAdminRegistry.sol";
import {TokenSetup} from "../TokenSetup.t.sol";

contract TokenAdminRegistrySetup is TokenSetup {
  event AdministratorRegistered(address indexed token, address indexed administrator);
  event PoolSet(address indexed token, address indexed pool);

  function setUp() public virtual override {
    TokenSetup.setUp();
  }
}

contract TokenAdminRegistry_getPools is TokenAdminRegistrySetup {
  function test_getPools_Success() public {
    address[] memory tokens = new address[](1);
    tokens[0] = s_sourceTokens[0];

    address[] memory got = s_tokenAdminRegistry.getPools(tokens);
    assertEq(got.length, 1);
    assertEq(got[0], s_sourcePoolByToken[tokens[0]]);

    got = s_tokenAdminRegistry.getPools(s_sourceTokens);
    assertEq(got.length, s_sourceTokens.length);
    for (uint256 i = 0; i < s_sourceTokens.length; i++) {
      assertEq(got[i], s_sourcePoolByToken[s_sourceTokens[i]]);
    }

    address doesNotExist = makeAddr("doesNotExist");
    tokens[0] = doesNotExist;
    got = s_tokenAdminRegistry.getPools(tokens);
    assertEq(got.length, 1);
    assertEq(got[0], address(0));
  }
}

contract TokenAdminRegistry_getPool is TokenAdminRegistrySetup {
  function test_getPool_Success() public {
    address got = s_tokenAdminRegistry.getPool(s_sourceTokens[0]);
    assertEq(got, s_sourcePoolByToken[s_sourceTokens[0]]);
  }

  function test_getPool_UnsupportedToken_Reverts() public {
    address doesNotExist = makeAddr("doesNotExist");
    vm.expectRevert(abi.encodeWithSelector(TokenAdminRegistry.UnsupportedToken.selector, doesNotExist));
    s_tokenAdminRegistry.getPool(doesNotExist);
  }
}

contract TokenAdminRegistry_setPool is TokenAdminRegistrySetup {
  function test_setPool_Success() public {
    address pool = makeAddr("pool");

    vm.expectEmit();
    emit PoolSet(s_sourceTokens[0], pool);

    s_tokenAdminRegistry.setPool(s_sourceTokens[0], pool);

    assertEq(s_tokenAdminRegistry.getPool(s_sourceTokens[0]), pool);
  }

  function test_setPool_OnlyAdministrator_Revert() public {
    vm.stopPrank();

    vm.expectRevert(
      abi.encodeWithSelector(TokenAdminRegistry.OnlyAdministrator.selector, address(this), s_sourceTokens[0])
    );
    s_tokenAdminRegistry.setPool(s_sourceTokens[0], makeAddr("pool"));
  }
}
