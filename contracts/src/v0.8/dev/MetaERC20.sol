// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import {SafeMath} from "@openzeppelin/contracts/utils/math/SafeMath.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {CCIPReceiver} from "../ccip/applications/CCIPReceiver.sol";
import {IRouterClient} from "../ccip/interfaces/router/IRouterClient.sol";
import {Client} from "../ccip/models/Client.sol";
import {IBurnMintERC20} from "../ccip/interfaces/pools/IBurnMintERC20.sol";

contract MetaERC20 is IERC20, CCIPReceiver, IBurnMintERC20 {
  using SafeMath for uint256;

  string public constant name = "BankToken";
  string public constant symbol = "BANKTOKEN";
  uint8 public constant decimals = 18;

  // NOTE: implements IERC20.totalSupply
  uint256 public override totalSupply;
  // NOTE: implements IERC20.balanceOf
  mapping(address => uint256) public override balanceOf;
  // NOTE: implements IERC20.allowance
  mapping(address => mapping(address => uint256)) public override allowance;

  IERC20 private s_ccipFeeToken;

  bytes32 public DOMAIN_SEPARATOR;
  // keccak256("MetaTransfer(address,address,uint256,uint256,uint64,uint64,uint256)");
  // arguments in order: owner,to,amount,nonce,chainId,gasLimit,deadline
  bytes32 public constant META_TRANSFER_TYPEHASH = 0xfa015aba6914852681b8c40c25589bc710e6970e3bc80c5140186542aec3299a;
  mapping(address => uint256) public nonces;

  struct MetaTransferMessage {
    address owner;
    address to; 
    uint256 amount; 
    uint256 deadline; 
    uint64 chainId; 
    uint64 gasLimit;
  }

  constructor(uint256 _totalSupply, IERC20 ccipFeeToken, address router) CCIPReceiver(router) public {
    totalSupply = _totalSupply;
    balanceOf[msg.sender] = totalSupply;
    s_ccipFeeToken = ccipFeeToken;

    DOMAIN_SEPARATOR = keccak256(
      abi.encode(
        keccak256('EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)'),
        keccak256(bytes(name)),
        keccak256(bytes('1')),
        block.chainid,
        address(this)
      )
    );
  }

  /**
   * @dev Moves `amount` tokens from the caller's account to `to`.
   *
   * Returns a boolean value indicating whether the operation succeeded.
   *
   * Emits a {Transfer} event.
   */
  function transfer(address to, uint256 amount) external override returns (bool) {
    _transfer(msg.sender, to, amount);
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
   */
  function approve(address spender, uint256 amount) external override returns (bool) {
    _approve(msg.sender, spender, amount);
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
   */
  function transferFrom(
    address from,
    address to,
    uint256 amount
  ) external override returns (bool) {
    if (allowance[from][msg.sender] != type(uint256).max) {
      allowance[from][msg.sender] = allowance[from][msg.sender].sub(amount);
    }
    _transfer(from, to, amount);
    return true;
  }

  function metaTransfer(MetaTransferMessage calldata m, uint8 v, bytes32 r, bytes32 s) external {
    require(m.deadline >= block.timestamp, 'EXPIRED');
    bytes32 digest = keccak256(
      abi.encodePacked(
        '\x19\x01',
        DOMAIN_SEPARATOR,
        keccak256(abi.encode(META_TRANSFER_TYPEHASH, 
          m.owner, 
          m.to, 
          m.amount, 
          nonces[m.owner]++, 
          m.chainId, 
          m.gasLimit, 
          m.deadline
        ))
      )
    );
    address recoveredAddress = ecrecover(digest, v, r, s);
    require(recoveredAddress != address(0) && recoveredAddress == m.owner, 'INVALID_SIGNATURE');
    if (m.chainId == block.chainid) {
      _transfer(m.owner, m.to, m.amount);
    } else {
      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(m.to),
        data: "", //empty   
        tokenAmounts: new Client.EVMTokenAmount[](m.amount),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: m.gasLimit, strict: false})), //not sure what these parameters do
        feeToken: address(s_ccipFeeToken)
      });
      IRouterClient(getRouter()).ccipSend(m.chainId, message);
    }
  }

  /**
   * @notice Fund this contract with configured feeToken and approve tokens to the router
   * @dev Requires prior approval from the msg.sender
   * @param amount The amount of feeToken to be funded
   */
  function fund(uint256 amount) external {
    s_ccipFeeToken.transferFrom(msg.sender, address(this), amount);
    s_ccipFeeToken.approve(address(getRouter()), amount);
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    // do nothing for now
  }

  //function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
  //  uint256 pingPongCount = abi.decode(message.data, (uint256));
  //  if (!s_isPaused) {
  //    _respond(pingPongCount + 1);
  //  }
  //}

  //TODO: Add access control for minting
  function mint(address account, uint256 amount) external override {
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
}