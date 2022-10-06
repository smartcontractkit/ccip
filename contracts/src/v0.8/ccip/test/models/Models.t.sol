// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenSetup} from "../TokenSetup.t.sol";
import {CCIP} from "../../models/Models.sol";

contract Models__toAny2EVMMessage is TokenSetup {
  using CCIP for CCIP.Any2EVMMessageFromSender;

  function testSuccess() public {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 42;
    amounts[1] = 1337;
    CCIP.Any2EVMMessageFromSender memory message = CCIP.Any2EVMMessageFromSender({
      sourceChainId: SOURCE_CHAIN_ID,
      sender: abi.encode(OWNER),
      receiver: STRANGER,
      data: abi.encode(STRANGER),
      destTokens: s_destTokens,
      destPools: s_destPools,
      amounts: amounts,
      gasLimit: 1234567890
    });

    CCIP.Any2EVMMessage memory messageForReceiver = message._toAny2EVMMessage();

    assertEq(message.sourceChainId, messageForReceiver.sourceChainId);
    assertEq(message.sender, messageForReceiver.sender);
    assertEq(message.data, messageForReceiver.data);
    assertEq(address(message.destTokens[0]), address(messageForReceiver.destTokens[0]));
    assertEq(address(message.destTokens[1]), address(messageForReceiver.destTokens[1]));
    assertEq(message.amounts, messageForReceiver.amounts);
  }
}
