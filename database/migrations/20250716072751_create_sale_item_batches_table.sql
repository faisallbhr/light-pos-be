-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sale_item_batches (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        sale_id BIGINT UNSIGNED NOT NULL,
        purchase_item_id BIGINT UNSIGNED NOT NULL,
        quantity INT NOT NULL,
        FOREIGN key (sale_id) REFERENCES sales (id) ON DELETE CASCADE,
        FOREIGN key (purchase_item_id) REFERENCES purchase_items (id) ON DELETE CASCADE
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS sale_item_batches;
-- +goose StatementEnd