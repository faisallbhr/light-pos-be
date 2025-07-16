-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    products (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        sku VARCHAR(100) UNIQUE NOT NULL,
        name VARCHAR(100) NOT NULL,
        image VARCHAR(255),
        buy_price INT NOT NULL,
        sell_price INT NOT NULL,
        stock INT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS products;
-- +goose StatementEnd