// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../mocks/MockERC20.sol";
import "../helpers/SubscriptionManagerHelper.sol";
import "../BaseTest.t.sol";
import "../../subscription/Subscription.sol";

contract SubscriptionSetup is BaseTest {
  event SubscriptionFunded(address receiver, uint256 funding);
  event SubscriptionCreated(address receiver);
  event PreparedSetSenders(address receiver, address[] newSenders);
  event NewSendersSet(address receiver, address[] newSenders);
  event PreparedWithdrawalRequest(address receiver, uint256 amount);
  event WithdrawalProcessed(address receiver, uint256 amount);

  Subscription s_subscriptionContract;
  IERC20 s_feeToken;
  uint256 internal constant APPROVED_AMOUNT = 100;

  SubscriptionManagerInterface s_receiver;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_receiver = new SubscriptionManagerHelper(OWNER);
    s_feeToken = new MockERC20("sLINK", "sLNK", OWNER, 2**256 - 1);
    s_subscriptionContract = new Subscription(
      SubscriptionInterface.SubscriptionConfig({
        setSubscriptionSenderDelay: SET_SUBSCRIPTION_SENDER_DELAY,
        withdrawalDelay: WITHDRAWAL_DELAY,
        feeToken: s_feeToken
      })
    );

    s_feeToken.approve(address(s_subscriptionContract), APPROVED_AMOUNT * 2);
    s_subscriptionContract.createSubscription(generateSubscriptionForOwner());
  }

  function generateSubscriptionForOwner() internal view returns (SubscriptionInterface.OffRampSubscription memory) {
    address[] memory senders = new address[](1);
    senders[0] = OWNER;
    return
      SubscriptionInterface.OffRampSubscription({
        senders: senders,
        receiver: s_receiver,
        strictSequencing: true,
        balance: APPROVED_AMOUNT
      });
  }
}

/// @notice #getSupportedTokensForExecutionFee
contract Subscription_getSupportedTokensForExecutionFee is SubscriptionSetup {
  // Success
  function testSuccess() public {
    address[] memory supportedToken = s_subscriptionContract.getSupportedTokensForExecutionFee();
    assertEq(address(s_feeToken), address(supportedToken[0]));
    assertEq(1, supportedToken.length);
  }
}

/// @notice #getSubscription
contract Subscription_getSubscription is SubscriptionSetup {
  // Success
  function testSuccess() public {
    assertEq(s_subscriptionContract.getSubscription(address(s_receiver)).balance, APPROVED_AMOUNT);
  }
}

/// @notice #getFeeToken
contract Subscription_getFeeToken is SubscriptionSetup {
  // Success
  function testSuccess() public {
    assertEq(address(s_subscriptionContract.getFeeToken()), address(s_feeToken));
  }
}

/// @notice #createSubscription
contract Subscription_createSubscription is SubscriptionSetup {
  // Success
  function testSuccess() public {
    SubscriptionInterface.OffRampSubscription memory subscription = generateSubscriptionForOwner();

    subscription.receiver = new SubscriptionManagerHelper(OWNER);

    vm.expectEmit(false, false, false, true);
    emit SubscriptionCreated(address(subscription.receiver));

    s_subscriptionContract.createSubscription(subscription);
  }

  // Reverts
  function testInvalidManagerReverts() public {
    vm.stopPrank();

    vm.expectRevert(SubscriptionInterface.InvalidManager.selector);
    s_subscriptionContract.createSubscription(generateSubscriptionForOwner());
  }

  function testSubscriptionAlreadyExistsReverts() public {
    vm.expectRevert(SubscriptionInterface.SubscriptionAlreadyExists.selector);
    s_subscriptionContract.createSubscription(generateSubscriptionForOwner());
  }

  function testApproveTooLowReverts() public {
    SubscriptionInterface.OffRampSubscription memory subscription = generateSubscriptionForOwner();
    subscription.receiver = new SubscriptionManagerHelper(OWNER);
    subscription.balance = APPROVED_AMOUNT + 1;

    vm.expectRevert("ERC20: transfer amount exceeds allowance");
    s_subscriptionContract.createSubscription(subscription);
  }
}

/// @notice #fundSubscription
contract Subscription_fundSubscription is SubscriptionSetup {
  // Success
  function testSuccess() public {
    uint256 balanceBefore = s_feeToken.balanceOf(OWNER);
    vm.expectEmit(false, false, false, true);
    emit SubscriptionFunded(address(s_receiver), APPROVED_AMOUNT);

    s_subscriptionContract.fundSubscription(address(s_receiver), APPROVED_AMOUNT);
    assertEq(balanceBefore - APPROVED_AMOUNT, s_feeToken.balanceOf(OWNER));
  }

  // Reverts
  function testFundingAmountNotPositiveReverts() public {
    vm.expectRevert(SubscriptionInterface.FundingAmountNotPositive.selector);
    s_subscriptionContract.fundSubscription(address(s_receiver), 0);
  }
}

/// @notice #prepareSetSubscriptionSenders
contract Subscription_prepareSetSubscriptionSenders is SubscriptionSetup {
  // Success
  function testSuccess() public {
    address[] memory newSenders = new address[](5);
    newSenders[0] = address(100);

    vm.expectEmit(false, false, false, true);
    emit PreparedSetSenders(address(s_receiver), newSenders);

    s_subscriptionContract.prepareSetSubscriptionSenders(address(s_receiver), newSenders);
  }

  // Reverts

  function testInvalidManagerReverts() public {
    vm.stopPrank();
    address[] memory newSenders = new address[](5);

    vm.expectRevert(SubscriptionInterface.InvalidManager.selector);
    s_subscriptionContract.prepareSetSubscriptionSenders(address(s_receiver), newSenders);
  }
}

/// @notice #setSubscriptionSenders
contract Subscription_setSubscriptionSenders is SubscriptionSetup {
  // Success
  function testSuccess() public {
    address[] memory newSenders = new address[](5);
    newSenders[0] = address(100);
    s_subscriptionContract.prepareSetSubscriptionSenders(address(s_receiver), newSenders);
    vm.warp(BLOCK_TIME + SET_SUBSCRIPTION_SENDER_DELAY);

    vm.expectEmit(false, false, false, true);
    emit NewSendersSet(address(s_receiver), newSenders);

    s_subscriptionContract.setSubscriptionSenders(address(s_receiver), newSenders);

    SubscriptionInterface.OffRampSubscription memory updatedSubscription = s_subscriptionContract.getSubscription(
      address(s_receiver)
    );

    assertEq(updatedSubscription.senders[0], newSenders[0]);
  }

  // Reverts
  function testInvalidManagerReverts() public {
    vm.stopPrank();
    address[] memory newSenders = new address[](5);

    vm.expectRevert(SubscriptionInterface.InvalidManager.selector);
    s_subscriptionContract.setSubscriptionSenders(address(s_receiver), newSenders);
  }

  function testDelayNotPassedYetReverts() public {
    address[] memory newSenders = new address[](5);
    newSenders[0] = address(100);
    s_subscriptionContract.prepareSetSubscriptionSenders(address(s_receiver), newSenders);
    vm.expectRevert(
      abi.encodeWithSelector(
        SubscriptionInterface.DelayNotPassedYet.selector,
        BLOCK_TIME + SET_SUBSCRIPTION_SENDER_DELAY
      )
    );
    s_subscriptionContract.setSubscriptionSenders(address(s_receiver), newSenders);
  }

  function testAddressMismatchReverts() public {
    address[] memory newSenders = new address[](5);
    newSenders[0] = address(100);
    s_subscriptionContract.prepareSetSubscriptionSenders(address(s_receiver), newSenders);
    newSenders[0] = address(101);
    vm.warp(BLOCK_TIME + SET_SUBSCRIPTION_SENDER_DELAY);
    vm.expectRevert(abi.encodeWithSelector(SubscriptionInterface.AddressMismatch.selector, address(100), address(101)));
    s_subscriptionContract.setSubscriptionSenders(address(s_receiver), newSenders);
  }
}

