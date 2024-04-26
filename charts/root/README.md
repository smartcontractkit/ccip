# Root chart
Root chart dedicated for integration testing

## Development
### Updating dependencies
`helm dependency update`

### Testing 
#### Render template locally
```
helm template app -f values.yaml -n crib-example . --output-dir .rendered \
--set=ingress.baseDomain="$DEVSPACE_INGRESS_BASE_DOMAIN" \
--set=ccip-scripts.ccipScriptsDeployment.enabled=true
```

### Testing using local version of crib-chainlink-cluster chart

1) In the root chart dependencies replace the remote chart reference with the local one for example:

```yaml
  - name: crib-chainlink-cluster
    version: 0.5.1
    repository: "file://../../../infra-charts/crib-chainlink-cluster"
```

2) Update dependencies in the root chart
```
helm dependency update
```

3) All set! Now you can render chart locally or test it with devspace