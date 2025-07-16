-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    stock_opname_items (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        stock_opname_id BIGINT UNSIGNED NOT NULL,
        product_id BIGINT UNSIGNED NOT NULL,
        system_stock INT NOT NULL,
        real_stock INT NOT NULL,
        difference INT NOT NULL,
        note TEXT,
        FOREIGN key (stock_opname_id) REFERENCES stock_opnames (id) ON DELETE CASCADE,
        FOREIGN key (product_id) REFERENCES products (id) ON DELETE CASCADE
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS stock_opname_items;
-- +goose StatementEnd