// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {ThirdPartyTokenPool} from "./ThirdPartyTokenPool.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";
import {IERC165} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/introspection/IERC165.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool that locks and releases tokens on their native chain.
/// This pool does not accept liquidity. It is intended for tokens that are native on a single chain.
/// The token is locked and released on its native chain. On alternative chains, the token is burned and minted.
/// @dev One token per pool.
contract ThirdPartyLockReleaseTokenPool is ThirdPartyTokenPool, ITypeAndVersion {
  using SafeERC20 for IERC20;

  event TokensMigrated(address newPool, uint256 amount);

  event Locked(
    address indexed caller,
    address indexed sender,
    bytes receiver,
    uint256 amount,
    uint64 destChainSelector
  );
  event Released(
    address indexed caller,
    address indexed receiver,
    bytes sender,
    uint256 amount,
    uint64 sourceChainSelector
  );

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ThirdPartyLockReleaseTokenPool 1.2.0";

  /// @dev The unique lock release pool flag to signal through EIP 165.
  bytes4 private constant LOCK_RELEASE_INTERFACE_ID = bytes4(keccak256("ThirdPartyLockReleaseTokenPool"));

  constructor(IERC20 token) ThirdPartyTokenPool(token) {}

  /// @notice Burn the token in the pool.
  /// @param originalSender Original sender of the tokens.
  /// @param receiver Receiver of the tokens on destination chain.
  /// @param amount Amount to burn.
  /// @param destChainSelector Destination chain Id.
  function lockOrBurn(
    address originalSender,
    bytes calldata receiver,
    uint256 amount,
    uint64 destChainSelector,
    bytes calldata
  ) external virtual override onlyLockOrBurnCaller returns (bytes memory) {
    _consumeLockOrBurnRateLimit(destChainSelector, amount);
    emit Locked(msg.sender, originalSender, receiver, amount, destChainSelector);
    return "";
  }

  /// @notice Mint tokens from the pool to the recipient.
  /// @param originalSender Original sender of the tokens.
  /// @param receiver Receiver of the tokens.
  /// @param amount Amount to release or mint.
  /// @param sourceChainSelector Source chain Id.
  function releaseOrMint(
    bytes memory originalSender,
    address receiver,
    uint256 amount,
    uint64 sourceChainSelector,
    bytes memory
  ) external virtual override onlyReleaseOrMintCaller {
    _consumeReleaseOrMintRateLimit(sourceChainSelector, amount);
    getToken().safeTransfer(receiver, amount);
    emit Released(msg.sender, receiver, originalSender, amount, sourceChainSelector);
  }

  /// @notice Migrate all tokens locked in this pool to a new pool, used for upgrading pools.
  /// @param newPool Address of the pool to migrate to.
  function migrateTokensTo(address newPool) external onlyOwner {
    if (newPool == address(0)) revert ZeroAddressNotAllowed();

    uint256 balance = getToken().balanceOf(address(this));
    getToken().transfer(newPool, balance);

    emit TokensMigrated(newPool, balance);
  }

  /// @notice returns the lock release interface flag used for EIP165 identification.
  function getLockReleaseInterfaceId() public pure returns (bytes4) {
    return LOCK_RELEASE_INTERFACE_ID;
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == LOCK_RELEASE_INTERFACE_ID || super.supportsInterface(interfaceId);
  }
}
