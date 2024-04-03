-- +goose Up

-- Backup logs table
ALTER TABLE evm.logs RENAME TO logs_backup;

-- Recreate logs table
create table evm.logs
(
    evm_chain_id    numeric(78)              not null,
    log_index       bigint                   not null,
    block_hash      bytea                    not null,
    block_number    bigint                   not null
        constraint logs_block_number_check
            check (block_number > 0),
    address         bytea                    not null,
    event_sig       bytea                    not null,
    topics          bytea[]                  not null,
    tx_hash         bytea                    not null,
    data            bytea                    not null,
    created_at      timestamp with time zone not null,
    block_timestamp timestamp with time zone not null,
    primary key (block_hash, log_index, evm_chain_id)
) PARTITION BY HASH (evm_chain_id);

create index evm_logs_partitioned_idx
    on evm.logs (evm_chain_id, block_number, address, event_sig);

create index evm_logs_partitioned_idx_data_word_one
    on evm.logs ("substring"(data, 1, 32));

create index evm_logs_partitioned_idx_data_word_two
    on evm.logs ("substring"(data, 33, 32));

create index evm_logs_partitioned_idx_data_word_three
    on evm.logs ("substring"(data, 65, 32));

create index evm_logs_partitioned_idx_data_word_four
    on evm.logs ("substring"(data, 97, 32));

create index evm_logs_partitioned_idx_topic_two
    on evm.logs ((topics[2]));

create index evm_logs_partitioned_idx_topic_three
    on evm.logs ((topics[3]));

create index evm_logs_partitioned_idx_topic_four
    on evm.logs ((topics[4]));

create index evm_logs_partitioned_idx_created_at
    on evm.logs (created_at);

create index idx_evm_partitioned_logs_ordered_by_block_and_created_at
    on evm.logs (evm_chain_id, address, event_sig, block_number, created_at);

create index evm_logs_partitioned_idx_tx_hash
    on evm.logs (tx_hash);

create index evm_logs_partitioned_by_timestamp
    on evm.logs (evm_chain_id, address, event_sig, block_timestamp, block_number);

-- Create the partitions (we assume 41 partitions, prime number to reduce collisions)
CREATE TABLE evm.logs_partition0 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 0);
CREATE TABLE evm.logs_partition1 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 1);
CREATE TABLE evm.logs_partition2 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 2);
CREATE TABLE evm.logs_partition3 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 3);
CREATE TABLE evm.logs_partition4 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 4);
CREATE TABLE evm.logs_partition5 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 5);
CREATE TABLE evm.logs_partition6 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 6);
CREATE TABLE evm.logs_partition7 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 7);
CREATE TABLE evm.logs_partition8 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 8);
CREATE TABLE evm.logs_partition9 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 9);
CREATE TABLE evm.logs_partition10 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 10);
CREATE TABLE evm.logs_partition11 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 11);
CREATE TABLE evm.logs_partition12 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 12);
CREATE TABLE evm.logs_partition13 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 13);
CREATE TABLE evm.logs_partition14 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 14);
CREATE TABLE evm.logs_partition15 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 15);
CREATE TABLE evm.logs_partition16 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 16);
CREATE TABLE evm.logs_partition17 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 17);
CREATE TABLE evm.logs_partition18 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 18);
CREATE TABLE evm.logs_partition19 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 19);
CREATE TABLE evm.logs_partition20 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 20);
CREATE TABLE evm.logs_partition21 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 21);
CREATE TABLE evm.logs_partition22 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 22);
CREATE TABLE evm.logs_partition23 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 23);
CREATE TABLE evm.logs_partition24 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 24);
CREATE TABLE evm.logs_partition25 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 25);
CREATE TABLE evm.logs_partition26 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 26);
CREATE TABLE evm.logs_partition27 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 27);
CREATE TABLE evm.logs_partition28 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 28);
CREATE TABLE evm.logs_partition29 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 29);
CREATE TABLE evm.logs_partition30 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 30);
CREATE TABLE evm.logs_partition31 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 31);
CREATE TABLE evm.logs_partition32 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 32);
CREATE TABLE evm.logs_partition33 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 33);
CREATE TABLE evm.logs_partition34 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 34);
CREATE TABLE evm.logs_partition35 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 35);
CREATE TABLE evm.logs_partition36 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 36);
CREATE TABLE evm.logs_partition37 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 37);
CREATE TABLE evm.logs_partition38 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 38);
CREATE TABLE evm.logs_partition39 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 39);
CREATE TABLE evm.logs_partition40 PARTITION OF evm.logs FOR VALUES WITH (MODULUS 41, REMAINDER 40);

INSERT INTO evm.logs SELECT * FROM evm.logs_backup;

-- +goose Down

DROP TABLE evm.logs;
ALTER TABLE evm.logs_backup RENAME TO logs;
