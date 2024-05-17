// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MultiOCR3Base} from "../../ocr/MultiOCR3Base.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MultiOCR3Helper} from "../helpers/MultiOCR3Helper.sol";

import {Vm} from "forge-std/Vm.sol";

contract MultiOCR3BaseSetup is BaseTest {
  // Signer private keys used for these test
  uint256 internal constant PRIVATE0 = 0x7b2e97fe057e6de99d6872a2ef2abf52c9b4469bc848c2465ac3fcd8d336e81d;
  uint256 internal constant PRIVATE1 = 0xab56160806b05ef1796789248e1d7f34a6465c5280899159d645218cd216cee6;
  uint256 internal constant PRIVATE2 = 0x6ec7caa8406a49b76736602810e0a2871959fbbb675e23a8590839e4717f1f7f;
  uint256 internal constant PRIVATE3 = 0x80f14b11da94ae7f29d9a7713ea13dc838e31960a5c0f2baf45ed458947b730a;

  address[] internal s_valid_signers;
  address[] internal s_valid_transmitters;

  //   bytes32[] internal s_rs;
  //   bytes32[] internal s_ss;
  //   bytes32 internal s_rawVs;

  bytes internal constant REPORT = abi.encode("testReport");
  MultiOCR3Helper internal s_multiOCR3;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_valid_transmitters = new address[](4);
    for (uint160 i = 0; i < 4; ++i) {
      s_valid_transmitters[i] = address(4 + i);
    }

    s_valid_signers = new address[](4);
    s_valid_signers[0] = vm.addr(PRIVATE0); //0xc110458BE52CaA6bB68E66969C3218A4D9Db0211
    s_valid_signers[1] = vm.addr(PRIVATE1); //0xc110a19c08f1da7F5FfB281dc93630923F8E3719
    s_valid_signers[2] = vm.addr(PRIVATE2); //0xc110fdF6e8fD679C7Cc11602d1cd829211A18e9b
    s_valid_signers[3] = vm.addr(PRIVATE3); //0xc11028017c9b445B6bF8aE7da951B5cC28B326C0

    s_multiOCR3 = new MultiOCR3Helper();

    // bytes32 testReportDigest = getTestReportDigest();

    // bytes32[] memory rs = new bytes32[](2);
    // bytes32[] memory ss = new bytes32[](2);
    // uint8[] memory vs = new uint8[](2);

    // // Calculate signatures
    // (vs[0], rs[0], ss[0]) = vm.sign(PRIVATE0, testReportDigest);
    // (vs[1], rs[1], ss[1]) = vm.sign(PRIVATE1, testReportDigest);

    // s_rs = rs;
    // s_ss = ss;
    // s_rawVs = bytes32(bytes1(vs[0] - 27)) | (bytes32(bytes1(vs[1] - 27)) >> 8);
  }

  function _getBasicConfigDigest(
    uint8 F,
    address[] memory signers,
    address[] memory transmitters
  ) internal view returns (bytes32) {
    bytes memory configBytes = abi.encode("");
    uint256 configVersion = 1;

    uint256 h = uint256(
      keccak256(
        abi.encode(
          block.chainid, address(s_multiOCR3), signers, transmitters, F, configBytes, configVersion, configBytes
        )
      )
    );
    uint256 prefixMask = type(uint256).max << (256 - 16); // 0xFFFF00..00
    uint256 prefix = 0x0001 << (256 - 16); // 0x000100..00
    return bytes32((prefix & prefixMask) | (h & ~prefixMask));
  }

  function _getTestReportDigest(
    uint8 F,
    address[] memory signers,
    address[] memory transmitters
  ) internal view returns (bytes32) {
    bytes32 configDigest = _getBasicConfigDigest(F, signers, transmitters);
    bytes32[3] memory reportContext = [configDigest, configDigest, configDigest];
    return keccak256(abi.encodePacked(keccak256(REPORT), reportContext));
  }

  function _assertOCRConfigEquality(
    MultiOCR3Base.OCRConfig memory configA,
    MultiOCR3Base.OCRConfig memory configB
  ) internal {
    vm.assertEq(configA.configInfo.configDigest, configB.configInfo.configDigest);
    vm.assertEq(configA.configInfo.F, configB.configInfo.F);
    vm.assertEq(configA.configInfo.n, configB.configInfo.n);
    vm.assertEq(configA.configInfo.uniqueReports, configB.configInfo.uniqueReports);
    vm.assertEq(configA.configInfo.isSignatureVerificationEnabled, configB.configInfo.isSignatureVerificationEnabled);

    vm.assertEq(configA.signers, configB.signers);
    vm.assertEq(configA.transmitters, configB.transmitters);
  }

  function _assertOCRConfigUnconfigured(MultiOCR3Base.OCRConfig memory config) internal {
    assertEq(config.configInfo.configDigest, bytes32());
    assertEq(config.signers.length, 0);
    assertEq(config.transmitters.length, 0);
  }
}

