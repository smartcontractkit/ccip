/* eslint-disable @typescript-eslint/no-var-requires */
const ExplicitImports = require('./rules/explicitImports.js')
const PrefixPrivateFunctionsWithUnderscore = require('./rules/prefixPrivateFunctionsWithUnderscore.js')
const PrefixInternalFunctionsWithUnderscore = require('./rules/prefixInternalFunctionsWithUnderscore.js')

module.exports = [
  ExplicitImports,
  PrefixPrivateFunctionsWithUnderscore,
  PrefixInternalFunctionsWithUnderscore,
]
