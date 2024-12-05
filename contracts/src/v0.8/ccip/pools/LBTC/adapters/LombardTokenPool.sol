// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {IRouterClient} from "../../../interfaces/IRouterClient.sol";
import {Client} from "../../../libraries/Client.sol";
import {IBridge} from "../IBridge.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {TokenPool} from "../../TokenPool.sol";
import {CLAdapter} from "./CLAdapter.sol";

contract LombardTokenPool is TokenPool {
    CLAdapter public adapter;
    bool public isAttestationEnabled;

    /// @notice msg.sender gets the ownership of the contract given
    /// token pool implementation
    constructor(
        IERC20 lbtc_,
        address ccipRouter_,
        address[] memory allowlist_,
        address rmnProxy_,
        bool attestationEnable_
    ) TokenPool(lbtc_, 8, allowlist_, rmnProxy_, ccipRouter_) {
        isAttestationEnabled = attestationEnable_;
    }

    function setAdapter(CLAdapter adapter_) external {
        adapter = adapter_;
    }

    /// @notice Burn the token in the pool
    /// @dev The _validateLockOrBurn check is an essential security check
    function lockOrBurn(
        Pool.LockOrBurnInV1 calldata lockOrBurnIn
    ) external virtual override returns (Pool.LockOrBurnOutV1 memory) {
        _validateLockOrBurn(lockOrBurnIn);

        // send out to burn
        i_token.transfer(address(adapter), lockOrBurnIn.amount);
        (uint256 burnedAmount, bytes memory payload) = adapter.initiateDeposit(
            lockOrBurnIn.remoteChainSelector,
            lockOrBurnIn.receiver,
            lockOrBurnIn.amount
        );

        emit Burned(lockOrBurnIn.originalSender, burnedAmount);

        bytes memory destPoolData;
        if (isAttestationEnabled) {
            destPoolData = abi.encode(sha256(payload));
        } else {
            destPoolData = payload;
        }

        return
            Pool.LockOrBurnOutV1({
                destTokenAddress: getRemoteToken(
                    lockOrBurnIn.remoteChainSelector
                ),
                destPoolData: destPoolData
            });
    }

    /// @notice Mint tokens from the pool to the recipient
    /// @dev The _validateReleaseOrMint check is an essential security check
    function releaseOrMint(
        Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
    ) external virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
        _validateReleaseOrMint(releaseOrMintIn);

        if (isAttestationEnabled) {
            adapter.initiateWithdrawal(
                releaseOrMintIn.remoteChainSelector,
                releaseOrMintIn.offchainTokenData
            );
        } else {
            adapter.initWithdrawalNoSignatures(
                releaseOrMintIn.remoteChainSelector,
                releaseOrMintIn.sourcePoolData
            );
        }

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
