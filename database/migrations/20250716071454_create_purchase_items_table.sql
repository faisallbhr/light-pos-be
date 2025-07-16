-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    purchase_items (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        purchase_id BIGINT UNSIGNED NOT NULL,
        product_id BIGINT UNSIGNED NOT NULL,
        quantity INT NOT NULL,
        buy_price INT NOT NULL,
        total_price INT NOT NULL,
        remaining_quantity INT NOT NULL,
        FOREIGN key (purchase_id) REFERENCES purchases (id),
        FOREIGN key (product_id) REFERENCES products (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS purchase_items;
-- +goose StatementEnd