-- +goose Up
CREATE TABLE test_table_3 (
                              key TEXT NOT NULL,
                              val JSONB NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              PRIMARY KEY ( key)
);

-- +goose Down
drop table test_table_3;

