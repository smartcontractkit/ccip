/* eslint-disable @typescript-eslint/no-var-requires */
const ExplicitImports = require('./rules/explicitImports.js')
const PrefixPrivateFunctionsWithUnderscore = require('./rules/prefixPrivateFunctionsWithUnderscore.js')
const PrefixInternalFunctionsWithUnderscore = require('./rules/prefixInternalFunctionsWithUnderscore.js')
const InheritedConstructorArgsNotInContractDefinition = require('./rules/inheritedConstructorArgsNotInContractDefinition.js')
const AllCapsConstantStorageVariables = require('./rules/allCapsConstantStorageVariables.js')
const PrefixStorageVariablesWithSUnderscore = require('./rules/prefixStorageVariablesWithSUnderscore.js')
const PrefixImmutableVariablesWithI = require('./rules/prefixImmutableVariablesWithI.js')
const NoRequireStatements = require('./rules/noRequireStatements.js')

module.exports = [
  ExplicitImports,
  PrefixPrivateFunctionsWithUnderscore,
  PrefixInternalFunctionsWithUnderscore,
  InheritedConstructorArgsNotInContractDefinition,
  AllCapsConstantStorageVariables,
  PrefixStorageVariablesWithSUnderscore,
  PrefixImmutableVariablesWithI,
  NoRequireStatements,
]
