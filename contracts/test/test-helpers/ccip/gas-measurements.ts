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
          NO_TOKENS: 2_408_105,
          ONE_TOKEN: 3_625_481,
        },
      },
    },
  },
  OnRamp: {
    requestCrossChainSend: {
      MESSAGE_ONLY: 188_431,
      ONE_TOKEN: 211_112,
      TWO_TOKENS: 321_445,
      THREE_TOKENS: 431_754,
    },
  },
}
