// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {BaseTest} from "../BaseTest.t.sol";
import {OCR2Base} from "../../ocr/OCR2Base.sol";
import {OCR2Helper} from "../helpers/OCR2Helper.sol";

contract OCR2BaseSetup is BaseTest {
  // Signer private keys used for these test
  // Private 0: 7b2e97fe057e6de99d6872a2ef2abf52c9b4469bc848c2465ac3fcd8d336e81d
  // Private 1: ab56160806b05ef1796789248e1d7f34a6465c5280899159d645218cd216cee6
  // Private 2: 6ec7caa8406a49b76736602810e0a2871959fbbb675e23a8590839e4717f1f7f
  // Private 3: 80f14b11da94ae7f29d9a7713ea13dc838e31960a5c0f2baf45ed458947b730a

  OCR2Helper s_OCR2Base;
  address[] s_valid_signers;
  address[] s_valid_transmitters;

  bytes32[] s_rs;
  bytes32[] s_ss;
  bytes32 constant s_rawVs = bytes32(uint256(1 << (256 - 8)));

  uint64 constant s_offchainConfigVersion = 3;
  uint8 constant s_f = 1;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_OCR2Base = new OCR2Helper();
    s_valid_signers = new address[](4);
    s_valid_transmitters = new address[](4);

    for (uint160 i = 0; i < 4; ++i) {
      s_valid_transmitters[i] = address(4 + i);
    }

    s_valid_signers[0] = 0xc110458BE52CaA6bB68E66969C3218A4D9Db0211;
    s_valid_signers[1] = 0xc110a19c08f1da7F5FfB281dc93630923F8E3719;
    s_valid_signers[2] = 0xc110fdF6e8fD679C7Cc11602d1cd829211A18e9b;
    s_valid_signers[3] = 0xc11028017c9b445B6bF8aE7da951B5cC28B326C0;

    bytes32[] memory rs = new bytes32[](3);
    bytes32[] memory ss = new bytes32[](3);

    rs[0] = 0x63b66c6cf62c3c79cb705c3d59d6c88a2116e496fcf40ce5247717768ea7ad45;
    ss[0] = 0x2fb845202d90ab3cfe5938e45b63d85217bf4df010a811744994cfa58e9a048f;

    rs[1] = 0x60f64b8490d0b3621f0b5258fb5dbe342cec2e0d31cbeb35169a52b7900657d7;
    ss[1] = 0x76719ca79072761168cfe0da7a0dcab92e97762c55e19f539a88bb9b4a0482b4;

    rs[2] = 0x7b34fcc75f34ee5072739531b8062c7ad3e72164d54484dea7a3b85a0b3dae29;
    ss[2] = 0x32086d9355be3832f0cc969153c8926e3677c7ebd4990acd2f16a2b75681bad0;

    s_rs = rs;
    s_ss = ss;
  }

  function getBasicConfigDigest(uint8 f, uint64 currentConfigCount) internal view returns (bytes32) {
    bytes memory configBytes = abi.encode("");
    return
      s_OCR2Base.configDigestFromConfigData(
        block.chainid,
        address(s_OCR2Base),
        currentConfigCount + 1,
        s_valid_signers,
        s_valid_transmitters,
        f,
        configBytes,
        s_offchainConfigVersion,
        configBytes
      );
  }
}

