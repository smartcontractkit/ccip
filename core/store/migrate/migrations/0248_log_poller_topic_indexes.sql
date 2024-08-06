-- +goose Up

create index evm_logs_idx_topic_two_extended
    on evm.logs (address, event_sig, evm_chain_id, (topics[2]));

-- +goose Down

drop index if exists evm.evm_logs_idx_topic_two_extended;