// contract OCR2Base_transmit is OCR2BaseSetup {
//   bytes32 internal s_configDigest;

//   function setUp() public virtual override {
//     OCR2BaseSetup.setUp();
//     bytes memory configBytes = abi.encode("");

//     s_configDigest = getBasicConfigDigest(s_f, 0);
//     s_OCR2Base.setOCR2Config(
//       s_valid_signers, s_valid_transmitters, s_f, configBytes, s_offchainConfigVersion, configBytes
//     );
//   }

//   function test_Transmit2SignersSuccess_gas() public {
//     vm.pauseGasMetering();
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];

//     vm.startPrank(s_valid_transmitters[0]);
//     vm.resumeGasMetering();
//     s_OCR2Base.transmit(reportContext, REPORT, s_rs, s_ss, s_rawVs);
//   }

//   // Reverts

//   function test_ForkedChain_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];

//     uint256 chain1 = block.chainid;
//     uint256 chain2 = chain1 + 1;
//     vm.chainId(chain2);
//     vm.expectRevert(abi.encodeWithSelector(OCR2Base.ForkedChain.selector, chain1, chain2));
//     vm.startPrank(s_valid_transmitters[0]);
//     s_OCR2Base.transmit(reportContext, REPORT, s_rs, s_ss, s_rawVs);
//   }

//   function test_WrongNumberOfSignatures_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];

//     vm.expectRevert(OCR2Base.WrongNumberOfSignatures.selector);
//     s_OCR2Base.transmit(reportContext, REPORT, new bytes32[](0), new bytes32[](0), s_rawVs);
//   }

//   function test_ConfigDigestMismatch_Revert() public {
//     bytes32 configDigest;
//     bytes32[3] memory reportContext = [configDigest, configDigest, configDigest];

//     vm.expectRevert(abi.encodeWithSelector(OCR2Base.ConfigDigestMismatch.selector, s_configDigest, configDigest));
//     s_OCR2Base.transmit(reportContext, REPORT, new bytes32[](0), new bytes32[](0), s_rawVs);
//   }

//   function test_SignatureOutOfRegistration_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];

//     bytes32[] memory rs = new bytes32[](2);
//     bytes32[] memory ss = new bytes32[](1);

//     vm.expectRevert(OCR2Base.SignaturesOutOfRegistration.selector);
//     s_OCR2Base.transmit(reportContext, REPORT, rs, ss, s_rawVs);
//   }

//   function test_UnAuthorizedTransmitter_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
//     bytes32[] memory rs = new bytes32[](2);
//     bytes32[] memory ss = new bytes32[](2);

//     vm.expectRevert(OCR2Base.UnauthorizedTransmitter.selector);
//     s_OCR2Base.transmit(reportContext, REPORT, rs, ss, s_rawVs);
//   }

//   function test_NonUniqueSignature_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
//     bytes32[] memory rs = s_rs;
//     bytes32[] memory ss = s_ss;

//     rs[1] = rs[0];
//     ss[1] = ss[0];
//     // Need to reset the rawVs to be valid
//     bytes32 rawVs = bytes32(bytes1(uint8(28) - 27)) | (bytes32(bytes1(uint8(28) - 27)) >> 8);

//     vm.startPrank(s_valid_transmitters[0]);
//     vm.expectRevert(OCR2Base.NonUniqueSignatures.selector);
//     s_OCR2Base.transmit(reportContext, REPORT, rs, ss, rawVs);
//   }

//   function test_UnauthorizedSigner_Revert() public {
//     bytes32[3] memory reportContext = [s_configDigest, s_configDigest, s_configDigest];
//     bytes32[] memory rs = new bytes32[](2);
//     rs[0] = s_configDigest;
//     bytes32[] memory ss = rs;

