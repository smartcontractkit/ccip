# CCIP Explorer Search
This tool facilitates querying CCIP transactions and generating a CSV report and console output that includes details such as network names, transaction hashes, fee tokens, and URLs to view transactions on various blockchain explorers.

## Command Line Flags

The sender's address is required:
* -sender: Sender's address (required)
* -receiver: Receiver's address
* -source: Source network shortcode
* -dest: Destination network shortcode
* -messageId: Unique identifier for the message (probably should not be an option)
* -feeToken: Fee token address

Additional flags for pagination: 
* -first
* -offset

## Example Command

```bash
go run . -sender=<address> -receiver=<address> -source=ETH -dest=POLY -first=5
```

## Output
After executing the script, it will generate a transactions.csv file in the current directory. This file includes details on each transaction queried, such as source and destination networks, fee token, and URLs for transaction verification on network-specific explorers.

Transactions will also be printed in the terminal with markdown formatted links which can be inserted into markdown or confluence docs.

## Configuration

### Network Mappings

Predefined network name mappings are used to identify networks. These mappings are hard-coded within the script:

```toml
var networkNameMapping = map[string]string{
    "POLY": "polygon-mainnet",
    "BSC": "binance_smart_chain-mainnet",
    "ETH": "ethereum-mainnet",
    // Add more mappings here...
}
```
### networks.toml Configuration

The networks.toml file has network-specific configurations such as explorer URLs and token address mappings. Here's an example structure of the networks.toml:

```toml
[networks."polygon-mainnet"]
url = "https://polygonscan.com/tx/"
tokens = { "0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270" = "WMATIC", "0xb0897686c545045afc77cf20ec7a532e3120e0f1" = "LINK" }
```