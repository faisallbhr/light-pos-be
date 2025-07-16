-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    purchases (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        invoice_number VARCHAR(100) UNIQUE,
        supplier_id BIGINT UNSIGNED,
        purchase_date TIMESTAMP NOT NULL,
        total_price INT NOT NULL,
        FOREIGN key (supplier_id) REFERENCES suppliers (id) ON DELETE CASCADE
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS purchases;
-- +goose StatementEnd