// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

import {ThirdPartyTokenPool} from "./ThirdPartyTokenPool.sol";

import {IERC165} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/introspection/IERC165.sol";

/// @notice This pool mints and burns a token.
/// @dev One token per pool. The token must grant this pool permission to call burn and mint.
contract ThirdPartyBurnMintTokenPool is ThirdPartyTokenPool, ITypeAndVersion {
  event Burned(
    address indexed caller,
    address indexed sender,
    bytes receiver,
    uint256 amount,
    uint64 destChainSelector
  );
  event Minted(
    address indexed caller,
    address indexed receiver,
    bytes sender,
    uint256 amount,
    uint64 sourceChainSelector
  );

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ThirdPartyBurnMintTokenPool 1.2.0";

  /// @dev The unique burn mint pool flag to signal through EIP 165.
  bytes4 private constant BURN_MINT_INTERFACE_ID = bytes4(keccak256("ThirdPartyBurnMintTokenPool"));

  constructor(IBurnMintERC20 token) ThirdPartyTokenPool(token) {}

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
    IBurnMintERC20(address(i_token)).burn(amount);
    emit Burned(msg.sender, originalSender, receiver, amount, destChainSelector);
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
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
    emit Minted(msg.sender, receiver, originalSender, amount, sourceChainSelector);
  }

  /// @notice returns the burn mint interface flag used for EIP165 identification.
  function getBurnMintInterfaceId() public pure returns (bytes4) {
    return BURN_MINT_INTERFACE_ID;
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == BURN_MINT_INTERFACE_ID || super.supportsInterface(interfaceId);
  }
}
