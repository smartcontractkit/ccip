class PrefixImmutableVariablesWithI {
  constructor(reporter, config) {
    this.ruleId = 'prefix-immutable-variables-with-i'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { subNodes } = ctx
    for (let subNode of subNodes) {
      const { type } = subNode
      if (type === 'StateVariableDeclaration') {
        for (let variable of subNode.variables) {
          const { type, visibility, isImmutable, name } = variable
          if (
            type === 'VariableDeclaration' &&
            visibility &&
            isImmutable && // immutable variables only
            !name.startsWith('i_')
          ) {
            this.reporter.error(
              ctx,
              this.ruleId,
              `Immutable variable ${name} is not prefixed with 'i_'`,
            )
          }
        }
      }
    }
  }
}

module.exports = PrefixImmutableVariablesWithI
