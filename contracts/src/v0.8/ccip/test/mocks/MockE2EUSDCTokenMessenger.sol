// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITokenMessenger} from "../../pools/USDC/ITokenMessenger.sol";
import {IMessageTransmitterWithRelay} from "./interfaces/IMessageTransmitterWithRelay.sol";
import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

// This contract mocks both the ITokenMessenger and IMessageTransmitter
// contracts involved with the Cross Chain Token Protocol.
contract MockE2EUSDCTokenMessenger is ITokenMessenger {
  uint32 private immutable i_messageBodyVersion;
  bytes32 public constant i_destinationTokenMessenger = keccak256("i_destinationTokenMessenger");
  uint64 public s_nonce;
  address private i_transmitter;

  // Local Message Transmitter responsible for sending and receiving messages to/from remote domains
  IMessageTransmitterWithRelay public immutable localMessageTransmitterWithRelay;

  constructor(uint32 version, address transmitter) {
    i_messageBodyVersion = version;
    s_nonce = 1;
    i_transmitter = transmitter;
    localMessageTransmitterWithRelay = IMessageTransmitterWithRelay(transmitter);
  }

  // The mock function is based on the same function in https://github.com/circlefin/evm-cctp-contracts/blob/master/src/TokenMessenger.sol
  function depositForBurnWithCaller(
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller
  ) external returns (uint64) {
    IBurnMintERC20(burnToken).transferFrom(msg.sender, address(this), amount);
    IBurnMintERC20(burnToken).burn(amount);
    // Format message body
    bytes memory _burnMessage = abi.encodePacked(
      i_messageBodyVersion,
      burnToken,
      mintRecipient,
      amount,
      bytes32(uint256(uint160((msg.sender))))
    );
    s_nonce = _sendDepositForBurnMessage(
      destinationDomain,
      i_destinationTokenMessenger,
      destinationCaller,
      _burnMessage
    );
    emit DepositForBurn(
      s_nonce,
      burnToken,
      amount,
      msg.sender,
      mintRecipient,
      destinationDomain,
      i_destinationTokenMessenger,
      destinationCaller
    );
    return s_nonce;
  }

  function messageBodyVersion() external view returns (uint32) {
    return i_messageBodyVersion;
  }

  function localMessageTransmitter() external view returns (address) {
    return i_transmitter;
  }

  /**
   * @notice Sends a BurnMessage through the local message transmitter
   * @dev calls local message transmitter's sendMessage() function if `_destinationCaller` == bytes32(0),
   * or else calls sendMessageWithCaller().
   * @param _destinationDomain destination domain
   * @param _destinationTokenMessenger address of registered TokenMessenger contract on destination domain, as bytes32
   * @param _destinationCaller caller on the destination domain, as bytes32. If `_destinationCaller` == bytes32(0),
   * any address can call receiveMessage() on destination domain.
   * @param _burnMessage formatted BurnMessage bytes (message body)
   * @return nonce unique nonce reserved by message
   */
  function _sendDepositForBurnMessage(
    uint32 _destinationDomain,
    bytes32 _destinationTokenMessenger,
    bytes32 _destinationCaller,
    bytes memory _burnMessage
  ) internal returns (uint64 nonce) {
    if (_destinationCaller == bytes32(0)) {
      return localMessageTransmitterWithRelay.sendMessage(_destinationDomain, _destinationTokenMessenger, _burnMessage);
    } else {
      return
        localMessageTransmitterWithRelay.sendMessageWithCaller(
          _destinationDomain,
          _destinationTokenMessenger,
          _destinationCaller,
          _burnMessage
        );
    }
  }
}
