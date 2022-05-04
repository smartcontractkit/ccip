export const GAS = {
  OffRamp: {
    report: 110_456,
    executeTransaction: {
      TWO_TOKENS: {
        FEES: {
          CONTRACT_RECEIVER: 536_304,
          EOA_RECEIVER: 235_608,
        },
        NO_FEES: {
          CONTRACT_RECEIVER: 489_646,
          EOA_RECEIVER: 249_896,
        },
      },
    },
  },
  OnRamp: {
    requestCrossChainSend: {
      MESSAGE_ONLY: 190_394,
      ONE_TOKEN: 213_081,
      TWO_TOKENS: 323_390,
      THREE_TOKENS: 433_687,
    },
  },
}
