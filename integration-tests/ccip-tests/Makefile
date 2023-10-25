## To Override the default config,
# example usage: make override_config overrideconfig="" override_toml=../config/config.toml env=../.env
.PHONY: override_config
override_config:
	cd ./testconfig/override && \
	go run . --overridewith=$(overrideconfig) --path=$(override_toml) --output=$(env) && \
	cd ../..

# example usage: make test_load_ccip_simulated_k8 image=chainlink-ccip tag=latest testimage=chainlink-ccip-tests:latest testname=TestLoadCCIPStableRequestTriggeringWithNetworkChaos
.PHONY: test_load_ccip
test_load_ccip:override_config
	source ./testconfig/override/$(env) && \
    CHAINLINK_IMAGE=$(image)  \
    CHAINLINK_VERSION=$(tag)  \
    ENV_JOB_IMAGE=$(testimage)  \
    TEST_SUITE=load \
    TEST_ARGS="-test.timeout 900h" \
    DETACH_RUNNER=true \
	go test -timeout 24h -count=1 -v -run ^$(testname)$$ ./load


# example usage: make test_smoke_ccip_simulated_local image=chainlink-ccip tag=latest testimage=chainlink-ccip-tests:latest testname=TestSmokeCCIPForBidirectionalLane override_toml=../config/config.toml env=test-config.env
.PHONY: test_smoke_ccip
include test-config.env
test_smoke_ccip: override_config
	source ./testconfig/override/$(env) && \
    CHAINLINK_IMAGE=$(image)  \
    CHAINLINK_VERSION=$(tag)  \
    ENV_JOB_IMAGE=$(testimage)  \
    TEST_SUITE=smoke \
    TEST_ARGS="-test.timeout 900h" \
    DETACH_RUNNER=true \
	go test -timeout 24h -count=1 -v -run ^$(testname)$$ ./smoke

# example usage: make test_smoke_ccip_simulated_local_besu image=chainlink-ccip tag=latest override_toml=../config/config.toml env=test-config.env
.PHONY: test_smoke_ccip_simulated_local_besu
include test-config.env
overrideconfig := "CCIP.Env.Networks = ['SIMULATED_BESU_NONDEV_1', 'SIMULATED_BESU_NONDEV_2']"
test_smoke_ccip_simulated_local_besu: override_config
	source ./testconfig/override/$(env) && \
  	CHAINLINK_IMAGE=$(image)  \
  	CHAINLINK_VERSION=$(tag)  \
    ENV_JOB_IMAGE=""  \
	go test -timeout 24h -count=1 -v -run ^TestSmokeCCIPForBidirectionalLane$$ ./smoke
