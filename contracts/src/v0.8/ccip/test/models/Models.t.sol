// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenSetup} from "../TokenSetup.t.sol";
import {Internal} from "../../models/Internal.sol";
import {Common} from "../../models/Common.sol";

contract Models__toAny2EVMMessage is TokenSetup {
  using Internal for Internal.Any2EVMMessageFromSender;

  function testSuccess() public {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedDestinationEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = 42;
    tokensAndAmounts[1].amount = 1337;
    Internal.Any2EVMMessageFromSender memory message = Internal.Any2EVMMessageFromSender({
      sourceChainId: SOURCE_CHAIN_ID,
      sender: abi.encode(OWNER),
      receiver: STRANGER,
      data: abi.encode(STRANGER),
      destTokensAndAmounts: tokensAndAmounts,
      destPools: s_destPools,
      gasLimit: 1234567890
    });

    Common.Any2EVMMessage memory messageForReceiver = message._toAny2EVMMessage();

    assertEq(message.sourceChainId, messageForReceiver.sourceChainId);
    assertEq(message.sender, messageForReceiver.sender);
    assertEq(message.data, messageForReceiver.data);
    assertEq(message.destTokensAndAmounts[0].token, messageForReceiver.destTokensAndAmounts[0].token);
    assertEq(message.destTokensAndAmounts[0].amount, messageForReceiver.destTokensAndAmounts[0].amount);
    assertEq(message.destTokensAndAmounts[1].token, messageForReceiver.destTokensAndAmounts[1].token);
    assertEq(message.destTokensAndAmounts[1].amount, messageForReceiver.destTokensAndAmounts[1].amount);
  }
}
