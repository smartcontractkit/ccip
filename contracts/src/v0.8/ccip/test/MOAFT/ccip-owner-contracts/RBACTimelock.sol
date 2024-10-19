// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @notice Contract module which acts as a timelocked controller with role-based
 * access control. When set as the owner of an `Ownable` smart contract, it
 * can enforce a timelock on `onlyOwner` maintenance operations and prevent
 * a list of blocked functions from being called. The timelock can be bypassed
 * by a bypasser or an admin in emergency situations that require quick action.
 *
 * Non-emergency actions are expected to follow the timelock.
 *
 * The contract has five roles. Each role can be inhabited by multiple
 * (potentially overlapping) addresses.
 *
 * 1) Admin: The admin manages membership for all roles (including the admin
 *    role itself). The admin automatically inhabits all other roles. The admin
 *    can call the bypasserExecuteBatch function to bypass any restrictions like
 *    the delay imposed by the timelock and the list of blocked functions. The
 *    admin can manage the list of blocked functions. In practice, the admin
 *    role is expected to (1) be inhabited by a contract requiring a secure
 *    quorum of votes before taking any action and (2) to be used rarely, namely
 *    only for emergency actions or configuration of the RBACTimelock.
 *
 * 2) Proposer: The proposer can schedule delayed operations that don't use any
 *    blocked function selector.
 *
 * 3) Executor: The executor can execute previously scheduled operations once
 *    their delay has expired. The contract enforces that the calls in an
 *    operation are executed with the correct args (target, data, value), but
 *    the executor can freely choose the gas limit. Since the executor is
 *    typically not particularly trusted, we recommend that (transitive) callees
 *    implement standard behavior of simply reverting if insufficient gas is
 *    provided. In particular, this means callees should not have non-reverting
 *    gas-dependent branches.
 *
 * 4) Canceller: The canceller can cancel operations that have been scheduled
 *    but not yet executed.
 *
 * 5) Bypasser: The bypasser can bypass any restrictions like the delay imposed
 *    by the timelock and the list of blocked functions to immediately execute
 *    operations, e.g. in case of emergencies.
 *
 * Note that this contract doesn't place any restrictions on the gas limit used
 * when executing operations. See the above comment on the executor role for
 * more details.
 *
 * @dev This contract is a modified version of OpenZeppelin's
 * contracts/governance/TimelockController.sol contract from v4.7.0, accessed in
 * commit 561d1061fc568f04c7a65853538e834a889751e8 of
 * github.com/OpenZeppelin/openzeppelin-contracts
 * Said contract is under "Copyright (c) 2016-2023 zOS Global Limited and
 * contributors" and its original MIT license can be found at
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/561d1061fc568f04c7a65853538e834a889751e8/LICENSE
 */
