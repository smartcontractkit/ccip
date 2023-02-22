// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IPriceRegistry} from "../../interfaces/prices/IPriceRegistry.sol";
import {IEVM2EVMOnRamp} from "../../interfaces/onRamp/IEVM2EVMOnRamp.sol";
import {IRouter} from "../../interfaces/router/IRouter.sol";

import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";
import {PriceRegistry} from "../../prices/PriceRegistry.sol";
import {RouterSetup} from "../router/RouterSetup.t.sol";
import {PriceRegistrySetup} from "../prices/PriceRegistry.t.sol";
import {Internal} from "../../models/Internal.sol";
import {Client} from "../../models/Client.sol";
import "../../offRamp/EVM2EVMOffRamp.sol";
import "../TokenSetup.t.sol";

contract EVM2EVMOnRampSetup is TokenSetup, PriceRegistrySetup {
  // Duplicate event of the CCIPSendRequested in the IOnRamp
  event CCIPSendRequested(Internal.EVM2EVMMessage message);

  uint256 internal immutable i_tokenAmount0 = 9;
  uint256 internal immutable i_tokenAmount1 = 7;

  bytes32 internal s_metadataHash;

  address[] internal s_allowList;

  EVM2EVMOnRamp internal s_onRamp;
  address[] s_offRamps;

  function setUp() public virtual override(TokenSetup, PriceRegistrySetup) {
    TokenSetup.setUp();
    PriceRegistrySetup.setUp();

    IEVM2EVMOnRamp.FeeTokenConfigArgs[] memory feeTokenConfigArgs = new IEVM2EVMOnRamp.FeeTokenConfigArgs[](2);
    feeTokenConfigArgs[0] = IEVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceTokens[0],
      feeAmount: 1,
      multiplier: 108e16,
      destGasOverhead: 1
    });
    feeTokenConfigArgs[1] = IEVM2EVMOnRamp.FeeTokenConfigArgs({
      token: s_sourceRouter.getWrappedNative(),
      feeAmount: 2,
      multiplier: 108e16,
      destGasOverhead: 2
    });
    s_onRamp = new EVM2EVMOnRamp(
      IEVM2EVMOnRamp.Chains({chainId: SOURCE_CHAIN_ID, destChainId: DEST_CHAIN_ID}),
      getTokensAndPools(s_sourceTokens, getCastedSourcePools()),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      address(s_sourceRouter),
      address(s_priceRegistry),
      feeTokenConfigArgs,
      s_sourceTokens[0],
      getNopsAndWeights()
    );

    s_metadataHash = keccak256(
      abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, address(s_onRamp))
    );

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    LockReleaseTokenPool(address(s_sourcePools[0])).setOnRamp(address(s_onRamp), true);
    LockReleaseTokenPool(address(s_sourcePools[1])).setOnRamp(address(s_onRamp), true);

    s_offRamps = new address[](2);
    s_offRamps[0] = address(10);
    s_offRamps[1] = address(11);
    IRouter.OnRampUpdate[] memory onRampUpdates = new IRouter.OnRampUpdate[](1);
    IRouter.OffRampUpdate[] memory offRampUpdates = new IRouter.OffRampUpdate[](1);
    onRampUpdates[0] = IRouter.OnRampUpdate({destChainId: DEST_CHAIN_ID, onRamp: address(s_onRamp)});
    offRampUpdates[0] = IRouter.OffRampUpdate({sourceChainId: SOURCE_CHAIN_ID, offRamps: s_offRamps});
    s_sourceRouter.applyRampUpdates(onRampUpdates, offRampUpdates);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    IERC20(s_sourceTokens[0]).approve(address(s_sourceRouter), 2**128);
    IERC20(s_sourceTokens[1]).approve(address(s_sourceRouter), 2**128);
  }

  function assertSameConfig(IEVM2EVMOnRamp.OnRampConfig memory a, IEVM2EVMOnRamp.OnRampConfig memory b) public {
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.maxGasLimit, b.maxGasLimit);
  }

  function _generateTokenMessage() public view returns (Client.EVM2AnyMessage memory) {
    Client.EVMTokenAmount[] memory tokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
    tokenAmounts[0].amount = i_tokenAmount0;
    tokenAmounts[1].amount = i_tokenAmount1;
    return
      Client.EVM2AnyMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokenAmounts: tokenAmounts,
        feeToken: s_sourceFeeToken,
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _generateEmptyMessage() public view returns (Client.EVM2AnyMessage memory) {
    return
      Client.EVM2AnyMessage({
        receiver: abi.encode(OWNER),
        data: "",
        tokenAmounts: new Client.EVMTokenAmount[](0),
        feeToken: s_sourceFeeToken,
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: GAS_LIMIT, strict: false}))
      });
  }

  function _messageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 seqNum,
    uint64 nonce,
    uint256 feeTokenAmount
  ) public view returns (Internal.EVM2EVMMessage memory) {
    Internal.EVM2EVMMessage memory messageEvent = Internal.EVM2EVMMessage({
      sequenceNumber: seqNum,
      feeTokenAmount: feeTokenAmount,
      sender: OWNER,
      nonce: nonce,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainId: SOURCE_CHAIN_ID,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokenAmounts: message.tokenAmounts,
      feeToken: message.feeToken,
      messageId: ""
    });

    messageEvent.messageId = Internal._hash(messageEvent, s_metadataHash);
    return messageEvent;
  }
}
