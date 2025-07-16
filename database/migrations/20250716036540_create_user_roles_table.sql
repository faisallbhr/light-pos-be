-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    user_roles (
        user_id BIGINT UNSIGNED NOT NULL,
        role_id BIGINT UNSIGNED NOT NULL,
        PRIMARY key (user_id, role_id),
        FOREIGN key (user_id) REFERENCES users (id),
        FOREIGN key (role_id) REFERENCES roles (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS user_roles;
-- +goose StatementEnd