# CCIP Tests

Here lives the integration tests for ccip, utilizing our [chainlink-testing-framework](https://github.com/smartcontractkit/chainlink-testing-framework) and [integration-tests](https://github.com/smartcontractkit/ccip/integration-tests)

## Running the tests

### Setting up test inputs :

In order to run the tests the first step is to set up the test inputs. There are two kinds of inputs -
1. Generic test input set via TOML - If no specific input is set the tests will run with default inputs mentioned in [default.toml](./testconfig/tomls/default.toml)
2. Secrets set via env variables. Please refer to [secrets.toml](./testconfig/secrets.env) for the list of env variables that need to be set.

If you want to override the default inputs, you need to set an env var `BASE64_TEST_CONFIG_OVERRIDE` containing the base64 encoded TOML file content.
For example, if you want to override the `Networks` input in test and want to run your test on `avalanche testnet` and `arbitrum goerli` network, you can create a TOML file with the following content:
```toml
[CCIP]

[CCIP.Env]
Networks = ['AVALANCHE_FUJI', 'ARBITRUM_GOERLI']
```
and then encode it using `base64` command and set the env var `BASE64_TEST_CONFIG_OVERRIDE` with the encoded content.
```bash
export BASE64_TEST_CONFIG_OVERRIDE=$(base64 -w 0 < path-to-toml-file)
```

Alternatively, you can also use the make command to invoke a go script to do the same.
```bash
## if overridestring is set, override_toml is ignored
make override_config overridestring="<overridden config string>" override_toml="<the toml file with overridden config string>" env="<.env file with BASE64_TEST_CONFIG_OVERRIDE value>"
```

In order to set the secrets, you need to set the env vars mentioned in [secrets.toml](./testconfig/secrets.env) file and source the file.  
```bash
source ./testconfig/secrets.env
```

Please note that the secrets.env should not be checked in to the repo and should be kept locally.
You can run the command to ignore the changes to the file.
```bash
git update-index --skip-worktree ./testconfig/secrets.env
```



To run the tests with
There are two ways to run the tests:
1. Using local docker containers
2. Using remote kubernetes cluster

### Using local docker containers

In order to run the tests locally, you need to have docker installed and running on your machine.
