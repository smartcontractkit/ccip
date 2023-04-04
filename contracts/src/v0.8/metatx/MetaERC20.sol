// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {IRouterClient} from "../ccip/interfaces/IRouterClient.sol";
import {Client} from "../ccip/models/Client.sol";
import {ERC2771Recipient} from "../vendor/ERC2771Recipient.sol";
import {IForwarder} from "../metatx/IForwarder.sol";
import {OwnerIsCreator} from "../ccip/OwnerIsCreator.sol";
import {IBurnMintERC20} from "../ccip/interfaces/pools/IBurnMintERC20.sol";

// TODO: implement lock-and-mint token pool instead of burn-and-mint
contract MetaERC20 is ERC2771Recipient, OwnerIsCreator, IBurnMintERC20 {
  string public constant name = "BankToken";
  string public constant symbol = "BANKTOKEN";
  uint8 public constant decimals = 18;

  /// @dev implements IERC20.totalSupply
  uint256 public override totalSupply;
  /// @dev implements IERC20.balanceOf
  mapping(address => uint256) public override balanceOf;
  /// @dev implements IERC20.allowance
  mapping(address => mapping(address => uint256)) public override allowance;
  /// @dev forwarder verifies signatures for meta transactions and forwards the 
  /// @dev request to this contract
  IForwarder private s_fowarder;
  IRouterClient private s_ccip_router;
  /// @dev boolean only used for testing. Should be set to false in production
  /// @dev go-ethereum.simulatedBackend (used for testing) doesn't allow custom chain IDs
  /// @dev so block.chainid is hard-coded to 1337. 
  bool private s_test_only_force_cross_chain_transfer;
  
  constructor(uint256 _totalSupply, address _ccip_router, bool _test_only_force_cross_chain_transfer) {
    totalSupply = _totalSupply;
    balanceOf[msg.sender] = totalSupply;
    s_ccip_router = IRouterClient(_ccip_router);
    s_test_only_force_cross_chain_transfer = _test_only_force_cross_chain_transfer;
  }

  function setForwarder(IForwarder forwarder) external onlyOwner {
    _setTrustedForwarder(address(forwarder));
  }

   /// @dev Moves `amount` tokens from the caller's account to `to`.
   /// Emits a {Transfer} event.
   /// @param to token receiver
   /// @param amount token amount to transfer in wei or equivalent
   /// @return bool true if transfer is successful, false otherwise
  function transfer(address to, uint256 amount) external override returns (bool) {
    _transfer(_msgSender(), to, amount);
    return true;
  }

  function _transfer(address from, address to, uint256 amount) private {
    balanceOf[from] = balanceOf[from] - amount;
    balanceOf[to] = balanceOf[to] + amount;
    emit Transfer(from, to, amount);
  }

  function _approve(address owner, address spender, uint amount) private {
    allowance[owner][spender] = amount;
    emit Approval(owner, spender, amount);
  }

  /// @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
  /// @dev Returns a boolean value indicating whether the operation succeeded.
  /// @dev IMPORTANT: Beware that changing an allowance with this method brings the risk
  /// @dev that someone may use both the old and the new allowance by unfortunate
  /// @dev transaction ordering. One possible solution to mitigate this race
  /// @dev condition is to first reduce the spender's allowance to 0 and set the
  /// @dev desired value afterwards:
  /// @dev https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
  /// @dev Emits an {Approval} event.
  /// @param spender address to approve for spending
  /// @param amount token amount to transfer in wei or equivalent
  /// @return bool true if transfer is successful, false otherwise
  function approve(address spender, uint256 amount) external override returns (bool) {
    _approve(_msgSender(), spender, amount);
    return true;
  }

  /// @dev Moves `amount` tokens from `from` to `to` using the
  /// @dev allowance mechanism. `amount` is then deducted from the caller's
  /// @dev allowance.
  /// @dev Returns a boolean value indicating whether the operation succeeded.
  /// @dev Emits a {Transfer} event.
  /// @param from address to transfer token from
  /// @param to address to transfer token to
  /// @param amount token amount to transfer in wei or equivalent
  /// @return bool true if transfer is successful, false otherwise
  function transferFrom(
    address from,
    address to,
    uint256 amount
  ) external override returns (bool) {
    if (allowance[from][_msgSender()] != type(uint256).max) {
      allowance[from][_msgSender()] = allowance[from][_msgSender()] - amount;
    }
    _transfer(from, to, amount);
    return true;
  }

  error InsufficientBalance(address from, uint256 currentBalance, uint256 requiredBalance);

  /// @dev Transfers "amount" of this token to receiver address in destination chain.
  /// @param receiver token receiver address in destination chain. Handles distribution of tokens to recipients
  /// @param amount total token amount to be transferred
  /// @param destinationChainId destination chain ID
  function metaTransfer(
    address receiver,
    uint256 amount, 
    uint64 destinationChainId
  ) external validateTrustedForwarder {
    if (!isCrossChainTransfer(destinationChainId)) {
      _transfer(_msgSender(), receiver, amount);
      return;
    }
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({
      token: address(this),
      amount: amount
    });
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: address(0), // use native token instead of ERC20 tokens
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000, strict: false}))
    });
    uint256 fee = s_ccip_router.getFee(destinationChainId, message);

    s_ccip_router.ccipSend{value: fee}(destinationChainId, message);
  }

  /// @notice Fund and approve MetaERC20 tokens
  /// @dev Requires prior approval from the msg.sender
  /// @param spender Approval is given to the spender
  /// @param amount The amount of feeToken to be funded
  function fund(address spender, uint256 amount) external {
    _transfer(_msgSender(), address(this), amount);
    _approve(address(this), spender, amount);
  }
  
  receive() external payable {}

  error WithdrawFailure();

  function withdrawNative() external onlyOwner {
    uint256 amount = address(this).balance;
    // Owner can receive Ether since the address of owner is payable
    (bool success, ) = owner().call{value: amount}("");
    if (!success) {
      revert WithdrawFailure();
    }
  }

  function mint(address account, uint256 amount) external {
    require(account != address(0), "ERC20: mint to the zero address");
    totalSupply += amount;
    balanceOf[account] += amount;
    emit Transfer(address(0), account, amount);
  }

  function burn(address account, uint256 amount) external {
    require(account != address(0), "ERC20: burn from the zero address");
    uint256 accountBalance = balanceOf[account];
    require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
    unchecked {
        balanceOf[account] = accountBalance - amount;
    }
    totalSupply -= amount;
    emit Transfer(account, address(0), amount);
  }

  function isCrossChainTransfer(uint64 chainId) private view returns (bool) {
    if (s_test_only_force_cross_chain_transfer) {
      return true;
    }
    return chainId != block.chainid;
  }

  error MustBeTrustedForwarder(address sender);

  modifier validateTrustedForwarder() {
    if (!isTrustedForwarder(msg.sender)) {
      revert MustBeTrustedForwarder(msg.sender);
    }
    _;
  }
}