-- +goose Up

CREATE EXTENSION pgcrypto;

INSERT INTO evm.logs(
    evm_chain_id,
    log_index,
    block_hash,
    block_number,
    address,
    event_sig,
    topics,
    tx_hash,
    data,
    created_at,
    block_timestamp
)
SELECT
    1337 AS evm_chain_id,
    gs AS log_index,
    rb1 AS block_hash,
    gs AS block_number,
    rb2 AS address,
    rb1 AS event_sig,
    ARRAY[rb1, rb2, rb1, rb2] AS topics,
    rb1 AS tx_hash,
    gen_random_bytes(128) AS data,
  timestamp '2023-01-01 00:00:00' + random() * (timestamp '2023-12-31 23:59:59' - timestamp '2023-01-01 00:00:00') AS created_at,
  timestamp '2023-01-01 00:00:00' + random() * (timestamp '2023-12-31 23:59:59' - timestamp '2023-01-01 00:00:00') AS block_timestamp
FROM generate_series(1, 200000) gs
    LEFT JOIN LATERAL (
    SELECT gen_random_bytes(32) AS rb1,
    gen_random_bytes(32) as rb2) as foo ON true;


INSERT INTO evm.logs(
    evm_chain_id,
    log_index,
    block_hash,
    block_number,
    address,
    event_sig,
    topics,
    tx_hash,
    data,
    created_at,
    block_timestamp
)
SELECT
    2337 AS evm_chain_id,
    gs AS log_index,
    rb1 AS block_hash,
    gs AS block_number,
    rb2 AS address,
    rb1 AS event_sig,
    ARRAY[rb1, rb2, rb1, rb2] AS topics,
    rb1 AS tx_hash,
    gen_random_bytes(128) AS data,
  timestamp '2023-01-01 00:00:00' + random() * (timestamp '2023-12-31 23:59:59' - timestamp '2023-01-01 00:00:00') AS created_at,
  timestamp '2023-01-01 00:00:00' + random() * (timestamp '2023-12-31 23:59:59' - timestamp '2023-01-01 00:00:00') AS block_timestamp
FROM generate_series(1, 200000) gs
    LEFT JOIN LATERAL (
    SELECT gen_random_bytes(32) AS rb1,
    gen_random_bytes(32) as rb2) as foo ON true;

DROP EXTENSION pgcrypto;

-- +goose Down

delete from evm.logs where evm_chain_id = 1337;
delete from evm.logs where evm_chain_id = 2337;