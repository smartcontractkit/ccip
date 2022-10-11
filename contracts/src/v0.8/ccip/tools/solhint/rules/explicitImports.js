class ExplicitImports {
  constructor(reporter, config) {
    this.ruleId = 'explicit-imports'
    this.reporter = reporter
    this.config = config
  }

  ImportDirective(ctx) {
    const { symbolAliases, path } = ctx
    if (symbolAliases == null) {
      this.reporter.error(
        ctx,
        this.ruleId,
        `Import "${path}" must explicitly import types from the imported file.`,
      )
    }
  }
}

module.exports = ExplicitImports
