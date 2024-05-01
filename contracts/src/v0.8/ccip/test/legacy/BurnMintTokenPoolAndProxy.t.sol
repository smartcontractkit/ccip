// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IPoolPriorTo1_5} from "../../interfaces/IPoolPriorTo1_5.sol";

import {BurnMintERC677} from "../../../shared/token/ERC677/BurnMintERC677.sol";
import {PriceRegistry} from "../../PriceRegistry.sol";
import {Client} from "../../libraries/Client.sol";
import {BurnMintTokenPoolAndProxy} from "../../pools/BurnMintTokenPoolAndProxy.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {EVM2EVMOnRampHelper} from "../helpers/EVM2EVMOnRampHelper.sol";
import {EVM2EVMOnRampSetup} from "../onRamp/EVM2EVMOnRampSetup.t.sol";
import {BurnMintTokenPool1_2, TokenPool1_2} from "./BurnMintTokenPool1_2.sol";
import {BurnMintTokenPool1_4, TokenPool1_4} from "./BurnMintTokenPool1_4.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract BurnMintTokenPoolAndProxySetup is EVM2EVMOnRampSetup {
  BurnMintTokenPoolAndProxy internal s_newPool;
  IPoolPriorTo1_5 internal s_legacyPool;
  BurnMintERC677 internal s_token;

  function setUp() public virtual override {
    super.setUp();
    // Create a system with a token and a legacy pool
    s_token = new BurnMintERC677("Test", "TEST", 18, 0);
    deal(address(s_token), OWNER, 1e20);
  }

  function _migrateToSelfServe() internal {
    // Deploy the new pool
    s_newPool = new BurnMintTokenPoolAndProxy(s_token, new address[](0), address(s_mockARM), address(s_sourceRouter));
    // Set the previous pool on the new pool
    s_newPool.setPreviousPool(s_legacyPool);

    // Configure the lanes just like the legacy pool
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddress: abi.encode(makeAddr("dest_pool")),
      allowed: true,
      outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: getInboundRateLimiterConfig()
    });
    s_newPool.applyChainUpdates(chainUpdates);

    // Register the token on the token admin registry
    s_tokenAdminRegistry.registerAdministratorPermissioned(address(s_token), OWNER);
    // Set the pool on the admin registry
    s_tokenAdminRegistry.setPool(address(s_token), address(s_newPool));
  }

  function test_success1_4() public {
    _deployPool1_4();
    _migrateToSelfServe();

    // NOTE
    // when this call is made, the SENDING SIDE of old lanes stop working.
    // Set the Router of the old pool to the new pool
    BurnMintTokenPool1_4(address(s_legacyPool)).setRouter(address(s_newPool));

    // Approve enough for a few calls
    s_token.approve(address(s_sourceRouter), 1000);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), 1e18);

    // Everything is configured, we can now send a ccip tx
    _ccipSend();

    // Turn off the legacy pool, this enabled the 1.5 pool logic.
    s_newPool.setPreviousPool(IPoolPriorTo1_5(address(0)));

    // The new pool is now active, but is has not been given permissions to burn/mint yet
    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.SenderNotBurner.selector, address(s_newPool)));
    _ccipSend();

    // When we do give burn/mint, the new pool is fully active
    s_token.grantMintAndBurnRoles(address(s_newPool));
    _ccipSend();
  }

  function test_success1_2() public {
    _deployPool1_2();
    _migrateToSelfServe();

    TokenPool1_2.RampUpdate[] memory rampUpdates = new TokenPool1_2.RampUpdate[](1);
    rampUpdates[0] = TokenPool1_2.RampUpdate({
      ramp: address(s_newPool),
      allowed: true,
      rateLimiterConfig: getInboundRateLimiterConfig()
    });
    // Since this call doesn't impact the usability of the old pool, we can do it whenever we want
    BurnMintTokenPool1_2(address(s_legacyPool)).applyRampUpdates(rampUpdates, rampUpdates);

    // Approve enough for a few calls
    s_token.approve(address(s_sourceRouter), 1000);
    IERC20(s_sourceFeeToken).approve(address(s_sourceRouter), 1e18);

    // Everything is configured, we can now send a ccip tx
    _ccipSend();

    // Turn off the legacy pool, this enabled the 1.5 pool logic.
    s_newPool.setPreviousPool(IPoolPriorTo1_5(address(0)));

    // The new pool is now active, but is has not been given permissions to burn/mint yet
    vm.expectRevert(abi.encodeWithSelector(BurnMintERC677.SenderNotBurner.selector, address(s_newPool)));
    _ccipSend();

    // When we do give burn/mint, the new pool is fully active
    s_token.grantMintAndBurnRoles(address(s_newPool));
    _ccipSend();
  }

  function _deployPool1_4() internal {
    s_legacyPool = new BurnMintTokenPool1_4(s_token, new address[](0), address(s_mockARM), address(s_sourceRouter));
    s_token.grantMintAndBurnRoles(address(s_legacyPool));

    TokenPool1_4.ChainUpdate[] memory legacyChainUpdates = new TokenPool1_4.ChainUpdate[](1);
    legacyChainUpdates[0] = TokenPool1_4.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowed: true,
      outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: getInboundRateLimiterConfig()
    });
    BurnMintTokenPool1_4(address(s_legacyPool)).applyChainUpdates(legacyChainUpdates);
  }

  function _deployPool1_2() internal {
    s_legacyPool = new BurnMintTokenPool1_2(s_token, new address[](0), address(s_mockARM));
    s_token.grantMintAndBurnRoles(address(s_legacyPool));

    TokenPool1_2.RampUpdate[] memory rampUpdates = new TokenPool1_2.RampUpdate[](1);
    rampUpdates[0] = TokenPool1_2.RampUpdate({
      ramp: address(s_onRamp),
      allowed: true,
      rateLimiterConfig: getInboundRateLimiterConfig()
    });
    BurnMintTokenPool1_2(address(s_legacyPool)).applyRampUpdates(rampUpdates, rampUpdates);
  }

  function _ccipSend() internal {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: address(s_token), amount: 100});

    s_sourceRouter.ccipSend(
      DEST_CHAIN_SELECTOR,
      Client.EVM2AnyMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: s_sourceFeeToken,
        extraArgs: ""
      })
    );
  }
}
