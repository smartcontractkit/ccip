class PrefixPrivateFunctionsWithUnderscore {
  constructor(reporter, config) {
    this.ruleId = 'prefix-private-functions-with-underscore'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { subNodes } = ctx
    for (let subNode of subNodes) {
      const { type, visibility, name } = subNode
      if (
        type === 'FunctionDefinition' &&
        visibility === 'private' &&
        !name.startsWith('_')
      ) {
        this.reporter.error(
          ctx,
          this.ruleId,
          `Private function ${name} is not prefixed with underscore (_)`,
        )
      }
    }
  }
}

module.exports = PrefixPrivateFunctionsWithUnderscore
