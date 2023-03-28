// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {SafeMath} from "@openzeppelin/contracts/utils/math/SafeMath.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {CCIPReceiver} from "../ccip/applications/CCIPReceiver.sol";
import {IRouterClient} from "../ccip/interfaces/IRouterClient.sol";
import {Client} from "../ccip/models/Client.sol";
import {ERC2771Recipient} from "../metatx/ERC2771Recipient.sol";
import {IForwarder} from "../metatx/IForwarder.sol";
import {OwnerIsCreator} from "../ccip/OwnerIsCreator.sol";

contract MetaERC20 is CCIPReceiver, IERC20, ERC2771Recipient, OwnerIsCreator {
  using SafeMath for uint256;

  string public constant name = "BankToken";
  string public constant symbol = "BANKTOKEN";
  uint8 public constant decimals = 18;

  // @dev implements IERC20.totalSupply
  uint256 public override totalSupply;
  // @dev implements IERC20.balanceOf
  mapping(address => uint256) public override balanceOf;
  // @dev implements IERC20.allowance
  mapping(address => mapping(address => uint256)) public override allowance;
  // @dev fee token used for cross chain transfer using CCIP
  IERC20 private s_ccipFeeToken;
  // @dev forwarder verifies signatures for meta transactions and forwards the 
  // request to this contract
  IForwarder private s_fowarder;
  
  constructor(uint256 _totalSupply, address router) CCIPReceiver(router) {
    totalSupply = _totalSupply;
    balanceOf[msg.sender] = totalSupply;
  }

  function setFeeToken(IERC20 feeToken) external onlyOwner {
    // TODO: validate fee token
    s_ccipFeeToken = feeToken;
  }

  function setForwarder(IForwarder forwarder) external onlyOwner {
    _setTrustedForwarder(address(forwarder));
  }

  /**
   * @dev Moves `amount` tokens from the caller's account to `to`.
   * Emits a {Transfer} event.
   * @param to token receiver
   * @param amount token amount to transfer in wei or equivalent
   * @return bool true if transfer is successful, false otherwise
   */
  function transfer(address to, uint256 amount) external override returns (bool) {
    _transfer(_msgSender(), to, amount);
    return true;
  }

  function _transfer(address from, address to, uint256 amount) private {
    balanceOf[from] = balanceOf[from].sub(amount);
    balanceOf[to] = balanceOf[to].add(amount);
    emit Transfer(from, to, amount);
  }

  function _approve(address owner, address spender, uint amount) private {
    allowance[owner][spender] = amount;
    emit Approval(owner, spender, amount);
  }

  /**
   * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
   *
   * Returns a boolean value indicating whether the operation succeeded.
   *
   * IMPORTANT: Beware that changing an allowance with this method brings the risk
   * that someone may use both the old and the new allowance by unfortunate
   * transaction ordering. One possible solution to mitigate this race
   * condition is to first reduce the spender's allowance to 0 and set the
   * desired value afterwards:
   * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
   *
   * Emits an {Approval} event.
   * 
   * @param spender address to approve for spending
   * @param amount token amount to transfer in wei or equivalent
   * @return bool true if transfer is successful, false otherwise
   */
  function approve(address spender, uint256 amount) external override returns (bool) {
    _approve(_msgSender(), spender, amount);
    return true;
  }

  /**
   * @dev Moves `amount` tokens from `from` to `to` using the
   * allowance mechanism. `amount` is then deducted from the caller's
   * allowance.
   *
   * Returns a boolean value indicating whether the operation succeeded.
   *
   * Emits a {Transfer} event.
   * 
   * @param from address to transfer token from
   * @param to address to transfer token to
   * @param amount token amount to transfer in wei or equivalent
   * @return bool true if transfer is successful, false otherwise
   */
  function transferFrom(
    address from,
    address to,
    uint256 amount
  ) external override returns (bool) {
    if (allowance[from][_msgSender()] != type(uint256).max) {
      allowance[from][_msgSender()] = allowance[from][_msgSender()].sub(amount);
    }
    _transfer(from, to, amount);
    return true;
  }

  /**
   * @dev Moves amount of this token to recipient address in destination chain.
   * encapsulate recipient address and amount in the CCIP message.
   * the amount is burnt from source chain and minted in the destination chain
   * 
   * @param destinationTokenAddress ERC20 token address in destination chain
   * @param recipientAddress token recipient address in destination chain
   * @param amount token amount to transfer in wei or equivalent
   * @param destinationChainId destination chain ID
   */
  function metaTransfer(address destinationTokenAddress, address recipientAddress, uint256 amount, uint64 destinationChainId) external validateTrustedForwarder {
    Client.EVMTokenAmount[] memory tokenAmounts;

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(destinationTokenAddress),
      data: abi.encode(recipientAddress, amount),
      tokenAmounts: tokenAmounts, //empty tokenAmounts 
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 500_000, strict: false})), // hard-coded gas limit for now. can be set to a mutable state variable
      feeToken: address(s_ccipFeeToken)
    });
    _burn(_msgSender(), amount);
    IRouterClient(getRouter()).ccipSend(destinationChainId, message);
  }
  
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    (address recipient, uint256 amount) = abi.decode(message.data, (address, uint256));
    _mint(recipient, amount);
  }

  /**
   * @notice Fund this contract with configured feeToken and approve tokens to the router
   * @dev Requires prior approval from the msg.sender
   * @param amount The amount of feeToken to be funded
   */
  function fund(uint256 amount) external {
    s_ccipFeeToken.transferFrom(_msgSender(), address(this), amount);
    s_ccipFeeToken.approve(address(getRouter()), amount);
  }

  function _mint(address account, uint256 amount) internal {
    require(account != address(0), "ERC20: mint to the zero address");
    totalSupply += amount;
    balanceOf[account] += amount;
    emit Transfer(address(0), account, amount);
  }

  function _burn(address account, uint256 amount) internal {
    require(account != address(0), "ERC20: burn from the zero address");
    uint256 accountBalance = balanceOf[account];
    require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
    unchecked {
        balanceOf[account] = accountBalance - amount;
    }
    totalSupply -= amount;
    emit Transfer(account, address(0), amount);
  }

  error MustBeTrustedForwarder(address sender);

  modifier validateTrustedForwarder() {
    if (!isTrustedForwarder(msg.sender)) {
      revert MustBeTrustedForwarder(msg.sender);
    }
    _;
  }
}