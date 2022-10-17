class InheritedConstructorArgsNotInContractDefinition {
  constructor(reporter, config) {
    this.ruleId = 'inherited-constructor-args-not-in-contract-definition'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { baseContracts } = ctx
    for (let baseContract of baseContracts) {
      if (baseContract.arguments.length > 0) {
        this.reporter.error(
          ctx,
          this.ruleId,
          `Inherited contract constructor arguments for ${baseContract.baseName.namePath} should be passed in the constructor definition, not the contract definition`,
        )
      }
    }
  }
}

module.exports = InheritedConstructorArgsNotInContractDefinition
