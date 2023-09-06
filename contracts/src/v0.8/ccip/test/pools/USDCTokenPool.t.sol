// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

import "../BaseTest.t.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {Router} from "../../Router.sol";
import {USDCTokenPool} from "../../pools/USDC/USDCTokenPool.sol";
import {BurnMintERC677} from "../../../shared/token/ERC677/BurnMintERC677.sol";
import {MockUSDC} from "../mocks/MockUSDC.sol";
import {USDCTokenPoolHelper} from "../helpers/USDCTokenPoolHelper.sol";

import {IERC165} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/introspection/IERC165.sol";

contract USDCTokenPoolSetup is BaseTest {
  IBurnMintERC20 internal s_token;
  MockUSDC internal s_mockUSDC;

  struct USDCMessage {
    uint32 version;
    uint32 sourceDomain;
    uint32 destinationDomain;
    uint64 nonce;
    bytes32 sender;
    bytes32 recipient;
    bytes32 destinationCaller;
    bytes messageBody;
  }

  uint32 internal constant SOURCE_DOMAIN_IDENTIFIER = 0x02020202;
  uint32 internal constant DEST_DOMAIN_IDENTIFIER = 0;

  bytes32 internal constant SOURCE_CHAIN_TOKEN_SENDER = bytes32(uint256(uint160(0x01111111221)));
  bytes32 internal constant SOURCE_CHAIN_USDC_TOKEN = bytes32(uint256(uint160(0x000000232355)));

  address internal s_routerAllowedOnRamp = address(3456);
  address internal s_routerAllowedOffRamp = address(234);
  Router internal s_router;

  USDCTokenPoolHelper internal s_usdcTokenPool;
  USDCTokenPoolHelper internal s_usdcTokenPoolWithAllowList;
  address[] internal s_allowedList;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new BurnMintERC677("LINK", "LNK", 18, 0);
    deal(address(s_token), OWNER, type(uint256).max);
    setUpRamps();

    s_mockUSDC = new MockUSDC(0);

    USDCTokenPool.USDCConfig memory config = USDCTokenPool.USDCConfig({
      version: s_mockUSDC.messageBodyVersion(),
      tokenMessenger: address(s_mockUSDC),
      messageTransmitter: address(s_mockUSDC)
    });

    s_usdcTokenPool = new USDCTokenPoolHelper(
      config,
      s_token,
      new address[](0),
      address(s_mockARM),
      DEST_DOMAIN_IDENTIFIER
    );

    s_allowedList.push(USER_1);
    s_usdcTokenPoolWithAllowList = new USDCTokenPoolHelper(
      config,
      s_token,
      s_allowedList,
      address(s_mockARM),
      DEST_DOMAIN_IDENTIFIER
    );

    TokenPool.RampUpdate[] memory onRamps = new TokenPool.RampUpdate[](1);
    onRamps[0] = TokenPool.RampUpdate({
      ramp: s_routerAllowedOnRamp,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });

    TokenPool.RampUpdate[] memory offRamps = new TokenPool.RampUpdate[](1);
    offRamps[0] = TokenPool.RampUpdate({
      ramp: s_routerAllowedOffRamp,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });

    s_usdcTokenPool.applyRampUpdates(onRamps, offRamps);
    s_usdcTokenPoolWithAllowList.applyRampUpdates(onRamps, offRamps);

    USDCTokenPool.DomainUpdate[] memory domains = new USDCTokenPool.DomainUpdate[](1);
    domains[0] = USDCTokenPool.DomainUpdate({
      destChainSelector: DEST_CHAIN_ID,
      domainIdentifier: 9999,
      allowedCaller: keccak256("allowedCaller"),
      enabled: true
    });

    s_usdcTokenPool.setDomains(domains);
    s_usdcTokenPoolWithAllowList.setDomains(domains);
  }

  function setUpRamps() internal {
    s_router = new Router(address(s_token), address(s_mockARM));

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_ID, onRamp: s_routerAllowedOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    address[] memory offRamps = new address[](1);
    offRamps[0] = s_routerAllowedOffRamp;
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_ID, offRamp: offRamps[0]});

    s_router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }

  function _generateUSDCMessage(USDCMessage memory usdcMessage) internal pure returns (bytes memory) {
    return
      abi.encodePacked(
        usdcMessage.version,
        usdcMessage.sourceDomain,
        usdcMessage.destinationDomain,
        usdcMessage.nonce,
        usdcMessage.sender,
        usdcMessage.recipient,
        usdcMessage.destinationCaller,
        usdcMessage.messageBody
      );
  }
}

