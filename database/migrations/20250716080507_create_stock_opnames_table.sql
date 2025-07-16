-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    stock_opnames (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        opname_date DATE NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        note TEXT,
        FOREIGN key (user_id) REFERENCES users (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS stock_opnames;
-- +goose StatementEnd