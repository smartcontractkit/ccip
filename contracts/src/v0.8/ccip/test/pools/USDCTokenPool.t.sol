// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {IBurnMintERC20} from "../../interfaces/pools/IBurnMintERC20.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {MockUSDC} from "../mocks/MockUSDC.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {Router} from "../../Router.sol";
import {USDCTokenPool} from "../../pools/USDC/USDCTokenPool.sol";

import {IERC165} from "../../../vendor/IERC165.sol";

contract USDCTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  USDCTokenPool internal s_usdcTokenPool;
  address s_routerAllowedOnRamp = address(3456);
  address s_routerAllowedOffRamp = address(234);
  Router s_router;

  MockUSDC internal s_mockUSDC;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new MockERC20("USDC", "USDC", OWNER, 2**256 - 1);
    setUpRamps();

    s_mockUSDC = new MockUSDC(42);

    USDCTokenPool.USDCConfig memory config = USDCTokenPool.USDCConfig({
      version: s_mockUSDC.messageBodyVersion(),
      tokenMessenger: address(s_mockUSDC),
      messageTransmitter: address(s_mockUSDC)
    });

    s_usdcTokenPool = new USDCTokenPool(config, IBurnMintERC20(address(s_token)), rateLimiterConfig());

    TokenPool.RampUpdate[] memory onRamps = new TokenPool.RampUpdate[](1);
    onRamps[0] = TokenPool.RampUpdate({ramp: s_routerAllowedOnRamp, allowed: true});

    TokenPool.RampUpdate[] memory offRamps = new TokenPool.RampUpdate[](1);
    offRamps[0] = TokenPool.RampUpdate({ramp: s_routerAllowedOffRamp, allowed: true});

    s_usdcTokenPool.applyRampUpdates(onRamps, offRamps);

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_ID,
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCaller")
    });

    s_usdcTokenPool.setDomains(domains);
  }

  function setUpRamps() internal {
    s_router = new Router(address(s_token));

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_ID, onRamp: s_routerAllowedOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    address[] memory offRamps = new address[](1);
    offRamps[0] = s_routerAllowedOffRamp;
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_ID, offRamp: offRamps[0]});

    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }
}

contract USDCTokenPool_lockOrBurn is USDCTokenPoolSetup {
  event DepositForBurn(
    uint64 indexed nonce,
    address indexed burnToken,
    uint256 amount,
    address indexed depositor,
    bytes32 mintRecipient,
    uint32 destinationDomain,
    bytes32 destinationTokenMessenger,
    bytes32 destinationCaller
  );
  event Burned(address indexed sender, uint256 amount);

  function testLockOrBurnSuccess(bytes32 destinationReceiver, uint256 amount) public {
    changePrank(s_routerAllowedOnRamp);
    deal(address(s_token), s_routerAllowedOnRamp, amount);
    s_token.approve(address(s_usdcTokenPool), amount);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_ID);

    vm.expectEmit();
    emit DepositForBurn(
      s_mockUSDC.s_nonce(),
      address(s_token),
      amount,
      address(s_usdcTokenPool),
      destinationReceiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.i_destinationTokenMessenger(),
      expectedDomain.allowedCaller
    );

    vm.expectEmit();
    emit Burned(s_routerAllowedOnRamp, amount);

    s_usdcTokenPool.lockOrBurn(OWNER, abi.encodePacked(destinationReceiver), amount, DEST_CHAIN_ID, bytes(""));
  }

  // Reverts
  function testUnknownDomainReverts() public {
    uint256 amount = 1000;
    changePrank(s_routerAllowedOnRamp);
    deal(address(s_token), s_routerAllowedOnRamp, amount);
    s_token.approve(address(s_usdcTokenPool), amount);

    uint64 wrongDomain = DEST_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.UnknownDomain.selector, wrongDomain));

    s_usdcTokenPool.lockOrBurn(OWNER, abi.encodePacked(address(0)), amount, wrongDomain, bytes(""));
  }

  function testPermissionsErrorReverts() public {
    vm.expectRevert(TokenPool.PermissionsError.selector);

    s_usdcTokenPool.lockOrBurn(OWNER, abi.encodePacked(address(0)), 0, DEST_CHAIN_ID, bytes(""));
  }
}

