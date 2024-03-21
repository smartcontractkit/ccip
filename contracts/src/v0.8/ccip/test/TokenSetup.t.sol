// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IPool} from "../interfaces/pools/IPool.sol";

import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {LockReleaseTokenPool} from "../pools/LockReleaseTokenPool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {Client} from "../libraries/Client.sol";
import {BurnMintERC677} from "../../shared/token/ERC677/BurnMintERC677.sol";
import {MaybeRevertingBurnMintTokenPool} from "./helpers/MaybeRevertingBurnMintTokenPool.sol";
import {RouterSetup} from "./router/RouterSetup.t.sol";
import {TokenAdminRegistry} from "../pools/TokenAdminRegistry.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract TokenSetup is RouterSetup {
  address[] internal s_sourceTokens;
  address[] internal s_destTokens;

  address[] internal s_destPools;

  address internal s_sourceFeeToken;
  address internal s_destFeeToken;

  TokenAdminRegistry internal s_tokenAdminRegistry;

  mapping(address token => address sourcePool) internal s_sourcePoolByToken;
  mapping(address token => address destPool) internal s_destPoolByToken;

  function _deploySourceToken(string memory tokenName, uint256 dealAmount) internal returns (address) {
    BurnMintERC677 token = new BurnMintERC677(tokenName, tokenName, 18, 0);
    s_sourceTokens.push(address(token));
    deal(address(token), OWNER, dealAmount);
    return address(token);
  }

  function _deployDestToken(string memory tokenName, uint256 dealAmount) internal returns (address) {
    BurnMintERC677 token = new BurnMintERC677(tokenName, tokenName, 18, 0);
    s_destTokens.push(address(token));
    deal(address(token), OWNER, dealAmount);
    return address(token);
  }

  function _deployLockReleasePool(address token, bool isSourcePool) internal {
    LockReleaseTokenPool pool = new LockReleaseTokenPool(
      IERC20(token),
      new address[](0),
      address(s_mockARM),
      true,
      address(s_sourceRouter)
    );

    if (isSourcePool) {
      s_sourcePoolByToken[address(token)] = address(pool);
    } else {
      s_destPoolByToken[address(token)] = address(pool);
      s_destPools.push(address(pool));
    }
  }

  function _deployTokenAndBurnMintPool(address token, bool isSourcePool) internal {
    BurnMintTokenPool pool = new BurnMintTokenPool(
      BurnMintERC677(token),
      new address[](0),
      address(s_mockARM),
      address(s_sourceRouter)
    );
    BurnMintERC677(token).grantMintAndBurnRoles(address(pool));

    if (isSourcePool) {
      s_sourcePoolByToken[address(token)] = address(pool);
    } else {
      s_destPoolByToken[address(token)] = address(pool);
      s_destPools.push(address(pool));
    }
  }

  function setUp() public virtual override {
    RouterSetup.setUp();

    bool isSetup = s_sourceTokens.length != 0;
    if (isSetup) {
      return;
    }

    // Source tokens & pools
    address sourceLink = _deploySourceToken("sLINK", type(uint256).max);
    _deployLockReleasePool(sourceLink, true);
    s_sourceFeeToken = sourceLink;

    address sourceEth = _deploySourceToken("sETH", 2 ** 128);
    _deployTokenAndBurnMintPool(sourceEth, true);

    // Destination tokens & pools
    address destLink = _deployDestToken("dLINK", type(uint256).max);
    _deployLockReleasePool(destLink, false);
    s_destFeeToken = destLink;

    address destEth = _deployDestToken("dETH", 2 ** 128);
    _deployTokenAndBurnMintPool(destEth, false);

    // Float the dest link lock release pool with funds
    IERC20(destLink).transfer(s_destPoolByToken[destLink], POOL_BALANCE);

    s_tokenAdminRegistry = new TokenAdminRegistry();

    // Set pools in the registry
    for (uint256 i = 0; i < s_sourceTokens.length; ++i) {
      address token = s_sourceTokens[i];
      address pool = s_sourcePoolByToken[token];
      s_tokenAdminRegistry.registerAdministratorPermissioned(token, OWNER);
      s_tokenAdminRegistry.setPool(token, pool);

      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        remotePoolAddress: s_destPoolByToken[s_destTokens[i]],
        allowed: true,
        outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: getInboundRateLimiterConfig()
      });

      TokenPool(pool).applyChainUpdates(chainUpdates);
    }

    for (uint256 i = 0; i < s_destTokens.length; ++i) {
      address token = s_destTokens[i];
      address pool = s_destPoolByToken[token];
      s_tokenAdminRegistry.registerAdministratorPermissioned(token, OWNER);
      s_tokenAdminRegistry.setPool(token, pool);

      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
      chainUpdates[0] = TokenPool.ChainUpdate({
        remoteChainSelector: SOURCE_CHAIN_SELECTOR,
        remotePoolAddress: s_sourcePoolByToken[s_sourceTokens[i]],
        allowed: true,
        outboundRateLimiterConfig: getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: getInboundRateLimiterConfig()
      });

      TokenPool(pool).applyChainUpdates(chainUpdates);
    }
  }

  function getCastedSourceEVMTokenAmountsWithZeroAmounts()
    internal
    view
    returns (Client.EVMTokenAmount[] memory tokenAmounts)
  {
    tokenAmounts = new Client.EVMTokenAmount[](s_sourceTokens.length);
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      tokenAmounts[i].token = s_sourceTokens[i];
    }
  }
}