/// @notice #prepareWithdrawal
contract Subscription_prepareWithdrawal is SubscriptionSetup {
  // Success
  function testSuccess() public {
    vm.expectEmit(false, false, false, true);
    emit PreparedWithdrawalRequest(address(s_receiver), APPROVED_AMOUNT);

    s_subscriptionContract.prepareWithdrawal(address(s_receiver), APPROVED_AMOUNT);
  }

  // Reverts
  function testBalanceTooLowReverts() public {
    vm.expectRevert(SubscriptionInterface.BalanceTooLow.selector);
    s_subscriptionContract.prepareWithdrawal(address(s_receiver), 2 * APPROVED_AMOUNT);
  }

  function testInvalidManagerReverts() public {
    vm.stopPrank();

    vm.expectRevert(SubscriptionInterface.InvalidManager.selector);
    s_subscriptionContract.prepareWithdrawal(address(s_receiver), APPROVED_AMOUNT);
  }
}

/// @notice #withdrawal
contract Subscription_withdrawal is SubscriptionSetup {
  // Success
  function testSuccess() public {
    uint256 balanceBefore = s_feeToken.balanceOf(OWNER);
    uint256 amount = APPROVED_AMOUNT;
    s_subscriptionContract.prepareWithdrawal(address(s_receiver), amount);
    vm.warp(BLOCK_TIME + WITHDRAWAL_DELAY);

    vm.expectEmit(false, false, false, true);
    emit WithdrawalProcessed(address(s_receiver), amount);

    s_subscriptionContract.withdrawal(address(s_receiver), amount);
    assertEq(balanceBefore + amount, s_feeToken.balanceOf(OWNER));
  }

  // Reverts
  function testInvalidManagerReverts() public {
    vm.stopPrank();

    vm.expectRevert(SubscriptionInterface.InvalidManager.selector);
    s_subscriptionContract.withdrawal(address(s_receiver), APPROVED_AMOUNT);
  }

  function testDelayNotPassedYetReverts() public {
    uint256 amount = APPROVED_AMOUNT;
    s_subscriptionContract.prepareWithdrawal(address(s_receiver), amount);
    vm.expectRevert(
      abi.encodeWithSelector(SubscriptionInterface.DelayNotPassedYet.selector, BLOCK_TIME + WITHDRAWAL_DELAY)
    );
    s_subscriptionContract.withdrawal(address(s_receiver), amount);
  }

  function testAmountMismatchReverts() public {
    uint256 amount = APPROVED_AMOUNT;
    s_subscriptionContract.prepareWithdrawal(address(s_receiver), amount);
    vm.warp(BLOCK_TIME + WITHDRAWAL_DELAY);
    amount = amount - 1;
    vm.expectRevert(abi.encodeWithSelector(SubscriptionInterface.AmountMismatch.selector, APPROVED_AMOUNT, amount));
    s_subscriptionContract.withdrawal(address(s_receiver), amount);
  }
}

/// @notice #getSubscriptionConfig
contract Subscription_getSubscriptionConfig is SubscriptionSetup {
  // Success
  function testSuccess() public {
    SubscriptionInterface.SubscriptionConfig memory subscriptionConfig = s_subscriptionContract.getSubscriptionConfig();

    assertEq(SET_SUBSCRIPTION_SENDER_DELAY, subscriptionConfig.setSubscriptionSenderDelay);
    assertEq(WITHDRAWAL_DELAY, subscriptionConfig.withdrawalDelay);
    assertEq(address(s_feeToken), address(subscriptionConfig.feeToken));
  }
}