contract USDCTokenPool_releaseOrMint is USDCTokenPoolSetup {
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testReleaseOrMintSuccess(address receiver, uint256 amount) public {
    amount = bound(amount, 0, rateLimiterConfig().capacity);
    changePrank(s_routerAllowedOffRamp);

    bytes memory message = bytes("message bytes");
    bytes memory attestation = bytes("attestation bytes");

    bytes memory extraData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation})
    );

    vm.expectEmit();
    emit Minted(s_routerAllowedOffRamp, receiver, amount);

    vm.expectCall(address(s_mockUSDC), abi.encodeWithSelector(MockUSDC.receiveMessage.selector, message, attestation));

    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), receiver, amount, SOURCE_CHAIN_ID, extraData);
  }

  // Reverts
  function testUnlockingUSDCFailedReverts() public {
    changePrank(s_routerAllowedOffRamp);
    s_mockUSDC.setShouldSucceed(false);

    bytes memory extraData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: bytes(""), attestation: bytes("")})
    );

    vm.expectRevert(USDCTokenPool.UnlockingUSDCFailed.selector);

    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), OWNER, 1, SOURCE_CHAIN_ID, extraData);
  }

  function testConsumingMoreThanMaxCapacityReverts() public {
    uint256 capacity = rateLimiterConfig().capacity;
    uint256 amount = 10 * capacity;
    address receiver = address(1);
    changePrank(s_routerAllowedOffRamp);

    bytes memory extraData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: bytes(""), attestation: bytes("")})
    );

    vm.expectRevert(abi.encodeWithSelector(RateLimiter.ConsumingMoreThanMaxCapacity.selector, capacity, amount));

    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), receiver, amount, SOURCE_CHAIN_ID, extraData);
  }
}

contract USDCTokenPool_supportsInterface is USDCTokenPoolSetup {
  function testSupportsInterfaceSuccess() public {
    assertTrue(s_usdcTokenPool.supportsInterface(s_usdcTokenPool.getUSDCInterfaceId()));
    assertTrue(s_usdcTokenPool.supportsInterface(type(IPool).interfaceId));
    assertTrue(s_usdcTokenPool.supportsInterface(type(IERC165).interfaceId));
  }
}

contract USDCTokenPool_setDomains is USDCTokenPoolSetup {
  event DomainsSet(USDCTokenPool.DomainUpdate[]);

  mapping(uint64 => USDCTokenPool.Domain) private s_chainToDomain;

  function testSetDomainsSuccess(
    bytes32[10] calldata allowedCallers,
    uint32[10] calldata domainIdentifiers,
    uint64[10] calldata destChainSelectors
  ) public {
    uint256 numberOfDomains = allowedCallers.length;
    USDCTokenPool.DomainUpdate[] memory domainUpdates = new USDCTokenPool.DomainUpdate[](numberOfDomains);
    for (uint256 i = 0; i < numberOfDomains; ++i) {
      domainUpdates[i] = USDCTokenPool.DomainUpdate({
        allowedCaller: allowedCallers[i],
        domainIdentifier: domainIdentifiers[i],
        destChainSelector: destChainSelectors[i]
      });

      s_chainToDomain[destChainSelectors[i]] = USDCTokenPool.Domain({
        domainIdentifier: domainIdentifiers[i],
        allowedCaller: allowedCallers[i]
      });
    }

    vm.expectEmit();
    emit DomainsSet(domainUpdates);

    s_usdcTokenPool.setDomains(domainUpdates);

    for (uint256 i = 0; i < numberOfDomains; ++i) {
      USDCTokenPool.Domain memory expected = s_chainToDomain[destChainSelectors[i]];
      USDCTokenPool.Domain memory got = s_usdcTokenPool.getDomain(destChainSelectors[i]);
      assertEq(got.allowedCaller, expected.allowedCaller);
      assertEq(got.domainIdentifier, expected.domainIdentifier);
    }
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    USDCTokenPool.DomainUpdate[] memory domainUpdates = new USDCTokenPool.DomainUpdate[](0);

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");

    s_usdcTokenPool.setDomains(domainUpdates);
  }
}
