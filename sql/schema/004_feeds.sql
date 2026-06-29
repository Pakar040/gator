-- +goose Up
ALTER TABLE feeds
ADD COlUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds
DROP COlUMN last_fetched_at;