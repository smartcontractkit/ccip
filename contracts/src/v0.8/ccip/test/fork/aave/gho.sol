// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {Router} from "../../../Router.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {EVM2EVMOffRamp} from "../../../offRamp/EVM2EVMOffRamp.sol";
import {EVM2EVMOnRamp} from "../../../onRamp/EVM2EVMOnRamp.sol";
import {BurnMintTokenPoolAndProxy} from "../../../pools/BurnMintTokenPoolAndProxy.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {TokenPoolAndProxy} from "../../legacy/TokenPoolAndProxy.t.sol";

import {console2} from "forge-std/Console2.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {IERC20} from "forge-std/interfaces/IERC20.sol";

contract GHO is Test {
  uint256 private constant TOKENS_TO_SEND = 1;
  bytes32 internal constant TypeAndVersion1_5_OffRamp = keccak256("EVM2EVMOffRamp 1.5.0");

  struct ChainConfig {
    Router router;
    bool isMigrated;
    uint64 block;
    uint64 chainSelector;
    address gho;
    address newOnRamp;
    address newOffRamp;
    address proxyPool;
  }

  ChainConfig public SEPOLIA = ChainConfig({
    router: Router(0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59),
    isMigrated: true,
    block: 6937993,
    chainSelector: 16015286601757825753,
    gho: 0xc4bF5CbDaBE595361438F8c6a187bDc330539c60,
    newOnRamp: address(0),
    newOffRamp: address(0),
    proxyPool: address(0)
  });

  ChainConfig public ARBITRUM_SEPOLIA = ChainConfig({
    router: Router(0x2a9C5afB0d0e4BAb2BCdaE109EC4b0c4Be15a165),
    isMigrated: true,
    block: 91386151,
    chainSelector: 3478487238524512106,
    gho: 0xb13Cfa6f8B2Eed2C37fB00fF0c1A59807C585810,
    newOnRamp: address(0),
    newOffRamp: address(0),
    proxyPool: address(0)
  });

  ChainConfig public ETHEREUM = ChainConfig({
    router: Router(0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D),
    isMigrated: false,
    block: 21032248,
    chainSelector: 5009297550715157269,
    gho: 0x40D16FC0246aD3160Ccc09B8D0D3A2cD28aE6C2f,
    newOnRamp: 0x69eCC4E2D8ea56E2d0a05bF57f4Fd6aEE7f2c284,
    newOffRamp: 0xdf615eF8D4C64d0ED8Fd7824BBEd2f6a10245aC9,
    proxyPool: 0x9Ec9F9804733df96D1641666818eFb5198eC50f0
  });

  ChainConfig public ARBITRUM = ChainConfig({
    router: Router(0x141fa059441E0ca23ce184B6A78bafD2A517DdE8),
    isMigrated: false,
    block: 266996190,
    chainSelector: 4949039107694359620,
    gho: 0x7dfF72693f6A4149b17e7C6314655f6A9F7c8B33,
    newOnRamp: 0x67761742ac8A21Ec4D76CA18cbd701e5A6F3Bef3,
    newOffRamp: 0x91e46cc5590A4B9182e47f40006140A7077Dec31,
    proxyPool: 0x26329558f08cbb40d6a4CCA0E0C67b29D64A8c50
  });

  function test_gho_sepolia() public {
    uint256 sepoliaForkId = vm.createFork(vm.envString("SEPOLIA_RPC_URL"), SEPOLIA.block);
    uint256 arbitrumSepoliaForkId = vm.createFork(vm.envString("ARB_SEPOLIA_RPC_URL"), ARBITRUM_SEPOLIA.block);

    validateDirection(SEPOLIA, ARBITRUM_SEPOLIA, sepoliaForkId, arbitrumSepoliaForkId);
  }

  function test_gho_arbitrum_sep() public {
    uint256 sepoliaForkId = vm.createFork(vm.envString("SEPOLIA_RPC_URL"), SEPOLIA.block);
    uint256 arbitrumSepoliaForkId = vm.createFork(vm.envString("ARB_SEPOLIA_RPC_URL"), ARBITRUM_SEPOLIA.block);

    validateDirection(ARBITRUM_SEPOLIA, SEPOLIA, arbitrumSepoliaForkId, sepoliaForkId);
  }

  function test_gho_ethereum() public {
    uint256 ethereumForkId = vm.createFork(vm.envString("ETHEREUM_RPC_URL"), ETHEREUM.block);
    uint256 arbitrumSepoliaForkId = vm.createFork(vm.envString("ARBITRUM_RPC_URL"), ARBITRUM.block);

    validateDirection(ETHEREUM, ARBITRUM, ethereumForkId, arbitrumSepoliaForkId);
  }

  function test_gho_arbitrum() public {
    uint256 ethereumForkId = vm.createFork(vm.envString("ETHEREUM_RPC_URL"), ETHEREUM.block);
    uint256 arbitrumSepoliaForkId = vm.createFork(vm.envString("ARBITRUM_RPC_URL"), ARBITRUM.block);

    validateDirection(ARBITRUM, ETHEREUM, arbitrumSepoliaForkId, ethereumForkId);
  }

  function validateDirection(
    ChainConfig memory source,
    ChainConfig memory dest,
    uint256 sourceForkId,
    uint256 destForkId
  ) public {
    vm.selectFork(sourceForkId);
    vm.deal(address(this), 10_000 ether);

    if (!source.isMigrated) {
      // Succeeds pre-migration
      this.sendTokenMsg(source.router, source.gho, dest.chainSelector);
      console2.log("GHO message sent pre migration");

      _migrateChain(source, dest);
      vm.selectFork(destForkId);
      _migrateChain(dest, source);
      vm.selectFork(sourceForkId);
      console2.log("Chains migrated");

      // This fails as the current pools are not correctly set up.
      vm.expectRevert();
      this.sendTokenMsg(source.router, source.gho, dest.chainSelector);
      console2.log("GHO message failed as expected post migration");

      // TODO Actually migrate using the normal AAVE migration path
      // Will use the normal CCIP path here to unblock the test, simply replace this function
      // with the actual method and run the test again.
      _setProxyAsRouter(source, dest);

      // We do the same migration on the dest chain
      vm.selectFork(destForkId);
      _setProxyAsRouter(dest, source);
      vm.selectFork(sourceForkId);
      console2.log("Pools migrated");
    }

    // Pools have now been migrated, ready to be tested.
    Internal.EVM2EVMMessage memory ghoMsg = sendTokenMsg(source.router, source.gho, dest.chainSelector);
    console2.log("GHO message sent post migration");

    vm.selectFork(destForkId);
    _executeMsg(dest.router, ghoMsg);
    console2.log("GHO message executed post migration");
  }

  function sendTokenMsg(
    Router router,
    address token,
    uint64 destChainSelector
  ) public returns (Internal.EVM2EVMMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: TOKENS_TO_SEND});

    deal(token, address(this), TOKENS_TO_SEND);

    IERC20(token).approve(address(router), TOKENS_TO_SEND);

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(makeAddr("GHO_receiver")),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: address(0),
      extraArgs: ""
    });
    uint256 fee = router.getFee(destChainSelector, message);

    // `vm.getRecordedLogs` consumes the buffer, we do this to ensure we only get the latest event
    vm.getRecordedLogs();
    vm.recordLogs();
    router.ccipSend{value: fee}(destChainSelector, message);

    Vm.Log[] memory logs = vm.getRecordedLogs();
    for (uint256 i = 0; i < logs.length; ++i) {
      if (logs[i].topics[0] == EVM2EVMOnRamp.CCIPSendRequested.selector) {
        return abi.decode(logs[i].data, (Internal.EVM2EVMMessage));
      }
    }
    revert("No CCIPSendRequested event found");
  }

  // Emulates the migration to 1.5, the methods used are not necessarily representative of the actual migration.
  function _migrateChain(ChainConfig memory source, ChainConfig memory dest) internal {
    // Check if the token admin reg already supports the token
    TokenAdminRegistry tokenAdminRegistry =
      TokenAdminRegistry(EVM2EVMOnRamp(source.newOnRamp).getStaticConfig().tokenAdminRegistry);

    if (tokenAdminRegistry.getPool(source.gho) == address(0)) {
      address adminRegOwner = tokenAdminRegistry.owner();
      vm.startPrank(adminRegOwner);
      tokenAdminRegistry.proposeAdministrator(source.gho, adminRegOwner);
      tokenAdminRegistry.acceptAdminRole(source.gho);

      tokenAdminRegistry.setPool(source.gho, source.proxyPool);

      vm.stopPrank();
    }

    address poolOwner = TokenPool(source.proxyPool).owner();
    if (TokenPool(source.proxyPool).getRouter() != address(source.router)) {
      vm.prank(poolOwner);
      TokenPool(source.proxyPool).setRouter(address(source.router));
    }

    if (TokenPool(source.proxyPool).getRemotePool(dest.chainSelector).length == 0) {
      TokenPool.ChainUpdate[] memory chains = new TokenPool.ChainUpdate[](1);
      chains[0] = TokenPool.ChainUpdate({
        remoteChainSelector: dest.chainSelector,
        remotePoolAddress: abi.encode(dest.proxyPool),
        remoteTokenAddress: abi.encode(dest.gho),
        allowed: true,
        outboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0}),
        inboundRateLimiterConfig: RateLimiter.Config({isEnabled: false, capacity: 0, rate: 0})
      });
      vm.prank(poolOwner);
      TokenPool(source.proxyPool).applyChainUpdates(chains);
    }

    EVM2EVMOnRamp.DynamicConfig memory dynamicConfig = EVM2EVMOnRamp(source.newOnRamp).getDynamicConfig();
    if (dynamicConfig.router != address(source.router)) {
      dynamicConfig.router = address(source.router);

      vm.prank(EVM2EVMOnRamp(source.newOnRamp).owner());
      EVM2EVMOnRamp(source.newOnRamp).setDynamicConfig(dynamicConfig);
    }

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: dest.chainSelector, onRamp: source.newOnRamp});
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: dest.chainSelector, offRamp: source.newOffRamp});

    vm.prank(source.router.owner());
    source.router.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }

  function _setProxyAsRouter(ChainConfig memory source, ChainConfig memory dest) internal {
    EVM2EVMOnRamp onRamp = EVM2EVMOnRamp(source.router.getOnRamp(dest.chainSelector));
    EVM2EVMOnRamp.StaticConfig memory staticConfig = onRamp.getStaticConfig();
    TokenAdminRegistry tokenAdminRegistry = TokenAdminRegistry(staticConfig.tokenAdminRegistry);
    BurnMintTokenPoolAndProxy ghoProxyPool = BurnMintTokenPoolAndProxy(tokenAdminRegistry.getPool(source.gho));
    TokenPool nonProxyPool = TokenPool(ghoProxyPool.getPreviousPool());

    address ghoOwner = nonProxyPool.owner();
    vm.prank(ghoOwner);
    nonProxyPool.setRouter(address(ghoProxyPool));
  }

  function _executeMsg(Router router, Internal.EVM2EVMMessage memory message) internal {
    EVM2EVMOffRamp offRamp = _getOffRamp(router, message.sourceChainSelector);

    vm.prank(address(offRamp));

    // Empty gas overrides so it uses the values from the message.
    offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), new uint32[](1));
  }

  function _getOffRamp(Router router, uint64 sourceChainSelector) internal view returns (EVM2EVMOffRamp) {
    Router.OffRamp[] memory offRamps = router.getOffRamps();
    for (uint256 i = 0; i < offRamps.length; ++i) {
      Router.OffRamp memory configOffRamp = offRamps[i];
      if (configOffRamp.sourceChainSelector == sourceChainSelector) {
        EVM2EVMOffRamp offRamp = EVM2EVMOffRamp(configOffRamp.offRamp);
        if (keccak256(bytes(offRamp.typeAndVersion())) == TypeAndVersion1_5_OffRamp) {
          return offRamp;
        }
      }
    }

    revert("No offRamp found");
  }
}
