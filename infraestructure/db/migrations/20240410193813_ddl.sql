-- +goose Up
-- +goose StatementBegin
create table contact (
    id uuid primary key,
    name varchar(255) not null,
    lastname varchar(255) not null,
    email varchar(255) not null,
    phone varchar(255) not null,
    create_at timestamp not null default (now()),
    updated_at timestamp not null default (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
