-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    product_categories (
        product_id BIGINT UNSIGNED NOT NULL,
        category_id BIGINT UNSIGNED NOT NULL,
        PRIMARY key (product_id, category_id),
        FOREIGN key (product_id) REFERENCES products (id) ON DELETE CASCADE,
        FOREIGN key (category_id) REFERENCES categories (id) ON DELETE CASCADE
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS product_categories;
-- +goose StatementEnd