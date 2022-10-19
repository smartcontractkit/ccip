class NoBlockSingleIfReverts {
  constructor(reporter, config) {
    this.ruleId = 'no-block-single-if-reverts'
    this.reporter = reporter
    this.config = config
  }

  ContractDefinition(ctx) {
    try {
      const { subNodes } = ctx
      for (let subNode of subNodes) {
        this.checkNodeForBlockedSingleRevertStatement(subNode)
      }
    } catch (e) {
      console.error(e)
    }
  }

  checkNodeForBlockedSingleRevertStatement(node) {
    const { type, body } = node
    if (type === 'IfStatement') {
      const { trueBody, falseBody } = node
      if (
        trueBody.statements &&
        trueBody.statements.length === 1 &&
        trueBody.statements[0].type === 'RevertStatement'
      ) {
        this.reporter.error(
          trueBody.statements[0],
          this.ruleId,
          'If statements with only a single revert expression must not be in a block.',
        )
      } else {
        if (trueBody && trueBody.statements)
          this.checkBodyStatements(trueBody.statements)
        if (falseBody && falseBody.statements)
          this.checkBodyStatements(falseBody.statements)
      }
    } else if (body && body.statements) {
      this.checkBodyStatements(body.statements)
    }
  }

  checkBodyStatements(statements) {
    for (const statement of statements) {
      this.checkNodeForBlockedSingleRevertStatement(statement)
    }
  }
}

module.exports = NoBlockSingleIfReverts
