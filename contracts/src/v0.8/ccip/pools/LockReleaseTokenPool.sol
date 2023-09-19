// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
/// Because of lock/unlock requiring liquidity, this pool contract also has function to add and remove
/// liquidity. This allows for proper bookkeeping for both user and liquidity provider balances.
/// @dev One token per LockReleaseTokenPool.
contract LockReleaseTokenPool is TokenPool {
  using SafeERC20 for IERC20;

  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed provider, uint256 indexed amount);

  error InsufficientLiquidity();
  error WithdrawalTooHigh();
  error LiquidityNotAccepted();

  /// @dev The unique lock release pool flag to signal through EIP 165.
  bytes4 private constant LOCK_RELEASE_INTERFACE_ID = bytes4(keccak256("LockReleaseTokenPool"));

  /// @dev Whether or not the pool accepts liquidity.
  /// External liquidity is not required when there is one canonical token deployed to a chain,
  /// and CCIP is facilitating mint/burn on all the other chains, in which case the invariant
  /// balanceOf(pool) on home chain == sum(totalSupply(mint/burn "wrapped" token) on all remote chains) should always hold
  bool internal immutable i_acceptLiquidity;

  mapping(address provider => uint256 balance) internal s_liquidityProviderBalances;

  constructor(
    IERC20 token,
    address[] memory allowlist,
    address armProxy,
    bool acceptLiquidity
  ) TokenPool(token, allowlist, armProxy) {
    i_acceptLiquidity = acceptLiquidity;
  }

  /// @notice Locks the token in the pool
  /// @param amount Amount to lock
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function lockOrBurn(
    address originalSender,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override onlyOnRamp checkAllowList(originalSender) whenHealthy returns (bytes memory) {
    _consumeOnRampRateLimit(amount);
    emit Locked(msg.sender, amount);
    return "";
  }

  /// @notice Release tokens from the pool to the recipient
  /// @param receiver Recipient address
  /// @param amount Amount to release
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64,
    bytes memory
  ) external override onlyOffRamp whenHealthy {
    _consumeOffRampRateLimit(amount);
    getToken().safeTransfer(receiver, amount);
    emit Released(msg.sender, receiver, amount);
  }

  /// @notice returns the lock release interface flag used for EIP165 identification.
  function getLockReleaseInterfaceId() public pure returns (bytes4) {
    return LOCK_RELEASE_INTERFACE_ID;
  }

  // @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure override returns (bool) {
    return interfaceId == LOCK_RELEASE_INTERFACE_ID || super.supportsInterface(interfaceId);
  }

  /// @notice Gets the amount of provided liquidity for a given address.
  /// @param provider The address for which to get the balance.
  /// @return The current provided liquidity.
  function getProvidedLiquidity(address provider) external view returns (uint256) {
    return s_liquidityProviderBalances[provider];
  }

  /// @notice Checks if the pool can accept liquidity.
  /// @return true if the pool can accept liquidity, false otherwise.
  function canAcceptLiquidity() external view returns (bool) {
    return i_acceptLiquidity;
  }

  /// @notice Adds liquidity to the pool. The tokens should be approved first.
  /// @param amount The amount of liquidity to provide.
  function addLiquidity(uint256 amount) external {
    if (!i_acceptLiquidity) revert LiquidityNotAccepted();
    i_token.safeTransferFrom(msg.sender, address(this), amount);
    s_liquidityProviderBalances[msg.sender] += amount;
    emit LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param amount The amount of liquidity to remove.
  function removeLiquidity(uint256 amount) external {
    if (s_liquidityProviderBalances[msg.sender] < amount) revert WithdrawalTooHigh();
    if (i_token.balanceOf(address(this)) < amount) revert InsufficientLiquidity();
    s_liquidityProviderBalances[msg.sender] -= amount;
    i_token.safeTransfer(msg.sender, amount);
    emit LiquidityRemoved(msg.sender, amount);
  }
}
