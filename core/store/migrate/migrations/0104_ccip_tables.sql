-- +goose Up
-- +goose StatementBegin
CREATE TABLE ccip_requests
(
    PRIMARY KEY (source_chain_id, dest_chain_id, seq_num),
    source_chain_id text           NOT NULL,
    dest_chain_id   text           NOT NULL,
    seq_num         numeric(78, 0) NOT NULL,
    sender          bytea          NOT NULL,
    receiver        bytea          NOT NULL,
    data            bytea          NOT NULL,
    tokens          text[]         NOT NULL,
    amounts         text[]         NOT NULL,
    executor        bytea, -- optional, can be null
    options         bytea,
    raw             bytea,
    status          text,
    created_at      timestamptz    NOT NULL,
    updated_at      timestamptz    NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_ccip_requests_seq ON ccip_requests USING brin (seq_num);
CREATE INDEX IF NOT EXISTS idx_ccip_requests_source ON ccip_requests USING brin (source_chain_id);
CREATE INDEX IF NOT EXISTS idx_ccip_requests_dest ON ccip_requests USING brin (dest_chain_id);
CREATE INDEX IF NOT EXISTS idx_ccip_requests_status ON ccip_requests USING brin (status);
CREATE INDEX IF NOT EXISTS idx_ccip_requests_updated ON ccip_requests USING brin (updated_at);

CREATE TABLE ccip_relay_reports
(
    root        bytea PRIMARY KEY,
    min_seq_num numeric(78, 0) NOT NULL,
    max_seq_num numeric(78, 0) NOT NULL,
    created_at  timestamptz    NOT NULL DEFAULT now()
        CONSTRAINT chk_root CHECK (octet_length(root) = 32)
);

CREATE TABLE ccip_relay_specs
(
    id                                         SERIAL PRIMARY KEY,
    contract_id                                text,
    relay                                      text,
    relay_config                               JSONB,
    p2p_bootstrap_peers                        text[],
    is_bootstrap_peer                          boolean                  NOT NULL,
    ocr_key_bundle_id                          bytea,
    monitoring_endpoint                        text,
    transmitter_id                             text,
    blockchain_timeout                         bigint,
    contract_config_tracker_subscribe_interval bigint,
    contract_config_tracker_poll_interval      bigint,
    contract_config_confirmations              integer                  NOT NULL,
    juels_per_fee_coin_pipeline                text                     NOT NULL,
    created_at                                 timestamp with time zone NOT NULL,
    updated_at                                 timestamp with time zone NOT NULL,

    on_ramp_id                                 text                     NOT NULL,
    off_ramp_id                                text                     NOT NULL,
    source_evm_chain_id                        numeric(78, 0)           NOT NULL REFERENCES evm_chains (id),
    dest_evm_chain_id                          numeric(78, 0)           NOT NULL REFERENCES evm_chains (id)
);

CREATE TABLE ccip_execution_specs
(
    id                                         SERIAL PRIMARY KEY,
    contract_id                                text,
    relay                                      text,
    relay_config                               JSONB,
    p2p_bootstrap_peers                        text[],
    is_bootstrap_peer                          boolean                  NOT NULL,
    ocr_key_bundle_id                          bytea,
    monitoring_endpoint                        text,
    transmitter_id                             text,
    blockchain_timeout                         bigint,
    contract_config_tracker_subscribe_interval bigint,
    contract_config_tracker_poll_interval      bigint,
    contract_config_confirmations              integer                  NOT NULL,
    juels_per_fee_coin_pipeline                text                     NOT NULL,
    created_at                                 timestamp with time zone NOT NULL,
    updated_at                                 timestamp with time zone NOT NULL,

    on_ramp_id                                 text                     NOT NULL,
    off_ramp_id                                text                     NOT NULL,
    executor_id                                text                     NOT NULL,
    source_evm_chain_id                        numeric(78, 0)           NOT NULL REFERENCES evm_chains (id),
    dest_evm_chain_id                          numeric(78, 0)           NOT NULL REFERENCES evm_chains (id)
);

CREATE TABLE ccip_contract_configs
(
    contract_address        bytea                    NOT NULL primary key,
    config_digest           bytea                    NOT NULL,
    config_count            bigint                   NOT NULL,
    signers                 bytea[],
    transmitters            text[],
    f                       smallint                 NOT NULL,
    onchain_config          bytea,
    offchain_config_version bigint                   NOT NULL,
    offchain_config         bytea,
    created_at              timestamp with time zone NOT NULL,
    updated_at              timestamp with time zone NOT NULL,
    CONSTRAINT ccip_contract_configs_config_digest_check CHECK ((octet_length(config_digest) = 32))
);

CREATE TABLE ccip_pending_transmissions
(
    contract_address      bytea                    NOT NULL,
    config_digest         bytea                    NOT NULL,
    epoch                 bigint                   NOT NULL,
    round                 bigint                   NOT NULL,
    "time"                timestamp with time zone NOT NULL,
    extra_hash            bytea,
    report                bytea,
    attributed_signatures bytea[],
    created_at            timestamp with time zone NOT NULL,
    updated_at            timestamp with time zone NOT NULL,
    CONSTRAINT ccip_pending_transmissions_config_digest_check CHECK ((octet_length(config_digest) = 32))
);

ALTER TABLE ONLY ccip_pending_transmissions
    ADD CONSTRAINT ccip_pending_transmissions_pkey
        PRIMARY KEY (contract_address, config_digest, epoch, round);

-- TODO: might be able to keep these in a generic larger genocr states table?
CREATE TABLE ccip_persistent_states
(
    contract_address       bytea                    NOT NULL,
    config_digest          bytea                    NOT NULL,
    epoch                  bigint                   NOT NULL,
    highest_sent_epoch     bigint                   NOT NULL,
    highest_received_epoch bigint[]                 NOT NULL,
    created_at             timestamp with time zone NOT NULL,
    updated_at             timestamp with time zone NOT NULL,
    CONSTRAINT ccip_persistent_states_config_digest_check
        CHECK ((octet_length(config_digest) = 32))
);

ALTER TABLE ONLY ccip_persistent_states
    ADD CONSTRAINT ccip_persistent_states_pkey
        PRIMARY KEY (contract_address, config_digest);

ALTER TABLE jobs
    ADD COLUMN ccip_relay_spec_id     INT REFERENCES ccip_relay_specs (id),
    ADD COLUMN ccip_execution_spec_id INT REFERENCES ccip_execution_specs (id),
    DROP CONSTRAINT chk_only_one_spec,
    ADD CONSTRAINT chk_only_one_spec CHECK (
            num_nonnulls(
                    ocr_oracle_spec_id,
                    ocr2_oracle_spec_id,
                    direct_request_spec_id,
                    flux_monitor_spec_id,
                    keeper_spec_id,
                    cron_spec_id,
                    webhook_spec_id,
                    vrf_spec_id,
                    blockhash_store_spec_id,
                    bootstrap_spec_id,
                    ccip_relay_spec_id,
                    ccip_execution_spec_id) = 1
        );

CREATE INDEX IF NOT EXISTS idx_ccip_relay_min_seq ON ccip_relay_reports USING brin (min_seq_num);
CREATE INDEX IF NOT EXISTS idx_ccip_relay_max_seq ON ccip_relay_reports USING brin (max_seq_num);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ccip_requests;
DROP TABLE ccip_relay_reports;

ALTER TABLE jobs
    DROP CONSTRAINT chk_only_one_spec,
    ADD CONSTRAINT chk_only_one_spec CHECK (
            num_nonnulls(
                    ocr_oracle_spec_id,
                    ocr2_oracle_spec_id,
                    direct_request_spec_id,
                    flux_monitor_spec_id,
                    keeper_spec_id,
                    cron_spec_id,
                    webhook_spec_id,
                    vrf_spec_id,
                    blockhash_store_spec_id,
                    bootstrap_spec_id) = 1
        );
ALTER TABLE jobs
    DROP COLUMN ccip_relay_spec_id;
ALTER TABLE jobs
    DROP COLUMN ccip_execution_spec_id;
DROP TABLE IF EXISTS ccip_relay_specs;
DROP TABLE IF EXISTS ccip_execution_specs;
DROP TABLE IF EXISTS ccip_contract_configs;
DROP TABLE ccip_pending_transmissions;
DROP TABLE ccip_persistent_states;
-- +goose StatementEnd