//     vm.startPrank(s_valid_transmitters[0]);
//     vm.expectRevert(OCR2Base.UnauthorizedSigner.selector);
//     s_OCR2Base.transmit(reportContext, REPORT, rs, ss, s_rawVs);
//   }
// }

contract MultiOCR3Base_setOCR3Configs is MultiOCR3BaseSetup {
  function test_SetConfigsZeroInput_Success() public {
    vm.recordLogs();
    s_multiOCR3.setOCR3Configs(new MultiOCR3Base.OCRConfigArgs[](0));

    // No logs emitted
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);
  }

  function test_SetConfigWithSigners_Success() public {
    uint8 F = 2;

    _assertOCRConfigUnconfigured(s_multiOCR3.latestConfigDetails(0));

    MultiOCR3Base.OCRConfigArgs[] memory ocrConfigs = new MultiOCR3Base.OCRConfigArgs[](1);
    ocrConfigs[0] = MultiOCR3Base.OCRConfigArgs({
      ocrPluginType: 0,
      configDigest: _getTestReportDigest(F, s_valid_signers, s_valid_transmitters),
      F: F,
      uniqueReports: false,
      isSignatureVerificationEnabled: true,
      signers: s_valid_signers,
      transmitters: s_valid_transmitters
    });

    vm.expectEmit();
    emit MultiOCR3Base.ConfigSet(
      ocrConfigs[0].ocrPluginType,
      ocrConfigs[0].configDigest,
      ocrConfigs[0].signers,
      ocrConfigs[0].transmitters,
      ocrConfigs[0].F
    );
    s_multiOCR3.setOCR3Configs(ocrConfigs);

    MultiOCR3Base.OCRConfig memory expectedConfig = MultiOCR3Base.OCRConfig({
      configInfo: MultiOCR3Base.ConfigInfo({
        configDigest: ocrConfigs[0].configDigest,
        F: ocrConfigs[0].F,
        n: s_valid_signers.length,
        uniqueReports: ocrConfigs[0].uniqueReports,
        isSignatureVerificationEnabled: ocrConfigs[0].isSignatureVerificationEnabled
      }),
      signers: s_valid_signers,
      transmitters: s_valid_transmitters
    });
    _assertOCRConfigEquality(s_multiOCR3.latestConfigDetails(0), expectedConfig);
  }

  function test_SetConfigWithoutSigners_Success() public {
    uint8 F = 2;
    address[] memory signers = new address[](0);

    _assertOCRConfigUnconfigured(s_multiOCR3.latestConfigDetails(0));

    MultiOCR3Base.OCRConfigArgs[] memory ocrConfigs = new MultiOCR3Base.OCRConfigArgs[](1);
    ocrConfigs[0] = MultiOCR3Base.OCRConfigArgs({
      ocrPluginType: 0,
      configDigest: _getTestReportDigest(F, signers, s_valid_transmitters),
      F: F,
      uniqueReports: false,
      isSignatureVerificationEnabled: false,
      signers: signers,
      transmitters: s_valid_transmitters
    });

    vm.expectEmit();
    emit MultiOCR3Base.ConfigSet(
      ocrConfigs[0].ocrPluginType,
      ocrConfigs[0].configDigest,
      ocrConfigs[0].signers,
      ocrConfigs[0].transmitters,
      ocrConfigs[0].F
    );
    s_multiOCR3.setOCR3Configs(ocrConfigs);

    MultiOCR3Base.OCRConfig memory expectedConfig = MultiOCR3Base.OCRConfig({
      configInfo: MultiOCR3Base.ConfigInfo({
        configDigest: ocrConfigs[0].configDigest,
        F: ocrConfigs[0].F,
        n: 0,
        uniqueReports: ocrConfigs[0].uniqueReports,
        isSignatureVerificationEnabled: ocrConfigs[0].isSignatureVerificationEnabled
      }),
      signers: signers,
      transmitters: s_valid_transmitters
    });
    _assertOCRConfigEquality(s_multiOCR3.latestConfigDetails(0), expectedConfig);
  }

  function test_SetMultipleConfigs_Success() public {
    // pluginType 2 remains unconfigured
    _assertOCRConfigUnconfigured(s_multiOCR3.latestConfigDetails(2));
  }

  //   function test_SetConfigSuccess_gas() public {
  //     vm.pauseGasMetering();
  //     bytes memory configBytes = abi.encode("");
  //     uint32 configCount = 0;

  //     bytes32 configDigest = getBasicConfigDigest(s_f, configCount++);

  //     address[] memory transmitters = s_OCR2Base.getTransmitters();
  //     assertEq(0, transmitters.length);

  //     vm.expectEmit();
  //     emit ConfigSet(
  //       0,
  //       configDigest,
  //       configCount,
  //       s_valid_signers,
  //       s_valid_transmitters,
  //       s_f,
  //       configBytes,
  //       s_offchainConfigVersion,
  //       configBytes
  //     );

  //     s_OCR2Base.setOCR2Config(
  //       s_valid_signers, s_valid_transmitters, s_f, configBytes, s_offchainConfigVersion, configBytes
  //     );

  //     transmitters = s_OCR2Base.getTransmitters();
  //     assertEq(s_valid_transmitters, transmitters);

  //     configDigest = getBasicConfigDigest(s_f, configCount++);

  //     vm.expectEmit();
  //     emit ConfigSet(
  //       uint32(block.number),
  //       configDigest,
  //       configCount,
  //       s_valid_signers,
  //       s_valid_transmitters,
  //       s_f,
  //       configBytes,
  //       s_offchainConfigVersion,
  //       configBytes
  //     );
  //     vm.resumeGasMetering();
  //     s_OCR2Base.setOCR2Config(
  //       s_valid_signers, s_valid_transmitters, s_f, configBytes, s_offchainConfigVersion, configBytes
  //     );
  //   }

  // Reverts
  // TODO: implement revert tests after re-introducing validations
  //   function test_RepeatAddress_Revert() public {
  //     address[] memory signers = new address[](10);
  //     signers[0] = address(1245678);
  //     address[] memory transmitters = new address[](10);
  //     transmitters[0] = signers[0];

  //     vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "repeated transmitter address"));
  //     s_OCR2Base.setOCR2Config(signers, transmitters, 2, abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_SingerCannotBeZeroAddress_Revert() public {
  //     uint256 f = 1;
  //     address[] memory signers = new address[](3 * f + 1);
  //     address[] memory transmitters = new address[](3 * f + 1);
  //     for (uint160 i = 0; i < 3 * f + 1; ++i) {
  //       signers[i] = address(i + 1);
  //       transmitters[i] = address(i + 1000);
  //     }

  //     signers[0] = address(0);

  //     vm.expectRevert(OCR2Base.OracleCannotBeZeroAddress.selector);
  //     s_OCR2Base.setOCR2Config(signers, transmitters, uint8(f), abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_TransmitterCannotBeZeroAddress_Revert() public {
  //     uint256 f = 1;
  //     address[] memory signers = new address[](3 * f + 1);
  //     address[] memory transmitters = new address[](3 * f + 1);
  //     for (uint160 i = 0; i < 3 * f + 1; ++i) {
  //       signers[i] = address(i + 1);
  //       transmitters[i] = address(i + 1000);
  //     }

  //     transmitters[0] = address(0);

  //     vm.expectRevert(OCR2Base.OracleCannotBeZeroAddress.selector);
  //     s_OCR2Base.setOCR2Config(signers, transmitters, uint8(f), abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_OracleOutOfRegister_Revert() public {
  //     address[] memory signers = new address[](10);
  //     address[] memory transmitters = new address[](0);

  //     vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "oracle addresses out of registration"));
  //     s_OCR2Base.setOCR2Config(signers, transmitters, 2, abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_FTooHigh_Revert() public {
  //     address[] memory signers = new address[](0);
  //     uint8 f = 1;

  //     vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "faulty-oracle f too high"));
  //     s_OCR2Base.setOCR2Config(signers, new address[](0), f, abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_FMustBePositive_Revert() public {
  //     uint8 f = 0;

  //     vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "f must be positive"));
  //     s_OCR2Base.setOCR2Config(new address[](0), new address[](0), f, abi.encode(""), 100, abi.encode(""));
  //   }

  //   function test_TooManySigners_Revert() public {
  //     address[] memory signers = new address[](32);

  //     vm.expectRevert(abi.encodeWithSelector(OCR2Base.InvalidConfig.selector, "too many signers"));
  //     s_OCR2Base.setOCR2Config(signers, new address[](0), 0, abi.encode(""), 100, abi.encode(""));
  //   }
}
