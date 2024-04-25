// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {IPool} from "../../interfaces/pools/IPool.sol";

import {CallWithExactGas} from "../../../shared/call/CallWithExactGas.sol";

import {ARM} from "../../ARM.sol";
import {AggregateRateLimiter} from "../../AggregateRateLimiter.sol";
import {Router} from "../../Router.sol";
import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {EVM2EVMMultiOffRamp} from "../../offRamp/EVM2EVMMultiOffRamp.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {EVM2EVMMultiOffRampHelper} from "../helpers/EVM2EVMMultiOffRampHelper.sol";
import {MaybeRevertingBurnMintTokenPool} from "../helpers/MaybeRevertingBurnMintTokenPool.sol";
import {ConformingReceiver} from "../helpers/receivers/ConformingReceiver.sol";
import {MaybeRevertMessageReceiver} from "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import {MaybeRevertMessageReceiverNo165} from "../helpers/receivers/MaybeRevertMessageReceiverNo165.sol";
import {ReentrancyAbuser} from "../helpers/receivers/ReentrancyAbuser.sol";
import {MockCommitStore} from "../mocks/MockCommitStore.sol";
import {OCR2Base} from "../ocr/OCR2Base.t.sol";
import {OCR2BaseNoChecks} from "../ocr/OCR2BaseNoChecks.t.sol";
import {EVM2EVMMultiOffRampSetup} from "./EVM2EVMMultiOffRampSetup.t.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

// TODO: re-add tests:
//       - constructor
//       - ccipReceive
//       - execute
//       - execute_upgrade
//       - executeSingleMessage
//       - report
//       - manuallyExecute
//       - getExecutionState
//       - trialExecute
//       - releaseOrMintTokens
//       - getAllRateLimitTokens
//       - updateRateLimitTokens

contract EVM2EVMMultiOffRamp_setDynamicConfig is EVM2EVMMultiOffRampSetup {
  // OffRamp event
  event ConfigSet(EVM2EVMMultiOffRamp.StaticConfig staticConfig, EVM2EVMMultiOffRamp.DynamicConfig dynamicConfig);

  function test_SetDynamicConfig_Success() public {
    EVM2EVMMultiOffRamp.StaticConfig memory staticConfig = s_offRamp.getStaticConfig();
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig = generateDynamicMultiOffRampConfig(USER_3, address(s_priceRegistry));
    bytes memory onchainConfig = abi.encode(dynamicConfig);

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    vm.expectEmit();
    uint32 configCount = 1;
    emit ConfigSet(
      uint32(block.number),
      getBasicConfigDigest(address(s_offRamp), s_f, configCount, onchainConfig),
      configCount + 1,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, onchainConfig, s_offchainConfigVersion, abi.encode("")
    );

    EVM2EVMMultiOffRamp.DynamicConfig memory newConfig = s_offRamp.getDynamicConfig();
    _assertSameConfig(dynamicConfig, newConfig);
  }

  function test_NonOwner_Revert() public {
    vm.startPrank(STRANGER);
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig = generateDynamicMultiOffRampConfig(USER_3, address(s_priceRegistry));

    vm.expectRevert("Only callable by owner");

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, abi.encode(dynamicConfig), s_offchainConfigVersion, abi.encode("")
    );
  }

  function test_RouterZeroAddress_Revert() public {
    EVM2EVMMultiOffRamp.DynamicConfig memory dynamicConfig = generateDynamicMultiOffRampConfig(ZERO_ADDRESS, ZERO_ADDRESS);

    vm.expectRevert(EVM2EVMMultiOffRamp.ZeroAddressNotAllowed.selector);

    s_offRamp.setOCR2Config(
      s_valid_signers, s_valid_transmitters, s_f, abi.encode(dynamicConfig), s_offchainConfigVersion, abi.encode("")
    );
  }
}

contract EVM2EVMMultiOffRamp_metadataHash is EVM2EVMMultiOffRampSetup {
  function test_MetadataHash_Success() public view {
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
      )
    );
  }

  function test_MetadataHashChangesOnSourceChain_Success() public view {
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR + 1, ON_RAMP_ADDRESS);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR + 1, DEST_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
      )
    );
    assertTrue(
      h != s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
    );
  }

  function test_MetadataHashChangesOnOnRampAddress_Success() public view {
    address mockOnRampAddress = address(uint160(ON_RAMP_ADDRESS) + 1);
    bytes32 h = s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, mockOnRampAddress);
    assertEq(
      h,
      keccak256(
        abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, mockOnRampAddress)
      )
    );
    assertTrue(
      h != s_offRamp.metadataHash(SOURCE_CHAIN_SELECTOR, ON_RAMP_ADDRESS)
    );
  }
}
