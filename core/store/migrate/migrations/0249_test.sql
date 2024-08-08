-- +goose Up

drop index if exists evm.logs_idx_topic_two;

drop index if exists evm.idx_evm_logs_ordered_by_block_and_created_at;

CREATE INDEX logs_by_block_timestamp
    ON evm.logs (address, event_sig, evm_chain_id, block_timestamp);

-- +goose Down

drop index if exists evm.logs_by_block_timestamp;
