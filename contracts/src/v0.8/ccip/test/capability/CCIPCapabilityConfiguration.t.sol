// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";

import {CCIPCapabilityConfiguration} from "../../capability/CCIPCapabilityConfiguration.sol";
import {CCIPCapabilityConfigurationHelper} from "../helpers/CCIPCapabilityConfigurationHelper.sol";

contract CCIPCapabilityConfigurationSetup is Test {
  address public constant OWNER = 0x82ae2B4F57CA5C1CBF8f744ADbD3697aD1a35AFe;
  address public constant CAPABILITY_REGISTRY = 0x272aF4BF7FBFc4944Ed59F914Cd864DfD912D55e;

  CCIPCapabilityConfigurationHelper public s_ccipCC;

  function setUp() public {
    changePrank(OWNER);
    s_ccipCC = new CCIPCapabilityConfigurationHelper(CAPABILITY_REGISTRY);
  }
}

contract CCIPCapabilityConfiguration_ConfigStateMachine is CCIPCapabilityConfigurationSetup {
  // Successful cases.

  function test__stateFromConfigLength_Success() public {
    uint256 configLen = 0;
    CCIPCapabilityConfiguration.ConfigState state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Init));

    configLen = 1;
    state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Running));

    configLen = 2;
    state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Staging));
  }

  function test__validateConfigStateTransition_Success() public {
    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Init, CCIPCapabilityConfiguration.ConfigState.Running
    );

    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Running, CCIPCapabilityConfiguration.ConfigState.Staging
    );

    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Staging, CCIPCapabilityConfiguration.ConfigState.Running
    );
  }

  // Reverts.

  function test_Fuzz__stateFromConfigLength_Reverts(uint256 configLen) public {
    vm.assume(configLen > 2);
    vm.expectRevert(CCIPCapabilityConfiguration.InvalidConfigLength.selector);
    s_ccipCC.stateFromConfigLength(configLen);
  }
}
