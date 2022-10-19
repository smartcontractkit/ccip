class NoRequireStatements {
  constructor(reporter, config) {
    this.ruleId = 'no-require-statements'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    const { subNodes } = ctx
    for (let subNode of subNodes) {
      if (
        subNode.type === 'FunctionDefinition' ||
        subNode.type === 'ModifierDefinition'
      ) {
        this.checkForRequireStatement(JSON.parse(JSON.stringify(subNode)))
      }
    }
  }

  checkForRequireStatement(node) {
    switch (node.type) {
      case 'ExpressionStatement':
        if (
          node.expression &&
          node.expression.expression &&
          node.expression.expression.name === 'require'
        ) {
          this.reporter.error(
            node,
            this.ruleId,
            `Use custom errors instead of revert statements`,
          )
        }
        return
      case 'IfStatement':
        if (node.trueBody) this.checkForRequireStatement(node.trueBody)
        if (node.falseBody) this.checkForRequireStatement(node.falseBody)
        if (!node.trueBody && !node.falseBody) {
          console.log(node)
        }
        return
      default:
        if (!node.body) return
        for (const statement of node.body.statements) {
          this.checkForRequireStatement(statement)
        }
    }
  }
}

module.exports = NoRequireStatements
