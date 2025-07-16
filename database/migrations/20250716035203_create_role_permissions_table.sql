-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    role_permissions (
        role_id BIGINT UNSIGNED NOT NULL,
        permission_id BIGINT UNSIGNED NOT NULL,
        PRIMARY key (role_id, permission_id),
        FOREIGN key (role_id) REFERENCES roles (id),
        FOREIGN key (permission_id) REFERENCES permissions (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE if EXISTS role_permissions;
-- +goose StatementEnd