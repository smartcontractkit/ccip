// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice This pool mints and burns LBTC tokens through the Cross Chain Transfer
/// Protocol (CCTP).
contract MockLBTCTokenPool is TokenPool, ITypeAndVersion {
    using SafeERC20 for IERC20;

    string public constant override typeAndVersion = "MockLBTCTokenPool 1.5.1";

    constructor(
        IERC20 token,
        address[] memory allowlist,
        address rmnProxy,
        address router
    ) TokenPool(token, 8, allowlist, rmnProxy, router) {

    }

    /// @notice Burn the token in the pool
    /// @dev emits ITokenMessenger.DepositForBurn
    /// @dev Assumes caller has validated destinationReceiver
    function lockOrBurn(
        Pool.LockOrBurnInV1 calldata lockOrBurnIn
    ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
        bytes memory payload;
        bytes memory destPoolData;
        payload = abi.encodePacked(hex"1234abcd");
        destPoolData = abi.encode(sha256(payload));

        IBurnMintERC20(address(i_token)).burn(lockOrBurnIn.amount);
        emit Burned(msg.sender, lockOrBurnIn.amount);

        return
            Pool.LockOrBurnOutV1({
            destTokenAddress: getRemoteToken(
                lockOrBurnIn.remoteChainSelector
            ),
            destPoolData: destPoolData
        });
    }

    function releaseOrMint(
        Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
    ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {

        // TODO: validate releaseOrMintIn.offchainTokenData?
        // Mint to the receiver
        IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, releaseOrMintIn.amount);

        emit Minted(
            msg.sender,
            releaseOrMintIn.receiver,
            releaseOrMintIn.amount
        );

        return
            Pool.ReleaseOrMintOutV1({
            destinationAmount: releaseOrMintIn.amount
        });
    }
}

