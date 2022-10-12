class PrefixInternalFunctionsWithUnderscore {
  constructor(reporter, config) {
    this.ruleId = 'prefix-internal-functions-with-underscore'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { subNodes } = ctx
    for (let subNode of subNodes) {
      const { type, visibility, name } = subNode
      if (
        type === 'FunctionDefinition' &&
        visibility === 'internal' &&
        !name.startsWith('_')
      ) {
        this.reporter.error(
          ctx,
          this.ruleId,
          `Internal function ${name} is not prefixed with underscore (_)`,
        )
      }
    }
  }
}

module.exports = PrefixInternalFunctionsWithUnderscore
