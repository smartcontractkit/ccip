# Rebalancer O11y

## On Chain Events

### Rebalancer.sol Events:
- `event LiquidityTransferred(uint64 indexed ocrSeqNum, uint64 indexed fromChainSelector, uint64 indexed - toChainSelector, address to, uint256 amount, bytes bridgeSpecificData, bytes bridgeReturnData);`
  - signals _rebalanceLiquidity OR _receiveLiquidity was called.
  - Liquidity was transferred to another chain.  
- `event LiquidityAddedToContainer(address indexed provider, uint256 indexed amount);`
  - signals addLiquidity was called. 
  - This event is emitted when liquidity is added to the container. Could technically be from anyone.
- `event LiquidityRemovedFromContainer(address indexed remover, uint256 indexed amount);`
  - signals removeLiquidity was called. 
  - This event is emitted when liquidity is removed from the container. Owner only.
- `event LiquidityContainerSet(address indexed newLiquidityContainer);`
  - signals setLocalLiquidityContainer was called.
  - config setting
- `event MinimumLiquiditySet(uint256 oldBalance, uint256 newBalance);`
  - signals setMinimumLiquidity was called.
  - config setting
- `event CrossChainRebalancerSet(uint64 indexed remoteChainSelector, IBridgeAdapter localBridge, address remoteToken, address remoteRebalancer, bool enabled);`
  - signals _setCrossChainRebalancer was called.
  - config setting
- `event FinalizationFailed(uint64 indexed ocrSeqNum, uint64 indexed remoteChainSelector, bytes bridgeSpecificData, bytes reason);`
  - signals _receiveLiquidity was called.
  - Failed to finalize a withdrawal. this could mean that the withdrawal was already finalized or that the withdrawal failed.


### Bridge Adapter Events:
There is no events from bridge adapters.

## Atlas Tables

### Rebalancer.sol Tables:

Follow the same pattern as Automation event tables.
```postgresql
CREATE TABLE IF NOT EXISTS "{{ network }}".rebalancer_events
(
    chain_id          numeric   NOT NULL,
    event_id          text      NOT NULL,
    event_name        text      NOT NULL,
    inputs            jsonb     NOT NULL,
    raw_log           text      NOT NULL,
    block_number      numeric   NOT NULL,
    transaction_hash  text      NOT NULL,
    transaction_index integer   NOT NULL,
    block_hash        text      NOT NULL,
    log_index         integer   NOT NULL,
    removed           boolean   NOT NULL,
    block_timestamp   timestamp NOT NULL,
    contract_address  citext      NOT NULL,
    created_at        timestamp NOT NULL DEFAULT now(),
    updated_at        timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY(event_id, transaction_hash, transaction_index, block_hash, log_index)

);
CREATE INDEX IF NOT EXISTS idx_contract_address_on_rebalancer ON "{{ network }}".rebalancer_events (contract_address);
CREATE INDEX IF NOT EXISTS idx_event_name_on_rebalancer ON "{{ network }}".rebalancer_events (event_name);
```

example query:
```postgresql
SELECT * FROM "{{ network }}".rebalancer_events WHERE event_name = 'LiquidityAddedToContainer';
```

A second table for transfers that is in a more readable format using https://github.com/smartcontractkit/chain-selectors to convert the chain selector to a chain name.
```postgresql
CREATE TABLE IF NOT EXISTS "{{ network }}".rebalancer_transfers
(
    chain_id             numeric   NOT NULL,
    event_id             text      NOT NULL,
    event_name           text      NOT NULL,
    ocr_seq_num          integer   NOT NULL,
    from_chain           citext    NOT NULL, -- ie ethereum-mainnet-arbitrum-1
    to_chain             citext    NOT NULL, -- ie ethereum-mainnet
    to_address           citext    NOT NULL,
    amount               numeric   NOT NULL,
    bridge_specific_data bytea     NOT NULL,
    bridge_return_data   bytea     NOT NULL,
    block_number         numeric   NOT NULL,
    transaction_hash     text      NOT NULL,
    transaction_index    integer   NOT NULL,
    block_hash           text      NOT NULL,
    log_index            integer   NOT NULL,
    removed              boolean   NOT NULL,
    block_timestamp      timestamp NOT NULL,
    contract_address     citext    NOT NULL,
    created_at           timestamp NOT NULL DEFAULT now(),
    updated_at           timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY (event_id, transaction_hash, transaction_index, block_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_rebalancer_contract_address ON "{{ network }}".rebalancer_events (contract_address);
CREATE INDEX IF NOT EXISTS idx_rebalancer_event_name ON "{{ network }}".rebalancer_events (event_name);
CREATE INDEX IF NOT EXISTS idx_rebalancer_from_chain ON "{{ network }}".rebalancer_events (from_chain);
CREATE INDEX IF NOT EXISTS idx_rebalancer_to_chain ON "{{ network }}".rebalancer_events (to_chain);
```

example query:
```postgresql
SELECT * FROM "{{ network }}".rebalancer_events WHERE from_chain = 'ethereum-mainnet-arbitrum-1';
SELECT * FROM "{{ network }}".rebalancer_events WHERE to_chain = 'ethereum-mainnet';
SELECT * FROM "{{ network }}".rebalancer_events WHERE from_chain = 'ethereum-mainnet-arbitrum-1' AND to_chain = 'ethereum-mainnet';
```

## Custom Telemetry Metrics (Maybe?)
Especially since all nodes will be run by us, we know telemetry will be turned on.

### Plugin Metrics
- rebalancer_total_liquidity
  - Description: The total liquidity in the rebalancer system across all chains in view.
  - Type: Gauge
  - Additional Labels: token
- rebalancer_pool_liquidity
  - Description: The liquidity in a specific chains pool.
  - Type: Gauge
  - Additional Labels: network_name, token

## Benthos
Even if we do custom telemetry we will need a benthos stream to read in the telemetry. Or we can just do it via Benthos without telemetry. Or do both and have redundancy in our liquidity data. Either way we will need at least 2 benthos stream to collect the event data.

### Metrics
- rebalancer_total_liquidity
  - We could have a stream calculating the total liquidity in the rebalancer system across all chains in view.
  - same as custom telemetry metric above. this could be as backup in case nodes don't report or DON goes down or just as an extra data point
- rebalancer_pool_liquidity (same as custom telemetry metric above)
  - We could have a stream calculating the liquidity in a specific chains pool.
  - same as custom telemetry metric above. this could be as backup in case nodes don't report or DON goes down or just as an extra data point
- rebalancer_events
  - Description: All events from the rebalancer contract. Used to visualize the sequence/rate of events.
  - Type: Counter
  - Additional Labels: event_name
~~- sent/received liquidity?
  - would this be graphable? since the metric may change multiple times quickly we may not get an accurate representation of the data.
  - use to/from as labels to see the direction of the liquidity.
  - need to think on this one more.~~
- rebalancer_transfers
  - Description: All transfers from the rebalancer contract. Used to visualize the rate of transfers.
  - Type: Counter
  - Additional Labels: from_chain, to_chain

### Streams
- A stream to handle all the event counting. Will also be used to output the data to postgres.
- A stream for transfers and a processor to convert the chain selector to a chain name.
- A stream to fetch all pool liquidity and then calculate the total liquidity.
- A stream to collect telemetry data if available.
