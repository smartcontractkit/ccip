## Chaos testing
To run experiments from CLI download latest `havoc` [release](https://github.com/smartcontractkit/havoc/releases/tag/v0.2.4)

Check CCIP `havoc` [configuration](./havoc.toml)

## Automated runs
Check this [file](../load/ccip_test.go) and `TestLoadCCIPStableRPSWithChaos`

## Running the tests from CLI
Monkey mode
```
./havoc -c havoc.toml run
```

Experiment mode
```
./havoc -c havoc.toml apply
```

## Grafana integration
Add environment variables
```
export HAVOC_LOG_LEVEL="info"
export GRAFANA_URL="..."
export GRAFANA_TOKEN="..."
```
`dashboard_uuids` are copied from your target dashboards (JSON Model tab in Grafana)
