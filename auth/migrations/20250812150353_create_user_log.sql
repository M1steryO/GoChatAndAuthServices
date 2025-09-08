-- +goose Up
create type user_action as ENUM ('create_account', 'auth', 'logout');
create table user_log
(
    id         serial primary key,
    user_id    int         not null,
    action     user_action not null,
    created_at timestamp   not null default now(),
    FOREIGN KEY (user_id) REFERENCES "user" (id) ON UPDATE CASCADE
);
-- +goose Down
-- +goose StatementBegin
drop table user_log;
drop type user_action;
-- +goose StatementEnd
