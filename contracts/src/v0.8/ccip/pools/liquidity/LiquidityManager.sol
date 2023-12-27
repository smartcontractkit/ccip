// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBridgeAdapter} from "./interfaces/IBridge.sol";
import {ILiquidityContainer} from "./interfaces/ILiquidityContainer.sol";
import {ILiquidityManager} from "./interfaces/ILiquidityManager.sol";

import {OCR3Base} from "../../ocr/OCR3Base.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract LiquidityManager is ILiquidityManager, OCR3Base {
  using SafeERC20 for IERC20;

  error ZeroAddress();
  error InvalidDestinationChain(uint64 chainSelector);
  error ZeroChainSelector();
  error InsufficientLiquidity(uint256 requested, uint256 available);

  event LiquidityTransferred(
    uint64 indexed fromChainSelector,
    uint64 indexed toChainSelector,
    address indexed to,
    uint256 amount
  );
  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed remover, uint256 indexed amount);

  struct CrossChainLiquidityManagerArgs {
    address destLiquidityManager;
    IBridgeAdapter bridge;
    address l2Token;
    uint64 destChainSelector;
    bool enabled;
  }

  struct CrossChainLiquidityManager {
    address destLiquidityManager;
    IBridgeAdapter bridge;
    address l2Token;
    bool enabled;
    // Potentially some fields related to the bridge
  }

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "LiquidityManager 1.3.0-dev";

  /// @notice The token that this pool manages liquidity for.
  IERC20 public immutable i_localToken;

  /// @notice The chain selector belonging to the chain this pool is deployed on.
  uint64 internal immutable i_localChainSelector;

  /// @notice Mapping of chain selector to liquidity container on other chains
  mapping(uint64 chainSelector => CrossChainLiquidityManager) private s_crossChainLiquidityContainers;

  /// @notice The liquidity container on the local chain
  /// @dev In the case of CCIP, this would be the token pool.
  ILiquidityContainer private s_localLiquidityContainer;

  constructor(IERC20 token, uint64 localChainSelector, ILiquidityContainer localLiquidityContainer) OCR3Base() {
    if (localChainSelector == 0) {
      revert ZeroChainSelector();
    }

    if (address(token) == address(0)) {
      revert ZeroAddress();
    }
    i_localToken = token;
    i_localChainSelector = localChainSelector;
    s_localLiquidityContainer = localLiquidityContainer;
  }

  // ================================================================
  // │                    Liquidity management                      │
  // ================================================================

  /// @inheritdoc ILiquidityManager
  function getLiquidity() public view returns (uint256 currentLiquidity) {
    return i_localToken.balanceOf(address(s_localLiquidityContainer));
  }

  function rebalanceLiquidity(uint64 chainSelector, uint256 amount) external onlyOwner {
    _rebalanceLiquidity(chainSelector, amount);
  }

  function _rebalanceLiquidity(uint64 chainSelector, uint256 amount) internal {
    uint256 currentBalance = getLiquidity();
    if (currentBalance < amount) {
      revert InsufficientLiquidity(amount, currentBalance);
    }

    CrossChainLiquidityManager memory destLiqManager = s_crossChainLiquidityContainers[chainSelector];

    if (!destLiqManager.enabled) {
      revert InvalidDestinationChain(chainSelector);
    }

    s_localLiquidityContainer.withdrawLiquidity(amount);
    i_localToken.approve(address(destLiqManager.bridge), amount);

    destLiqManager.bridge.sendERC20(
      address(i_localToken),
      destLiqManager.l2Token,
      destLiqManager.destLiquidityManager,
      amount
    );

    emit LiquidityTransferred(i_localChainSelector, chainSelector, destLiqManager.destLiquidityManager, amount);
  }

  function _receiveLiquidity(uint64 chainSelector, uint256 amount, bytes memory bridgeData) internal {
    // TODO
  }

  /// @notice Adds liquidity to the multi-chain system.
  /// @dev Anyone can call this function, but anyone other than the owner should regard
  /// adding liquidity as a donation to the system, as there is no way to get it out.
  /// This function is open to anyone to be able to quickly add funds to the system
  /// without having to go through potentially complicated multisig schemes to do it from
  /// the owner address.
  function addLiquidity(uint256 amount) external {
    i_localToken.safeTransferFrom(msg.sender, address(this), amount);

    // Make sure this is tether compatible, as they have strange approval requirements
    // Should be good since all approvals are always immediately used.
    i_localToken.approve(address(s_localLiquidityContainer), amount);
    s_localLiquidityContainer.provideLiquidity(amount);

    emit LiquidityAdded(msg.sender, amount);
  }

  function removeLiquidity(uint256 amount) external onlyOwner {
    uint256 currentBalance = i_localToken.balanceOf(address(s_localLiquidityContainer));
    if (currentBalance < amount) {
      revert InsufficientLiquidity(amount, currentBalance);
    }

    s_localLiquidityContainer.withdrawLiquidity(amount);
    i_localToken.safeTransfer(msg.sender, amount);

    emit LiquidityRemoved(msg.sender, amount);
  }

  function _report(bytes calldata report, uint64) internal override {
    ILiquidityManager.LiquidityInstructions memory instructions = abi.decode(
      report,
      (ILiquidityManager.LiquidityInstructions)
    );

    uint256 sendInstructions = instructions.sendLiquidityParams.length;
    for (uint256 i = 0; i < sendInstructions; ++i) {
      _rebalanceLiquidity(
        instructions.sendLiquidityParams[i].destChainSelector,
        instructions.sendLiquidityParams[i].amount
      );
    }

    uint256 receiveInstructions = instructions.receiveLiquidityParams.length;
    for (uint256 i = 0; i < receiveInstructions; ++i) {
      _receiveLiquidity(
        instructions.receiveLiquidityParams[i].sourceChainSelector,
        instructions.receiveLiquidityParams[i].amount,
        instructions.receiveLiquidityParams[i].bridgeData
      );
    }

    // todo emit?
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  function getCrossChainLiquidityManager(
    uint64 chainSelector
  ) external view returns (CrossChainLiquidityManager memory) {
    return s_crossChainLiquidityContainers[chainSelector];
  }

  function setCrossChainLiquidityManager(
    CrossChainLiquidityManagerArgs[] calldata crossChainLiquidityManagers
  ) external onlyOwner {
    for (uint256 i = 0; i < crossChainLiquidityManagers.length; ++i) {
      setCrossChainLiquidityManager(crossChainLiquidityManagers[i]);
    }
  }

  function setCrossChainLiquidityManager(
    CrossChainLiquidityManagerArgs calldata crossChainLiqManager
  ) public onlyOwner {
    if (crossChainLiqManager.destChainSelector == 0) {
      revert ZeroChainSelector();
    }

    if (crossChainLiqManager.destLiquidityManager == address(0) || address(crossChainLiqManager.bridge) == address(0)) {
      revert ZeroAddress();
    }

    s_crossChainLiquidityContainers[crossChainLiqManager.destChainSelector] = CrossChainLiquidityManager({
      destLiquidityManager: crossChainLiqManager.destLiquidityManager,
      bridge: crossChainLiqManager.bridge,
      l2Token: crossChainLiqManager.l2Token,
      enabled: crossChainLiqManager.enabled
    });
  }

  function getLocalLiquidityContainer() external view returns (address) {
    return address(s_localLiquidityContainer);
  }

  function setLocalLiquidityContainer(ILiquidityContainer localLiquidityContainer) external onlyOwner {
    s_localLiquidityContainer = localLiquidityContainer;
  }
}
