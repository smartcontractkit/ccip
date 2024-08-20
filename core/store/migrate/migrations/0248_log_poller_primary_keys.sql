-- +goose Up

ALTER TABLE evm.logs DROP CONSTRAINT logs_pkey;
ALTER TABLE evm.logs ADD COLUMN id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY;
ALTER TABLE evm.log_poller_blocks  DROP CONSTRAINT log_poller_blocks_pkey;
ALTER TABLE evm.log_poller_blocks ADD COLUMN id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY;

-- +goose Down

ALTER TABLE evm.log_poller_blocks DROP COLUMN id;
ALTER TABLE evm.log_poller_blocks ADD PRIMARY KEY (block_number, evm_chain_id);
ALTER TABLE evm.logs DROP COLUMN id;
ALTER TABLE evm.logs ADD PRIMARY KEY (block_hash, log_index, evm_chain_id);
