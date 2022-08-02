// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import "../helpers/BaseOffRampHelper.sol";
import "../mocks/MockBlobVerifier.sol";
import "../../pools/NativeTokenPool.sol";

contract BaseOffRampSetup is TokenSetup {
  event OffRampConfigSet(BaseOffRampInterface.OffRampConfig config);

  BaseOffRampHelper s_offRamp;
  BaseOffRampInterface.OffRampConfig s_offRampConfig;

  MockBlobVerifier s_mockBlobVerifier;
  NativeTokenPool s_nativePool;

  uint256 immutable POOL_BALANCE = 5000;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_offRampConfig = BaseOffRampInterface.OffRampConfig({
      executionDelaySeconds: 0,
      maxDataSize: 500,
      maxTokensLength: 5,
      permissionLessExecutionThresholdSeconds: 500
    });

    s_mockBlobVerifier = new MockBlobVerifier();

    PoolInterface.BucketConfig memory bucketConfig = PoolInterface.BucketConfig({rate: 1e16, capacity: 1e16});

    PoolInterface[] memory pools = new PoolInterface[](2);
    pools[0] = s_sourcePools[0];
    vm.warp(0);

    s_nativePool = new NativeTokenPool(s_sourceTokens[1], bucketConfig, bucketConfig);
    pools[1] = s_nativePool;
    vm.warp(BLOCK_TIME);

    s_offRamp = new BaseOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_offRampConfig,
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      s_sourceTokens,
      pools,
      2**10
    );

    s_nativePool.setOffRamp(s_offRamp, true);
    s_nativePool.getToken().transfer(address(s_nativePool), POOL_BALANCE);
  }

  function assertSameConfig(BaseOffRampInterface.OffRampConfig memory a, BaseOffRampInterface.OffRampConfig memory b)
    public
  {
    assertEq(a.executionDelaySeconds, b.executionDelaySeconds);
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.permissionLessExecutionThresholdSeconds, b.permissionLessExecutionThresholdSeconds);
  }
}

/// @notice #constructor
contract BaseOffRamp_constructor is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    // owner
    assertEq(OWNER, s_offRamp.owner());

    assertEq(address(s_mockBlobVerifier), address(s_offRamp.getBlobVerifier()));

    assertEq(SOURCE_CHAIN_ID, s_offRamp.SOURCE_CHAIN_ID());
    assertEq(DEST_CHAIN_ID, s_offRamp.CHAIN_ID());

    assertSameConfig(s_offRampConfig, s_offRamp.getConfig());
  }

  // Revert
  function testTokenConfigMismatchReverts() public {
    vm.expectRevert(TokenPoolRegistry.InvalidTokenPoolConfig.selector);

    PoolInterface[] memory pools = new PoolInterface[](1);

    IERC20[] memory wrongTokens = new IERC20[](5);
    s_offRamp = new BaseOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_offRampConfig,
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      wrongTokens,
      pools,
      2**10
    );
  }
}

/// @notice #getExecutionState
contract BaseOffRamp_getExecutionState is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    // setting the execution state is done with a helper function. This
    // is normally not exposed.
    s_offRamp.setExecutionState(1, CCIP.MessageExecutionState.FAILURE);
    s_offRamp.setExecutionState(10, CCIP.MessageExecutionState.IN_PROGRESS);
    s_offRamp.setExecutionState(33, CCIP.MessageExecutionState.UNTOUCHED);
    s_offRamp.setExecutionState(50, CCIP.MessageExecutionState.SUCCESS);

    assertEq(uint256(CCIP.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(1)));
    assertEq(uint256(CCIP.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(10)));
    assertEq(uint256(CCIP.MessageExecutionState.UNTOUCHED), uint256(s_offRamp.getExecutionState(33)));
    assertEq(uint256(CCIP.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(50)));
  }
}

/// @notice #getBlobVerifier
contract BaseOffRamp_getBlobVerifier is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    assertEq(address(s_mockBlobVerifier), address(s_offRamp.getBlobVerifier()));
  }
}

/// @notice #getConfig
contract BaseOffRamp_getConfig is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    assertSameConfig(s_offRampConfig, s_offRamp.getConfig());
  }
}

/// @notice #setConfig
contract BaseOffRamp_setConfig is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    BaseOffRampInterface.OffRampConfig memory newConfig = generateNewConfig();

    vm.expectEmit(false, false, false, true);
    emit OffRampConfigSet(newConfig);

    s_offRamp.setConfig(newConfig);

    assertSameConfig(newConfig, s_offRamp.getConfig());
  }

  // Revert
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_offRamp.setConfig(s_offRampConfig);
  }

  function generateNewConfig() internal pure returns (BaseOffRampInterface.OffRampConfig memory) {
    return
      BaseOffRampInterface.OffRampConfig({
        executionDelaySeconds: 20,
        maxDataSize: 1,
        maxTokensLength: 15,
        permissionLessExecutionThresholdSeconds: 200
      });
  }
}

/// @notice #_releaseOrMintToken internal function
contract BaseOffRamp__releaseOrMintToken is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    uint256 startingBalance = s_sourceTokens[1].balanceOf(OWNER);
    uint256 amount = POOL_BALANCE / 2;
    s_offRamp.releaseOrMintToken(s_sourceTokens[1], amount, OWNER);
    assertEq(startingBalance + amount, s_sourceTokens[1].balanceOf(OWNER));
  }

  // Revert
  function testExceedsPoolReverts() public {
    vm.expectRevert("ERC20: transfer amount exceeds balance");
    s_offRamp.releaseOrMintToken(s_sourceTokens[1], POOL_BALANCE * 2, OWNER);
  }

  function testUnsupportedTokenReverts() public {
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, s_destTokens[1]));
    s_offRamp.releaseOrMintToken(s_destTokens[1], POOL_BALANCE / 2, OWNER);
  }
}

/// @notice #_releaseOrMintTokens
contract BaseOffRamp__releaseOrMintTokens is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    uint256 startingBalance = s_sourceTokens[1].balanceOf(OWNER);

    IERC20[] memory tokens = new IERC20[](2);
    tokens[0] = s_sourceTokens[1];
    tokens[1] = s_sourceTokens[1];

    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 100;
    amounts[1] = 50;

    s_offRamp.releaseOrMintTokens(tokens, amounts, OWNER);
    assertEq(startingBalance + amounts[0] + amounts[1], s_sourceTokens[1].balanceOf(OWNER));
  }

  // Revert

  function testUnsupportedTokenReverts() public {
    uint256[] memory amounts = new uint256[](2);
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, s_destTokens[0]));
    s_offRamp.releaseOrMintTokens(s_destTokens, amounts, OWNER);
  }

  function testTokenAndAmountMisMatchReverts() public {
    uint256[] memory amounts = new uint256[](1);

    vm.expectRevert(BaseOffRampInterface.TokenAndAmountMisMatch.selector);
    s_offRamp.releaseOrMintTokens(s_destTokens, amounts, OWNER);
  }
}

/// @notice #_verifyMessages
contract BaseOffRamp__verifyMessages is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    bytes32[] memory mockBytes = new bytes32[](5);
    // Since we use a mock blob verifier it should always return 1
    (uint256 timestamp, ) = s_offRamp.verifyMessages(mockBytes, mockBytes, 1, mockBytes, 1);
    assertEq(1, timestamp);
  }
}

/// @notice #_getPool
contract BaseOffRamp__getPool is BaseOffRampSetup {
  // Success
  function testSuccess() public {
    address expectedPoolAddress = address(s_nativePool);
    address actualPoolAddress = address(s_offRamp.getPool(s_nativePool.getToken()));
    assertEq(expectedPoolAddress, actualPoolAddress);
  }

  // Revert
  function testUnsupportedTokenReverts() public {
    IERC20 wrongToken = IERC20(address(1));

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, wrongToken));
    s_offRamp.getPool_helper(wrongToken);
  }
}
