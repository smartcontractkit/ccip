class PrefixStorageVariablesWithSUnderscore {
  static applicableVisibilities = ['private', 'internal'] // public variables are ignored

  constructor(reporter, config) {
    this.ruleId = 'prefix-storage-variables-with-s-underscore'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { subNodes } = ctx
    for (let subNode of subNodes) {
      const { type } = subNode
      if (type === 'StateVariableDeclaration') {
        for (let variable of subNode.variables) {
          const { type, visibility, isDeclaredConst, isImmutable, name } =
            variable
          if (
            type === 'VariableDeclaration' &&
            !isDeclaredConst && // const variables ignored
            !isImmutable && // immutable variables ignored
            PrefixStorageVariablesWithSUnderscore.applicableVisibilities.includes(
              visibility,
            ) &&
            !name.startsWith('s_')
          ) {
            this.reporter.error(
              ctx,
              this.ruleId,
              `Private / internal variable ${name} is not prefixed with s_`,
            )
          }
        }
      }
    }
  }
}

module.exports = PrefixStorageVariablesWithSUnderscore