contract OCR2Base_transmit is OCR2BaseSetup {
  bytes32 s_configDigest;

  function setUp() public virtual override {
    OCR2BaseSetup.setUp();
    bytes memory configBytes = abi.encode("");

    s_configDigest = getBasicConfigDigest(s_f, 0);
    s_OCR2Base.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      configBytes,
      s_offchainConfigVersion,
      configBytes
    );
  }

  function testTransmit3SignersSuccess_gas() public {
    vm.pauseGasMetering();
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");

    changePrank(s_valid_transmitters[0]);
    vm.resumeGasMetering();
    s_OCR2Base.transmit(reportContext, report, s_rs, s_ss, s_rawVs);
  }

  // Reverts
  function testWrongNumberOfSignaturesReverts() public {
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");

    vm.expectRevert(OCR2Base.WrongNumberOfSignatures.selector);
    s_OCR2Base.transmit(reportContext, report, new bytes32[](0), new bytes32[](0), s_rawVs);
  }

  function testConfigDigestMismatchReverts() public {
    bytes32 configDigest;

    bytes32[3] memory reportContext = [configDigest, configDigest, configDigest];
    bytes memory report = abi.encode("testReport");

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.ConfigDigestMismatch.selector, s_configDigest, configDigest));
    s_OCR2Base.transmit(reportContext, report, new bytes32[](0), new bytes32[](0), s_rawVs);
  }

  function testSignatureOutOfRegistrationReverts() public {
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");

    bytes32[] memory rs = new bytes32[](3);
    bytes32[] memory ss = new bytes32[](2);

    vm.expectRevert(OCR2Base.SignaturesOutOfRegistration.selector);
    s_OCR2Base.transmit(reportContext, report, rs, ss, s_rawVs);
  }

  function testUnAuthorizedTransmitterReverts() public {
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");
    bytes32[] memory rs = new bytes32[](3);
    bytes32[] memory ss = new bytes32[](3);

    vm.expectRevert(OCR2Base.UnauthorizedTransmitter.selector);
    s_OCR2Base.transmit(reportContext, report, rs, ss, s_rawVs);
  }

  function testNonUniqueSignatureReverts() public {
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");
    bytes32[] memory rs = s_rs;
    bytes32[] memory ss = s_ss;

    rs[2] = rs[1];
    ss[2] = ss[1];

    changePrank(s_valid_transmitters[0]);
    vm.expectRevert(OCR2Base.NonUniqueSignatures.selector);
    s_OCR2Base.transmit(reportContext, report, rs, ss, s_rawVs);
  }

  function testUnauthorizedSignerReverts() public {
    bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
    bytes memory report = abi.encode("testReport");
    bytes32[] memory rs = new bytes32[](3);
    rs[0] = s_configDigest;
    bytes32[] memory ss = rs;

    changePrank(s_valid_transmitters[0]);
    vm.expectRevert(OCR2Base.UnauthorizedSigner.selector);
    s_OCR2Base.transmit(reportContext, report, rs, ss, s_rawVs);
  }
}

contract OCR2Base_setOCR2Config is OCR2BaseSetup {
  event ConfigSet(
    uint32 previousConfigBlockNumber,
    bytes32 configDigest,
    uint64 configCount,
    address[] signers,
    address[] transmitters,
    uint8 f,
    bytes onchainConfig,
    uint64 offchainConfigVersion,
    bytes offchainConfig
  );

  function testSetConfigSuccess_gas() public {
    vm.pauseGasMetering();
    bytes memory configBytes = abi.encode("");
    uint32 configCount = 0;

    bytes32 configDigest = getBasicConfigDigest(s_f, configCount++);

    address[] memory transmitters = s_OCR2Base.getTransmitters();
    assertEq(0, transmitters.length);

    vm.expectEmit();
    emit ConfigSet(
      0,
      configDigest,
      configCount,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      configBytes,
      s_offchainConfigVersion,
      configBytes
    );

    s_OCR2Base.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      configBytes,
      s_offchainConfigVersion,
      configBytes
    );

    transmitters = s_OCR2Base.getTransmitters();
    assertEq(s_valid_transmitters, transmitters);

    configDigest = getBasicConfigDigest(s_f, configCount++);

    vm.expectEmit();
    emit ConfigSet(
      uint32(block.number),
      configDigest,
      configCount,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      configBytes,
      s_offchainConfigVersion,
      configBytes
    );
    vm.resumeGasMetering();
    s_OCR2Base.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      configBytes,
      s_offchainConfigVersion,
      configBytes
    );
  }

  // Reverts
  function testRepeatAddressReverts() public {
    address[] memory signers = new address[](10);
    address[] memory oracles = new address[](10);

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "repeated transmitter address"));
    s_OCR2Base.setOCR2Config(signers, oracles, 2, abi.encode(""), 100, abi.encode(""));
  }

  function testOracleOutOfRegisterReverts() public {
    address[] memory signers = new address[](10);
    address[] memory transmitters = new address[](0);

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "oracle addresses out of registration"));
    s_OCR2Base.setOCR2Config(signers, transmitters, 2, abi.encode(""), 100, abi.encode(""));
  }

  function testFTooHighReverts() public {
    address[] memory signers = new address[](0);
    uint8 f = 1;

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "faulty-oracle f too high"));
    s_OCR2Base.setOCR2Config(signers, new address[](0), f, abi.encode(""), 100, abi.encode(""));
  }

  function testFMustBePositiveReverts() public {
    uint8 f = 0;

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "f must be positive"));
    s_OCR2Base.setOCR2Config(new address[](0), new address[](0), f, abi.encode(""), 100, abi.encode(""));
  }

  function testTooManySignersReverts() public {
    address[] memory signers = new address[](32);

    vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "too many signers"));
    s_OCR2Base.setOCR2Config(signers, new address[](0), 0, abi.encode(""), 100, abi.encode(""));
  }
}
