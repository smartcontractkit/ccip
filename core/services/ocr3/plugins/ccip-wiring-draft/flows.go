package ccip_wiring_draft

import "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip-wiring-draft/roleuis"

func initCCIP(
	capabilityAuthor roleuis.CapabilityAuthor,
	nodeOperator roleuis.NodeOperator,
) {
	const ccipCapabilityID = "ccip"

	// capability author initializes the ccip capability by creating the config contract
	// and adding the new capability to the registry
	{
		capabilityCfgAddr, _ := capabilityAuthor.DeployConfigContract(
			ccipCapabilityID, []byte(
				`{
					"f_chain": {eth: 5, arb: 4},
					"transmitters": {1, 4, 5},
					"observationTimeout": 15s,
					"rpcTimeout": 10s,
					"chain_cfg": {
						"eth": {
							"contract": "0x1234...",
							"offramp": "0x123",
						}
					}
				}`,
			))

		_ = capabilityAuthor.AddCapabilityToRegistry(ccipCapabilityID, capabilityCfgAddr, []byte(`{
			string type = "ccip-commit" / "ccip-exec"
			string version = "3.0.0"
			bool removed = false
			ResponseType responseType = REPORT
			ConfigurationSelector[] configurationSelectors = []{
			address contract = "0x1234...", bytes4 functionSelector = bytes4(
				"function getConfiguration(uint256 configurationId)"
			)}
		}`))
		// 'CapabilityAdded(ccip)' is emitted
	}

	// node operator can add the new capability in his capabilities stack
	{
		capabilityID, _ := nodeOperator.AddCapability(
			ccipCapabilityID, []byte(
				`{i support arbitrum ethereum and polygon}`,
			))

		_ = nodeOperator.PublishNode(capabilityID, []byte(
			`{publish my node that runs this capability}`,
		)) // 'NodeAdded(nodeID)' is emitted
	}

	// update config and let dons know

}
