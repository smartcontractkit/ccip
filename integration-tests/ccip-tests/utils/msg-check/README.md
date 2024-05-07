# Msg Checker

This tool queries detailed information about messages from CCIP Explorer using the message ID. It provides message stats, including timestamps for send, commit, and receipt phases, as well as the duration between these phases. The tool also resolves the fee token address to a more human-readable form using a provided token mapping file.

## Configuration

### Token Mapping

A tokens.toml file is used to map token addresses to their corresponding human-readable names. Here is an example of how token mappings are defined:

```toml
# Ethereum Mainnet Tokens
"0x514910771AF9Ca656af840dff83E8264EcF986CA" = "LINK"
"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" = "WETH"
```

## Usage

Provide a message id to get details:

```bash
go run . -msg=<message_id>
```

## Output

Output
After executing the script, the console will display detailed information about the queried message, including:

* Message ID
* Fee token (resolved to a human-readable name if possible)
* Timestamps for various stages
* Duration between different stages
* Additionally, a messages.csv file will be generated in the current directory, containing the same detailed information in a structured format suitable for further analysis or record-keeping.