contract USDCTokenPool_lockOrBurn is USDCTokenPoolSetup {
  error SenderNotAllowed(address sender);

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
  event TokensConsumed(uint256 tokens);

  function testFuzz_LockOrBurnSuccess(bytes32 destinationReceiver, uint256 amount) public {
    vm.assume(amount < rateLimiterConfig().capacity);
    vm.assume(amount > 0);
    changePrank(s_routerAllowedOnRamp);
    s_token.approve(address(s_usdcTokenPool), amount);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPool.getDomain(DEST_CHAIN_ID);

    vm.expectEmit();
    emit TokensConsumed(amount);

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

    bytes memory encodedNonce = s_usdcTokenPool.lockOrBurn(
      OWNER,
      abi.encodePacked(destinationReceiver),
      amount,
      DEST_CHAIN_ID,
      bytes("")
    );
    uint64 nonce = abi.decode(encodedNonce, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
  }

  function testFuzz_LockOrBurnWithAllowListSuccess(bytes32 destinationReceiver, uint256 amount) public {
    vm.assume(amount < rateLimiterConfig().capacity);
    vm.assume(amount > 0);
    changePrank(s_routerAllowedOnRamp);
    s_token.approve(address(s_usdcTokenPoolWithAllowList), amount);

    USDCTokenPool.Domain memory expectedDomain = s_usdcTokenPoolWithAllowList.getDomain(DEST_CHAIN_ID);

    vm.expectEmit();
    emit TokensConsumed(amount);
    vm.expectEmit();
    emit DepositForBurn(
      s_mockUSDC.s_nonce(),
      address(s_token),
      amount,
      address(s_usdcTokenPoolWithAllowList),
      destinationReceiver,
      expectedDomain.domainIdentifier,
      s_mockUSDC.i_destinationTokenMessenger(),
      expectedDomain.allowedCaller
    );
    vm.expectEmit();
    emit Burned(s_routerAllowedOnRamp, amount);

    bytes memory encodedNonce = s_usdcTokenPoolWithAllowList.lockOrBurn(
      s_allowedList[0],
      abi.encodePacked(destinationReceiver),
      amount,
      DEST_CHAIN_ID,
      bytes("")
    );
    uint64 nonce = abi.decode(encodedNonce, (uint64));
    assertEq(s_mockUSDC.s_nonce() - 1, nonce);
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

  function testLockOrBurnWithAllowListReverts() public {
    changePrank(s_routerAllowedOnRamp);

    vm.expectRevert(abi.encodeWithSelector(SenderNotAllowed.selector, STRANGER));

    s_usdcTokenPoolWithAllowList.lockOrBurn(STRANGER, abi.encodePacked(address(0)), 1000, DEST_CHAIN_ID, bytes(""));
  }
}

contract USDCTokenPool_releaseOrMint is USDCTokenPoolSetup {
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testFuzz_ReleaseOrMintSuccess(address recipient, uint256 amount) public {
    amount = bound(amount, 0, rateLimiterConfig().capacity);

    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: 0x060606060606,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(recipient))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    bytes memory message = _generateUSDCMessage(usdcMessage);
    bytes memory attestation = bytes("attestation bytes");

    bytes memory extraData = abi.encode(
      abi.encode(
        USDCTokenPool.SourceTokenDataPayload({nonce: usdcMessage.nonce, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})
      ),
      abi.encode(USDCTokenPool.MessageAndAttestation({message: message, attestation: attestation}))
    );

    vm.expectEmit();
    emit Minted(s_routerAllowedOffRamp, recipient, amount);

    vm.expectCall(address(s_mockUSDC), abi.encodeWithSelector(MockUSDC.receiveMessage.selector, message, attestation));

    changePrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), recipient, amount, SOURCE_CHAIN_ID, extraData);
  }

  // https://etherscan.io/tx/0xac9f501fe0b76df1f07a22e1db30929fd12524bc7068d74012dff948632f0883
  function testReleaseOrMintRealTxSuccess() public {
    bytes
      memory encodedUsdcMessage = hex"000000000000000300000000000000000000127a00000000000000000000000019330d10d9cc8751218eaf51e8885d058642e08a000000000000000000000000bd3fa81b58ba92a82136038b25adec7066af3155000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000004af08f56978be7dce2d1be3c65c005b41e79401c000000000000000000000000000000000000000000000000000000002057ff7a0000000000000000000000003a23f943181408eac424116af7b7790c94cb97a50000000000000000000000000000000000000000000000000000000000000000000000000000008274119237535fd659626b090f87e365ff89ebc7096bb32e8b0e85f155626b73ae7c4bb2485c184b7cc3cf7909045487890b104efb62ae74a73e32901bdcec91df1bb9ee08ccb014fcbcfe77b74d1263fd4e0b0e8de05d6c9a5913554364abfd5ea768b222f50c715908183905d74044bb2b97527c7e70ae7983c443a603557cac3b1c000000000000000000000000000000000000000000000000000000000000";
    bytes memory attestation = bytes("attestation bytes");

    uint32 nonce = 4730;
    uint32 sourceDomain = 3;

    bytes memory extraData = abi.encode(
      abi.encode(USDCTokenPool.SourceTokenDataPayload({nonce: nonce, sourceDomain: sourceDomain})),
      abi.encode(USDCTokenPool.MessageAndAttestation({message: encodedUsdcMessage, attestation: attestation}))
    );

    vm.expectCall(
      address(s_mockUSDC),
      abi.encodeWithSelector(MockUSDC.receiveMessage.selector, encodedUsdcMessage, attestation)
    );

    changePrank(s_routerAllowedOffRamp);
    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), OWNER, 100, SOURCE_CHAIN_ID, extraData);
  }

  // Reverts
  function testUnlockingUSDCFailedReverts() public {
    changePrank(s_routerAllowedOffRamp);
    s_mockUSDC.setShouldSucceed(false);

    uint256 amount = 13255235235;

    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: SOURCE_DOMAIN_IDENTIFIER,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: 0x060606060606,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(uint160(address(s_mockUSDC)))),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    bytes memory extraData = abi.encode(
      abi.encode(
        USDCTokenPool.SourceTokenDataPayload({nonce: usdcMessage.nonce, sourceDomain: SOURCE_DOMAIN_IDENTIFIER})
      ),
      abi.encode(
        USDCTokenPool.MessageAndAttestation({message: _generateUSDCMessage(usdcMessage), attestation: bytes("")})
      )
    );

    vm.expectRevert(USDCTokenPool.UnlockingUSDCFailed.selector);

    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), OWNER, amount, SOURCE_CHAIN_ID, extraData);
  }

  function testTokenMaxCapacityExceededReverts() public {
    uint256 capacity = rateLimiterConfig().capacity;
    uint256 amount = 10 * capacity;
    address recipient = address(1);
    changePrank(s_routerAllowedOffRamp);

    bytes memory extraData = abi.encode(
      USDCTokenPool.MessageAndAttestation({message: bytes(""), attestation: bytes("")})
    );

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.TokenMaxCapacityExceeded.selector, capacity, amount, address(s_token))
    );

    s_usdcTokenPool.releaseOrMint(abi.encode(OWNER), recipient, amount, SOURCE_CHAIN_ID, extraData);
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

  mapping(uint64 destChainSelector => USDCTokenPool.Domain domain) private s_chainToDomain;

  // Setting lower fuzz run as 256 runs was causing differing gas results in snapshot.
  /// forge-config: default.fuzz.runs = 32
  /// forge-config: ccip.fuzz.runs = 32
  function testFuzz_SetDomainsSuccess(
    bytes32[5] calldata allowedCallers,
    uint32[5] calldata domainIdentifiers,
    uint64[5] calldata destChainSelectors
  ) public {
    uint256 numberOfDomains = allowedCallers.length;
    USDCTokenPool.DomainUpdate[] memory domainUpdates = new USDCTokenPool.DomainUpdate[](numberOfDomains);
    for (uint256 i = 0; i < numberOfDomains; ++i) {
      vm.assume(allowedCallers[i] != bytes32(0) && domainIdentifiers[i] != 0 && destChainSelectors[i] != 0);

      domainUpdates[i] = USDCTokenPool.DomainUpdate({
        allowedCaller: allowedCallers[i],
        domainIdentifier: domainIdentifiers[i],
        destChainSelector: destChainSelectors[i],
        enabled: true
      });

      s_chainToDomain[destChainSelectors[i]] = USDCTokenPool.Domain({
        domainIdentifier: domainIdentifiers[i],
        allowedCaller: allowedCallers[i],
        enabled: true
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

  function testInvalidDomainReverts() public {
    bytes32 validCaller = bytes32(uint256(25));
    // Ensure valid domain works
    USDCTokenPool.DomainUpdate[] memory domainUpdates = new USDCTokenPool.DomainUpdate[](1);
    domainUpdates[0] = USDCTokenPool.DomainUpdate({
      allowedCaller: validCaller,
      domainIdentifier: 0, // ensures 0 is valid, as this is eth mainnet
      destChainSelector: 45690,
      enabled: true
    });

    s_usdcTokenPool.setDomains(domainUpdates);

    // Make update invalid on allowedCaller
    domainUpdates[0].allowedCaller = bytes32(0);
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidDomain.selector, domainUpdates[0]));

    s_usdcTokenPool.setDomains(domainUpdates);

    // Make valid again
    domainUpdates[0].allowedCaller = validCaller;

    // Make invalid on destChainSelector
    domainUpdates[0].destChainSelector = 0;
    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidDomain.selector, domainUpdates[0]));

    s_usdcTokenPool.setDomains(domainUpdates);
  }
}

