pragma solidity ^0.8.24;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

import {EnumerableSet} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

import {Router} from "../../Router.sol";

abstract contract USDCBridgeMigrator is OwnerIsCreator {
  using EnumerableSet for EnumerableSet.UintSet;

  address internal s_circleUSDCMigrator;

  mapping(uint64 chainSelector => uint256 lockedBalance) internal s_lockedTokensByChainSelector;

  EnumerableSet.UintSet internal s_ExecutedCCTPChainMigrations;

  uint64 internal s_proposedUSDCMigrationChain;

  event CCTPMigrationProposed(uint64 remoteChainSelector);
  event CCTPMigrationExecuted(uint64 remoteChainSelector, uint256 USDCBurned);
  event CCTPMigrationCancelled(uint64 existingProposalSelector);

  error onlyCircle();
  error ExistingMigrationProposal();
  error NoExistingMigrationProposal();
  error NoMigrationProposalPending();
  error InvalidChainSelector(uint64 remoteChainSelector);

  IBurnMintERC20 internal immutable i_USDC;
  Router internal immutable i_router;

  constructor(address token, address router) {
    i_USDC = IBurnMintERC20(token);
    i_router = Router(router);
  }

  function burnLockedUSDC() public {
    if (msg.sender != s_circleUSDCMigrator) revert onlyCircle();
    if (s_proposedUSDCMigrationChain == 0) revert ExistingMigrationProposal();

    uint64 burnChainSelector = s_proposedUSDCMigrationChain;
    uint256 tokensToBurn = s_lockedTokensByChainSelector[burnChainSelector];

    // Even though USDC is a trusted call, ensure CEI by updating state first
    delete s_lockedTokensByChainSelector[burnChainSelector];
    delete s_proposedUSDCMigrationChain;

    s_ExecutedCCTPChainMigrations.add(burnChainSelector);

    i_USDC.burn(tokensToBurn);

    // TODO: Should the "should-use-alt-mechanism" mapping be updated automatically here for the chain
    // selector on execution, or require another manual update?

    emit CCTPMigrationExecuted(burnChainSelector, tokensToBurn);
  }

  function proposeCCTPMigration(uint64 remoteChainSelector) external onlyOwner {
    // Prevent overwriting existing migration proposals until the current one is finished
    if (s_proposedUSDCMigrationChain != 0) revert ExistingMigrationProposal();

    // Ensure that the chain is supported by CCIP and non-zero, hasn't already been executed on, and is
    // a valid CCIP-supported chain selector
    if (
      remoteChainSelector == 0 || s_ExecutedCCTPChainMigrations.contains(remoteChainSelector)
        || !i_router.isChainSupported(remoteChainSelector)
    ) revert InvalidChainSelector(remoteChainSelector);

    s_proposedUSDCMigrationChain = remoteChainSelector;

    emit CCTPMigrationProposed(remoteChainSelector);
  }

  /// @notice Cancel an existing proposal to migrate a lane to CCTP.
  function cancelExistingCCTPMigrationProposal() external onlyOwner {
    if (s_proposedUSDCMigrationChain == 0) revert NoExistingMigrationProposal();

    uint64 currentProposalChainSelector = s_proposedUSDCMigrationChain;
    delete s_proposedUSDCMigrationChain;

    emit CCTPMigrationCancelled(currentProposalChainSelector);
  }

  function getCurrentProposedCCTPChainMigration() public view returns (uint64) {
    return s_proposedUSDCMigrationChain;
  }

  /// @dev The function should only be invoked once the address has been confirmed by Circle prior to
  /// chain expansion.
  function setCircleMigratorAddress(address migrator) external onlyOwner {
    s_circleUSDCMigrator = migrator;
  }

  function getLockedTokensForChain(uint64 remoteChainSelector) public view returns (uint256) {
    return s_lockedTokensByChainSelector[remoteChainSelector];
  }
}
