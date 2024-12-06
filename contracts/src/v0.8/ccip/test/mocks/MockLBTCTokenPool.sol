// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";


contract MockLBTCTokenPool is TokenPool, ITypeAndVersion {
    using SafeERC20 for IERC20;

    string public constant override typeAndVersion = "LBTCTokenPool 1.5.1";

    constructor(
        IERC20 token,
        address[] memory allowlist,
        address rmnProxy,
        address router
    ) TokenPool(token, 6, allowlist, rmnProxy, router) {
    }

    /// @notice Burn the token in the pool
    /// @dev The _validateLockOrBurn check is an essential security check
    function lockOrBurn(
        Pool.LockOrBurnInV1 calldata lockOrBurnIn
    ) external virtual override returns (Pool.LockOrBurnOutV1 memory) {
        bytes memory payload;
        bytes memory destPoolData;
        payload = abi.encodePacked(hex"1234abcd");
        destPoolData = abi.encode(sha256(payload));

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
    ) external virtual override returns (Pool.ReleaseOrMintOutV1 memory) {

        // TODO: validate releaseOrMintIn.offchainTokenData?

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