contract USDCTokenPool_setConfig is USDCTokenPoolSetup {
  event ConfigSet(USDCTokenPool.USDCConfig);

  function testSetConfigSuccess() public {
    // Assert existing approval
    assertEq(type(uint256).max, s_usdcTokenPool.getToken().allowance(address(s_usdcTokenPool), address(s_mockUSDC)));

    MockUSDC newMockUSDC = new MockUSDC(0);

    USDCTokenPool.USDCConfig memory newConfig = USDCTokenPool.USDCConfig({
      version: 0,
      tokenMessenger: address(newMockUSDC),
      messageTransmitter: address(123456789)
    });

    USDCTokenPool.USDCConfig memory oldConfig = s_usdcTokenPool.getConfig();

    vm.expectEmit();
    emit ConfigSet(newConfig);
    s_usdcTokenPool.setConfig(newConfig);

    USDCTokenPool.USDCConfig memory gotConfig = s_usdcTokenPool.getConfig();
    assertEq(gotConfig.tokenMessenger, newConfig.tokenMessenger);
    assertEq(gotConfig.messageTransmitter, newConfig.messageTransmitter);
    assertEq(gotConfig.version, newConfig.version);

    assertEq(0, s_usdcTokenPool.getToken().allowance(address(s_usdcTokenPool), oldConfig.tokenMessenger));
    assertEq(
      type(uint256).max,
      s_usdcTokenPool.getToken().allowance(address(s_usdcTokenPool), gotConfig.tokenMessenger)
    );
    // Assert old approval is removed
    assertEq(0, s_usdcTokenPool.getToken().allowance(address(s_usdcTokenPool), address(s_mockUSDC)));
  }

  // Reverts

  function testInvalidMessageVersionReverts() public {
    USDCTokenPool.USDCConfig memory newConfig = USDCTokenPool.USDCConfig({
      version: 1,
      tokenMessenger: address(100),
      messageTransmitter: address(1)
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, newConfig.version));
    s_usdcTokenPool.setConfig(newConfig);
  }

  function testInvalidTokenMessengerVersionReverts() public {
    uint32 wrongVersion = 5;
    MockUSDC newMockUSDC = new MockUSDC(wrongVersion);

    USDCTokenPool.USDCConfig memory newConfig = USDCTokenPool.USDCConfig({
      version: 0,
      tokenMessenger: address(newMockUSDC),
      messageTransmitter: address(1)
    });

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidTokenMessengerVersion.selector, wrongVersion));
    s_usdcTokenPool.setConfig(newConfig);
  }

  function testInvalidConfigReverts() public {
    USDCTokenPool.USDCConfig memory newConfig = USDCTokenPool.USDCConfig({
      version: 0,
      tokenMessenger: address(0),
      messageTransmitter: address(123456789)
    });

    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    s_usdcTokenPool.setConfig(newConfig);

    newConfig.tokenMessenger = address(235);
    newConfig.messageTransmitter = address(0);

    vm.expectRevert(USDCTokenPool.InvalidConfig.selector);
    s_usdcTokenPool.setConfig(newConfig);
  }

  function testOnlyOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");

    s_usdcTokenPool.setConfig(
      USDCTokenPool.USDCConfig({version: 0, tokenMessenger: address(100), messageTransmitter: address(1)})
    );
  }
}

