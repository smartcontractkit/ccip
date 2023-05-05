// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../applications/GovernanceDapp.sol";
import "../onRamp/EVM2EVMOnRampSetup.t.sol";

// setup
contract GovernanceDappSetup is EVM2EVMOnRampSetup {
  GovernanceDapp s_governanceDapp;
  IERC20 s_feeToken;

  GovernanceDapp.FeeConfig s_feeConfig;
  GovernanceDapp.CrossChainClone s_crossChainClone;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_crossChainClone = GovernanceDapp.CrossChainClone({chainId: DEST_CHAIN_ID, contractAddress: address(1)});

    s_feeToken = IERC20(s_sourceTokens[0]);
    s_governanceDapp = new GovernanceDapp(address(s_sourceRouter), s_feeConfig, s_feeToken);
    s_governanceDapp.addClone(s_crossChainClone);

    uint256 fundingAmount = 1e18;

    // Fund the contract with LINK tokens
    s_feeToken.approve(address(s_governanceDapp), fundingAmount);
    s_governanceDapp.fund(fundingAmount);
  }
}

/// @notice #constructor
contract GovernanceDapp_constructor is GovernanceDappSetup {
  function testConstructorSuccess() public {
    // typeAndVersion
    assertEq("GovernanceDapp 1.0.0", s_governanceDapp.typeAndVersion());
  }
}

/// @notice #voteForNewFeeConfig
contract GovernanceDapp_voteForNewFeeConfig is GovernanceDappSetup {
  event ConfigPropagated(uint64 chainId, address contractAddress);

  function testVoteForNewFeeConfigSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({feeAmount: 10000, changedAtBlock: 100});
    bytes memory data = abi.encode(feeConfig);
    Internal.EVM2EVMMessage memory message = Internal.EVM2EVMMessage({
      sequenceNumber: 1,
      sourceChainId: SOURCE_CHAIN_ID,
      sender: address(s_governanceDapp),
      receiver: s_crossChainClone.contractAddress,
      nonce: 1,
      data: data,
      tokenAmounts: new Client.EVMTokenAmount[](0),
      gasLimit: 3e5,
      strict: false,
      feeToken: s_sourceFeeToken,
      feeTokenAmount: 64800216001, // todo
      messageId: ""
    });
    message.messageId = Internal._hash(message, s_metadataHash);

    vm.expectEmit();
    emit CCIPSendRequested(message);

    vm.expectEmit();
    emit ConfigPropagated(s_crossChainClone.chainId, s_crossChainClone.contractAddress);

    s_governanceDapp.voteForNewFeeConfig(feeConfig);
  }
}

/// @notice #ccipReceive
contract GovernanceDapp_ccipReceive is GovernanceDappSetup {
  function testCcipReceiveSuccess() public {
    GovernanceDapp.FeeConfig memory feeConfig = GovernanceDapp.FeeConfig({feeAmount: 10000, changedAtBlock: 100});

    Client.Any2EVMMessage memory message = Client.Any2EVMMessage({
      messageId: bytes32("a"),
      sourceChainSelector: SOURCE_CHAIN_ID,
      sender: abi.encode(OWNER),
      data: abi.encode(feeConfig),
      destTokenAmounts: new Client.EVMTokenAmount[](0)
    });

    changePrank(address(s_sourceRouter));

    s_governanceDapp.ccipReceive(message);

    GovernanceDapp.FeeConfig memory newConfig = s_governanceDapp.getFeeConfig();
    assertEq(feeConfig.changedAtBlock, newConfig.changedAtBlock);
    assertEq(feeConfig.feeAmount, newConfig.feeAmount);
  }
  // Revert
}
