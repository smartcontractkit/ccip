# example usage: make test_smoke_ccip_simulated_local image=chainlink-ccip tag=latest testname=TestSmokeCCIPForBidirectionalLane
.PHONY: test_smoke_ccip_simulated_local
test_load_ccip_simulated_k8:
	DATABASE_URL="postgresql://postgres:node@localhost:5432/chainlink_test?sslmode=disable" \
    # the test will run in simulated networks
    SELECTED_NETWORKS="SIMULATED,SIMULATED_1,SIMULATED_2"  \
    # the test will not run in local docker
    CCIP_DEPLOY_ON_LOCAL=False  \
    # the test will create new environment with new contracts, chainlink nodes and jobs
    CCIP_TESTS_ON_EXISTING_DEPLOYMENT=False  \
    # the test will use SELECTED_NETWORKS, CCIP_NO_OF_NETWORKS, CCIP_NO_OF_LANES_PER_PAIR to create network pairs, no need to set CCIP_NETWORK_PAIRS
    CCIP_NETWORK_PAIRS="" \
    # the test will create 10 chains, if the chains are not simulated, the test expects 10 chains to be available in SELECTED_NETWORKS value
    # if simulated, the SELECTED_NETWORKS should be set as SIMULATED,SIMULATED_1,SIMULATED_2, the rest of the chains will be created by the test
    # by mimicking the values of SIMULATED_1 and using one of the availale chain ids from test chainids in chain selector repo
    CCIP_NO_OF_NETWORKS=10  \
    # the test will create 2 routers for each chains
    CCIP_NO_OF_LANES_PER_PAIR=2  \
    # the test will use this funding amount (native) for each chainlink node, keeping it high for running long and high frequency load tests
    CCIP_CHAINLINK_NODE_FUNDING=1000 \
    # the k8 env will be destroyed after 24h
    CCIP_KEEP_ENV_TTL=24h \
    # the test will use data-only messages
    CCIP_MSG_TYPE=WithoutToken \
    # the image and version to be used for chainlink nodes. The image can be located from
    CHAINLINK_IMAGE=$(image)  \
    CHAINLINK_VERSION=$(tag)  \
    ENV_JOB_IMAGE=""  \
	go test -timeout 24h -count=1 -v -run ^TestLoadCCIPStableRequestTriggeringWithNetworkChaos$$ ./load


# example usage: make test_smoke_ccip_simulated_local image=chainlink-ccip tag=latest testname=TestSmokeCCIPForBidirectionalLane
.PHONY: test_smoke_ccip_simulated_local
test_smoke_ccip_simulated_local:
	DATABASE_URL="postgresql://postgres:node@localhost:5432/chainlink_test?sslmode=disable" \
    SELECTED_NETWORKS="SIMULATED,SIMULATED_1,SIMULATED_2"  \
    CCIP_DEPLOY_ON_LOCAL=True  \
    CCIP_NETWORK_PAIRS="" \
    CCIP_NO_OF_NETWORKS=2  \
    CCIP_NO_OF_LANES_PER_PAIR=1  \
    CCIP_MSG_TYPE=WithToken \
    CHAINLINK_IMAGE=$(image)  \
    CHAINLINK_VERSION=$(tag)  \
    CCIP_TESTS_ON_EXISTING_DEPLOYMENT=False  \
    ENV_JOB_IMAGE=""  \
	go test -timeout 24h -count=1 -v -run ^$(testname)$$ ./smoke

# example usage: make test_smoke_ccip_simulated_local_besu image=chainlink tag=latest
.PHONY: test_smoke_ccip_simulated_local_besu
test_smoke_ccip_simulated_local_besu:
	DATABASE_URL="postgresql://postgres:node@localhost:5432/chainlink_test?sslmode=disable" \
    SELECTED_NETWORKS="SIMULATED,SIMULATED_BESU_NONDEV_1,SIMULATED_BESU_NONDEV_2"  \
    CCIP_DEPLOY_ON_LOCAL=True  \
    CCIP_NO_OF_NETWORKS=2  \
    CCIP_MSG_TYPE=WithToken \
    CHAINLINK_IMAGE=$(image)  \
    CHAINLINK_VERSION=$(tag)  \
    CCIP_TESTS_ON_EXISTING_DEPLOYMENT=False  \
    ENV_JOB_IMAGE=""  \
	go test -timeout 24h -count=1 -v -run ^TestSmokeCCIPForBidirectionalLane$$ ./smoke
