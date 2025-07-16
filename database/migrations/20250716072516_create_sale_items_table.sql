-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sale_items (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        sale_id BIGINT UNSIGNED NOT NULL,
        product_id BIGINT UNSIGNED NOT NULL,
        quantity INT NOT NULL,
        sell_price INT NOT NULL,
        total_price INT NOT NULL,
        FOREIGN key (sale_id) REFERENCES sales (id) ON DELETE CASCADE,
        FOREIGN key (product_id) REFERENCES products (id) ON DELETE CASCADE
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS sale_items;
-- +goose StatementEnd