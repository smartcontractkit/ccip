-- +goose Up
-- +goose StatementBegin
CREATE TABLE ccip_requests
(
    PRIMARY KEY (source_chain_id, dest_chain_id, on_ramp, off_ramp, seq_num),
    source_chain_id text           NOT NULL,
    dest_chain_id   text           NOT NULL,
    on_ramp         bytea          NOT NULL,
    off_ramp        bytea          NOT NULL,
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

CREATE INDEX IF NOT EXISTS idx_ccip_relay_min_seq ON ccip_relay_reports USING brin (min_seq_num);
CREATE INDEX IF NOT EXISTS idx_ccip_relay_max_seq ON ccip_relay_reports USING brin (max_seq_num);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ccip_requests;
DROP TABLE ccip_relay_reports;
-- +goose StatementEnd