contract USDCTokenPool__validateMessage is USDCTokenPoolSetup {
  function testFuzz_ValidateMessageSuccess(uint32 sourceDomain, uint64 nonce) public {
    vm.pauseGasMetering();
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: sourceDomain,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: nonce,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(299999)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.resumeGasMetering();
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({nonce: nonce, sourceDomain: sourceDomain})
    );
  }

  // Reverts

  function testValidateInvalidMessageReverts() public {
    USDCMessage memory usdcMessage = USDCMessage({
      version: 0,
      sourceDomain: 1553252,
      destinationDomain: DEST_DOMAIN_IDENTIFIER,
      nonce: 387289284924,
      sender: SOURCE_CHAIN_TOKEN_SENDER,
      recipient: bytes32(uint256(92398429395823)),
      destinationCaller: bytes32(uint256(uint160(address(s_usdcTokenPool)))),
      messageBody: bytes("")
    });

    USDCTokenPool.SourceTokenDataPayload memory sourceTokenData = USDCTokenPool.SourceTokenDataPayload({
      nonce: usdcMessage.nonce,
      sourceDomain: usdcMessage.sourceDomain
    });

    bytes memory encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);

    uint32 expectedSourceDomain = usdcMessage.sourceDomain + 1;

    vm.expectRevert(
      abi.encodeWithSelector(USDCTokenPool.InvalidSourceDomain.selector, expectedSourceDomain, usdcMessage.sourceDomain)
    );
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({nonce: usdcMessage.nonce, sourceDomain: expectedSourceDomain})
    );

    uint64 expectedNonce = usdcMessage.nonce + 1;

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidNonce.selector, expectedNonce, usdcMessage.nonce));
    s_usdcTokenPool.validateMessage(
      encodedUsdcMessage,
      USDCTokenPool.SourceTokenDataPayload({nonce: expectedNonce, sourceDomain: usdcMessage.sourceDomain})
    );

    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER + 1;
    vm.expectRevert(
      abi.encodeWithSelector(
        USDCTokenPool.InvalidDestinationDomain.selector,
        DEST_DOMAIN_IDENTIFIER,
        usdcMessage.destinationDomain
      )
    );

    s_usdcTokenPool.validateMessage(
      _generateUSDCMessage(usdcMessage),
      USDCTokenPool.SourceTokenDataPayload({nonce: usdcMessage.nonce, sourceDomain: usdcMessage.sourceDomain})
    );
    usdcMessage.destinationDomain = DEST_DOMAIN_IDENTIFIER;

    uint32 wrongVersion = usdcMessage.version + 1;

    usdcMessage.version = wrongVersion;
    encodedUsdcMessage = _generateUSDCMessage(usdcMessage);

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPool.InvalidMessageVersion.selector, wrongVersion));
    s_usdcTokenPool.validateMessage(encodedUsdcMessage, sourceTokenData);
  }
}
