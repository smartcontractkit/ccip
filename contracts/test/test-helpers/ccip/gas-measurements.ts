export const GAS = {
  OffRamp: {
    report: 110_456,
    executeTransaction: {
      ONE_MESSAGE: {
        TWO_TOKENS: {
          FEES: {
            CONTRACT_RECEIVER: 511_884,
            EOA_RECEIVER: 235_591,
          },
          NO_FEES: {
            CONTRACT_RECEIVER: 465_214,
            EOA_RECEIVER: 249_891,
          },
        },
      },
      TEN_MESSAGES: {
        NO_FEES: {
          NO_TOKENS: 2_408_163,
          ONE_TOKEN: 3_625_293,
        },
      },
    },
  },
  OnRamp: {
    requestCrossChainSend: {
      MESSAGE_ONLY: 133_548,
      ONE_TOKEN: 133_512,
      TWO_TOKENS: 202_574,
      THREE_TOKENS: 271_600,
    },
  },
}
