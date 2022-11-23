# Foundry Test Guidelines

We're using Foundry to test our CCIP smart contracts here. This enables us to test in Solidity. If you need to add tests for anything outside the CCIP contracts, please write them in hardhat (for the time being).

## Directory Structure

Mimic the source contract file structure as closely as possible. Example:

`./offRamp/ge/SomeOffRamp.sol` should have a test contract `./test/offRamp/ge/SomeOffRamp.t.sol`.

## Test File Structure

Break the test file down into multiple contracts, each contract testing a specific function inside the source contract.

For Example, here's a source contract `SomeOffRamp`:

```
contract SomeOffRamp {

  constructor() {
    ... set some state
  }

  function firstFunction() public {
    ...
  }

  function theNextFunction() public {
    ...
  }

  function _anInternalFunction() internal {
    ...
  }
}
```

Our test file `SomeOffRamp.t.sol` should be structured like this:

```
contract SomeOffRamp_constructor {
  // constructor state setup tests here
}

contract SomeOffRamp_firstFunction {
  // first function tests here
}

contract SomeOffRamp_theNextFunction {
  // tests here too...
}

contract SomeOffRamp_anInternalFunction {
  // This function will require a helper contract
  // to expose it.
}
```

## Test Structure

Inside each test contract (described above), group tests into `Success` and `Reverts` with inline comments:

```
contract SomeOffRamp_firstFunction {
  // Success

  function testZeroValueSuccess() public {
    ...
  }

    ...


  // Reverts

  function testOwnerReverts() public {
    // test that an ownable function reverts when not called by the owner
    ...
  }

  ...

}
```

Function naming should follow this structure:

`test{description of test}[Success|Reverts]`

Try to cover all the code paths present in each function being tested. In most cases, this will result in many more failure tests than success tests.

If a test file requires a complicated setUp, or if it requires many helper functions (like `_generateAMessageWithNoTokensStruct()`), create a separate file to perform this setup in. Using the example above, `SomeOffRampSetup.t.sol`. Inherit this and call the setUp function in the test file.