contract RBACTimelock is AccessControlEnumerable, IERC721Receiver, IERC1155Receiver {
  using EnumerableSet for EnumerableSet.Bytes32Set;

  struct Call {
    address target;
    uint256 value;
    bytes data;
  }

  bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
  bytes32 public constant PROPOSER_ROLE = keccak256("PROPOSER_ROLE");
  bytes32 public constant EXECUTOR_ROLE = keccak256("EXECUTOR_ROLE");
  bytes32 public constant CANCELLER_ROLE = keccak256("CANCELLER_ROLE");
  bytes32 public constant BYPASSER_ROLE = keccak256("BYPASSER_ROLE");
  uint256 internal constant _DONE_TIMESTAMP = uint256(1);

  mapping(bytes32 => uint256) private _timestamps;
  uint256 private _minDelay;
  EnumerableSet.Bytes32Set private _blockedFunctionSelectors;

  /**
   * @dev Emitted when a call is scheduled as part of operation `id`.
   */
  event CallScheduled(
    bytes32 indexed id,
    uint256 indexed index,
    address target,
    uint256 value,
    bytes data,
    bytes32 predecessor,
    bytes32 salt,
    uint256 delay
  );

  /**
   * @dev Emitted when a call is performed as part of operation `id`.
   */
  event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data);

  /**
   * @dev Emitted when a call is performed via bypasser.
   */
  event BypasserCallExecuted(uint256 indexed index, address target, uint256 value, bytes data);

  /**
   * @dev Emitted when operation `id` is cancelled.
   */
  event Cancelled(bytes32 indexed id);

  /**
   * @dev Emitted when the minimum delay for future operations is modified.
   */
  event MinDelayChange(uint256 oldDuration, uint256 newDuration);

  /**
   * @dev Emitted when a function selector is blocked.
   */
  event FunctionSelectorBlocked(bytes4 indexed selector);

  /**
   * @dev Emitted when a function selector is unblocked.
   */
  event FunctionSelectorUnblocked(bytes4 indexed selector);

  /**
   * @dev Initializes the contract with the following parameters:
   *
   * - `minDelay`: initial minimum delay for operations
   * - `admin`: account to be granted admin role
   * - `proposers`: accounts to be granted proposer role
   * - `executors`: accounts to be granted executor role
   * - `cancellers`: accounts to be granted canceller role
   * - `bypassers`: accounts to be granted bypasser role
   */
  constructor(
    uint256 minDelay,
    address admin,
    address[] memory proposers,
    address[] memory executors,
    address[] memory cancellers,
    address[] memory bypassers
  ) {
    _setRoleAdmin(ADMIN_ROLE, ADMIN_ROLE);
    _setRoleAdmin(PROPOSER_ROLE, ADMIN_ROLE);
    _setRoleAdmin(EXECUTOR_ROLE, ADMIN_ROLE);
    _setRoleAdmin(CANCELLER_ROLE, ADMIN_ROLE);
    _setRoleAdmin(BYPASSER_ROLE, ADMIN_ROLE);

    _setupRole(ADMIN_ROLE, admin);

    // register proposers
    for (uint256 i = 0; i < proposers.length; ++i) {
      _setupRole(PROPOSER_ROLE, proposers[i]);
    }

    // register executors
    for (uint256 i = 0; i < executors.length; ++i) {
      _setupRole(EXECUTOR_ROLE, executors[i]);
    }

    // register cancellers
    for (uint256 i = 0; i < cancellers.length; ++i) {
      _setupRole(CANCELLER_ROLE, cancellers[i]);
    }

    // register bypassers
    for (uint256 i = 0; i < bypassers.length; ++i) {
      _setupRole(BYPASSER_ROLE, bypassers[i]);
    }

    _minDelay = minDelay;
    emit MinDelayChange(0, minDelay);
  }

  /**
   * @dev Modifier to make a function callable only by a certain role or the
   * admin role.
   */
  modifier onlyRoleOrAdminRole(
    bytes32 role
  ) {
    address sender = _msgSender();
    if (!hasRole(ADMIN_ROLE, sender)) {
      _checkRole(role, sender);
    }
    _;
  }

  /**
   * @dev Contract might receive/hold ETH as part of the maintenance process.
   */
  receive() external payable {}

  /**
   * @dev See {IERC165-supportsInterface}.
   */
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual override(IERC165, AccessControlEnumerable) returns (bool) {
    return interfaceId == type(IERC1155Receiver).interfaceId || super.supportsInterface(interfaceId);
  }

  /**
   * @dev Returns whether an id correspond to a registered operation. This
   * includes both Pending, Ready and Done operations.
   */
  function isOperation(
    bytes32 id
  ) public view virtual returns (bool registered) {
    return getTimestamp(id) > 0;
  }

  /**
   * @dev Returns whether an operation is pending or not.
   */
  function isOperationPending(
    bytes32 id
  ) public view virtual returns (bool pending) {
    return getTimestamp(id) > _DONE_TIMESTAMP;
  }

  /**
   * @dev Returns whether an operation is ready or not.
   */
  function isOperationReady(
    bytes32 id
  ) public view virtual returns (bool ready) {
    uint256 timestamp = getTimestamp(id);
    return timestamp > _DONE_TIMESTAMP && timestamp <= block.timestamp;
  }

  /**
   * @dev Returns whether an operation is done or not.
   */
  function isOperationDone(
    bytes32 id
  ) public view virtual returns (bool done) {
    return getTimestamp(id) == _DONE_TIMESTAMP;
  }

  /**
   * @dev Returns the timestamp at with an operation becomes ready (0 for
   * unset operations, 1 for done operations).
   */
  function getTimestamp(
    bytes32 id
  ) public view virtual returns (uint256 timestamp) {
    return _timestamps[id];
  }

  /**
   * @dev Returns the minimum delay for an operation to become valid.
   *
   * This value can be changed by executing an operation that calls `updateDelay`.
   */
  function getMinDelay() public view virtual returns (uint256 duration) {
    return _minDelay;
  }

  /**
   * @dev Returns the identifier of an operation containing a batch of
   * transactions.
   */
  function hashOperationBatch(
    Call[] calldata calls,
    bytes32 predecessor,
    bytes32 salt
  ) public pure virtual returns (bytes32 hash) {
    return keccak256(abi.encode(calls, predecessor, salt));
  }

  /**
   * @dev Schedule an operation containing a batch of transactions.
   *
   * Emits one {CallScheduled} event per transaction in the batch.
   *
   * Requirements:
   *
   * - the caller must have the 'proposer' or 'admin' role.
   * - all payloads must not start with a blocked function selector.
   */
  function scheduleBatch(
    Call[] calldata calls,
    bytes32 predecessor,
    bytes32 salt,
    uint256 delay
  ) public virtual onlyRoleOrAdminRole(PROPOSER_ROLE) {
    bytes32 id = hashOperationBatch(calls, predecessor, salt);
    _schedule(id, delay);
    for (uint256 i = 0; i < calls.length; ++i) {
      _checkFunctionSelectorNotBlocked(calls[i].data);
      emit CallScheduled(id, i, calls[i].target, calls[i].value, calls[i].data, predecessor, salt, delay);
    }
  }

  /**
   * @dev Schedule an operation that becomes valid after a given delay.
   */
  function _schedule(bytes32 id, uint256 delay) private {
    require(!isOperation(id), "RBACTimelock: operation already scheduled");
    require(delay >= getMinDelay(), "RBACTimelock: insufficient delay");
    _timestamps[id] = block.timestamp + delay;
  }

  /**
   * @dev Cancel an operation.
   *
   * Requirements:
   *
   * - the caller must have the 'canceller' or 'admin' role.
   */
  function cancel(
    bytes32 id
  ) public virtual onlyRoleOrAdminRole(CANCELLER_ROLE) {
    require(isOperationPending(id), "RBACTimelock: operation cannot be cancelled");
    delete _timestamps[id];

    emit Cancelled(id);
  }

  /**
   * @dev Execute an (ready) operation containing a batch of transactions.
   * Note that we perform a raw call to each target. Raw calls to targets that
   * don't have associated contract code will always succeed regardless of
   * payload.
   *
   * Emits one {CallExecuted} event per transaction in the batch.
   *
   * Requirements:
   *
   * - the caller must have the 'executor' or 'admin' role.
   */
  function executeBatch(
    Call[] calldata calls,
    bytes32 predecessor,
    bytes32 salt
  ) public payable virtual onlyRoleOrAdminRole(EXECUTOR_ROLE) {
    bytes32 id = hashOperationBatch(calls, predecessor, salt);

    _beforeCall(id, predecessor);
    for (uint256 i = 0; i < calls.length; ++i) {
      _execute(calls[i]);
      emit CallExecuted(id, i, calls[i].target, calls[i].value, calls[i].data);
    }
    _afterCall(id);
  }

  /**
   * @dev Execute an operation's call.
   */
  function _execute(
    Call calldata call
  ) internal virtual {
    (bool success,) = call.target.call{value: call.value}(call.data);
    require(success, "RBACTimelock: underlying transaction reverted");
  }

  /**
   * @dev Checks before execution of an operation's calls.
   */
  function _beforeCall(bytes32 id, bytes32 predecessor) private view {
    require(isOperationReady(id), "RBACTimelock: operation is not ready");
    require(predecessor == bytes32(0) || isOperationDone(predecessor), "RBACTimelock: missing dependency");
  }

  /**
   * @dev Checks after execution of an operation's calls.
   */
  function _afterCall(
    bytes32 id
  ) private {
    require(isOperationReady(id), "RBACTimelock: operation is not ready");
    _timestamps[id] = _DONE_TIMESTAMP;
  }

  /**
   * @dev Changes the minimum timelock duration for future operations.
   *
   * Emits a {MinDelayChange} event.
   *
   * Requirements:
   *
   * - the caller must have the 'admin' role.
   */
  function updateDelay(
    uint256 newDelay
  ) external virtual onlyRole(ADMIN_ROLE) {
    emit MinDelayChange(_minDelay, newDelay);
    _minDelay = newDelay;
  }

  /**
   * @dev See {IERC721Receiver-onERC721Received}.
   */
  function onERC721Received(address, address, uint256, bytes memory) public virtual override returns (bytes4) {
    return this.onERC721Received.selector;
  }

  /**
   * @dev See {IERC1155Receiver-onERC1155Received}.
   */
  function onERC1155Received(address, address, uint256, uint256, bytes memory) public virtual override returns (bytes4) {
    return this.onERC1155Received.selector;
  }

  /**
   * @dev See {IERC1155Receiver-onERC1155BatchReceived}.
   */
  function onERC1155BatchReceived(
    address,
    address,
    uint256[] memory,
    uint256[] memory,
    bytes memory
  ) public virtual override returns (bytes4) {
    return this.onERC1155BatchReceived.selector;
  }

  /*
   * New functions not present in original OpenZeppelin TimelockController
   */

  /**
   * @dev Blocks a function selector from being used, i.e. schedule
   * operations with this function selector will revert.
   * Note that blocked selectors are only checked when an operation is being
   * scheduled, not when it is executed. You may want to check any pending
   * operations for whether they contain the blocked selector and cancel them.
   *
   * Requirements:
   *
   * - the caller must have the 'admin' role.
   */
  function blockFunctionSelector(
    bytes4 selector
  ) external onlyRole(ADMIN_ROLE) {
    if (_blockedFunctionSelectors.add(selector)) {
      emit FunctionSelectorBlocked(selector);
    }
  }

  /**
   * @dev Unblocks a previously blocked function selector so it can be used again.
   * Requirements:
   *
   * - the caller must have the 'admin' role.
   */
  function unblockFunctionSelector(
    bytes4 selector
  ) external onlyRole(ADMIN_ROLE) {
    if (_blockedFunctionSelectors.remove(selector)) {
      emit FunctionSelectorUnblocked(selector);
    }
  }

  /**
   * @dev Returns the number of blocked function selectors.
   */
  function getBlockedFunctionSelectorCount() external view returns (uint256) {
    return _blockedFunctionSelectors.length();
  }

  /**
   * @dev Returns the blocked function selector with the given index. Function
   * selectors are not sorted in any particular way, and their ordering may
   * change at any point.
   *
   * WARNING: When using {getBlockedFunctionSelectorCount} and
   * {getBlockedFunctionSelectorAt} via RPC, make sure you perform all queries
   * on the same block. When using these functions within an onchain
   * transaction, make sure that the state of this contract hasn't changed in
   * between invocations to avoid time-of-check time-of-use bugs.
   * See the following
   * https://forum.openzeppelin.com/t/iterating-over-elements-on-enumerableset-in-openzeppelin-contracts/2296[forum
   * post] for more information.
   */
  function getBlockedFunctionSelectorAt(
    uint256 index
  ) external view returns (bytes4) {
    return bytes4(_blockedFunctionSelectors.at(index));
  }

  /**
   * @dev Directly execute a batch of transactions, bypassing any other
   * checks.
   * Note that we perform a raw call to each target. Raw calls to targets that
   * don't have associated contract code will always succeed regardless of
   * payload.
   *
   * Emits one {BypasserCallExecuted} event per transaction in the batch.
   *
   * Requirements:
   *
   * - the caller must have the 'bypasser' or 'admin' role.
   */
  function bypasserExecuteBatch(
    Call[] calldata calls
  ) public payable virtual onlyRoleOrAdminRole(BYPASSER_ROLE) {
    for (uint256 i = 0; i < calls.length; ++i) {
      _execute(calls[i]);
      emit BypasserCallExecuted(i, calls[i].target, calls[i].value, calls[i].data);
    }
  }

  /**
   * @dev Checks to see if the function being scheduled is blocked.  This
   * is used when trying to schedule or batch schedule an operation.
   */
  function _checkFunctionSelectorNotBlocked(
    bytes calldata data
  ) private view {
    if (data.length < 4) {
      return;
    }
    bytes4 selector = bytes4(data[:4]);
    require(!_blockedFunctionSelectors.contains(bytes32(selector)), "RBACTimelock: selector is blocked");
  }
}
