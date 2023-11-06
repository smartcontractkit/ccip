-- +goose Up
CREATE INDEX logs_idx_data_word_four ON evm.logs (substring(data from 97 for 32));
CREATE INDEX logs_idx_data_word_five ON evm.logs (substring(data from 129 for 32));


-- +goose Down
DROP INDEX IF EXISTS evm.logs_idx_data_word_five;
DROP INDEX IF EXISTS evm.logs_idx_data_word_four;