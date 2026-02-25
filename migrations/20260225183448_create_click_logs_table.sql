-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS click_logs (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    short_code VARCHAR(10) NOT NULL,
    ip_address VARCHAR(45) NOT NULL,
    user_agent TEXT,
    clicked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_analytics_short_code (short_code),
    INDEX idx_analytics_clicked_at (clicked_at)
) ENGINE=InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS click_logs;
-- +goose StatementEnd