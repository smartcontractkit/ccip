// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../applications/GovernanceDapp.sol";
import "../../onRamp/subscription/EVM2EVMSubscriptionOnRampSetup.t.sol";

// setup
contract GovernanceDappSetup is EVM2EVMSubscriptionOnRampSetup {
  GovernanceDapp s_governanceDapp;
  IERC20 s_feeToken;

  GovernanceDapp.FeeConfig s_feeConfig;
  GovernanceDapp.CrossChainClone s_crossChainClone;

  Any2EVMOffRampRouterInterface internal s_receivingRouter;

  function setUp() public virtual override {
    EVM2EVMSubscriptionOnRampSetup.setUp();

    s_receivingRouter = Any2EVMOffRampRouterInterface(address(100));

    s_crossChainClone = GovernanceDapp.CrossChainClone({chainId: DEST_CHAIN_ID, contractAddress: address(1)});

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_governanceDapp = new GovernanceDapp(s_receivingRouter, s_onRampRouter, s_feeConfig);
    s_governanceDapp.addClone(s_crossChainClone);
  }
}

/// @notice #constructor
contract GovernanceDapp_constructor is GovernanceDappSetup {
  // Success
  function testSuccess() public {
    // typeAndVersion
    assertEq("GovernanceDapp 1.0.0", s_governanceDapp.typeAndVersion());
  }
}

/// @notice #voteForNewFeeConfig
contract GovernanceDapp_voteForNewFeeConfig is GovernanceDappSetup {
  event ConfigPropagated(uint256 chainId, address contractAddress);

  // Success
  function testSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({
      feeAmount: 10000,
      subscriptionManager: address(10),
      changedAtBlock: 100
    });
    bytes memory data = abi.encode(feeConfig);
    CCIP.EVM2EVMSubscriptionMessage memory subscriptionMsg = CCIP.EVM2EVMSubscriptionMessage({
      sequenceNumber: 1,
      sourceChainId: SOURCE_CHAIN_ID,
      sender: address(s_governanceDapp),
      receiver: s_crossChainClone.contractAddress,
      nonce: 1,
      data: data,
      tokensAndAmounts: new CCIP.EVMTokenAndAmount[](0),
      gasLimit: 3e5
    });

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(subscriptionMsg);

    vm.expectEmit(false, false, false, true);
    emit ConfigPropagated(s_crossChainClone.chainId, s_crossChainClone.contractAddress);

    s_governanceDapp.voteForNewFeeConfig(feeConfig);
  }
}

/// @notice #ccipReceive
contract GovernanceDapp_ccipReceive is GovernanceDappSetup {
  // Success

  function testSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({
      feeAmount: 10000,
      subscriptionManager: address(10),
      changedAtBlock: 100
    });

    CCIP.Any2EVMMessage memory message = CCIP.Any2EVMMessage({
      sourceChainId: SOURCE_CHAIN_ID,
      sender: abi.encode(OWNER),
      data: abi.encode(feeConfig),
      destTokensAndAmounts: new CCIP.EVMTokenAndAmount[](0)
    });

    changePrank(address(s_receivingRouter));

    s_governanceDapp.ccipReceive(message);

    GovernanceDapp.FeeConfig memory newConfig = s_governanceDapp.getFeeConfig();
    assertEq(feeConfig.subscriptionManager, newConfig.subscriptionManager);
    assertEq(feeConfig.changedAtBlock, newConfig.changedAtBlock);
    assertEq(feeConfig.feeAmount, newConfig.feeAmount);
  }
  // Revert
}
