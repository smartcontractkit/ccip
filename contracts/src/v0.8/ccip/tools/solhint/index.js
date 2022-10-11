/* eslint-disable @typescript-eslint/no-var-requires */
const ExplicitImports = require('./rules/explicitImports.js')
const PrefixPrivateFunctionsWithUnderscore = require('./rules/prefixPrivateFunctionsWithUnderscore.js')

module.exports = [ExplicitImports, PrefixPrivateFunctionsWithUnderscore]
