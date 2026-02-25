-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    long_url TEXT NOT NULL,
    -- This column automatically generates a binary hash of the long_url for O(1) lookups
    url_hash BINARY(16) AS (UNHEX(MD5(long_url))) STORED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Index for the redirect (GET /:code)
    INDEX idx_short_code (short_code),
    -- Unique Index for deduplication (POST /shorten)
    UNIQUE INDEX idx_url_hash (url_hash)
) ENGINE=InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd