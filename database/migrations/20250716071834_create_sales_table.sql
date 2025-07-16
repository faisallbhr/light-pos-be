-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sales (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        invoice_number VARCHAR(100) UNIQUE NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        sale_date DATE NOT NULL,
        total_price INT NOT NULL,
        payment_type VARCHAR(50) NOT NULL,
        paid_amount INT NOT NULL,
        `change` INT NOT NULL,
        FOREIGN key (user_id) REFERENCES users (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS sales;
-- +goose StatementEnd