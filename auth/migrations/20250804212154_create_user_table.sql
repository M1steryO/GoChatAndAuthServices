-- +goose Up
create type user_role AS ENUM ('USER', 'ADMIN');
create table "user"
(
    id         serial primary key,
    email      varchar(255) unique not null,
    name       varchar(255)        not null,
    role       user_role           not null default 'USER',
    password   varchar(255)        not null,
    created_at timestamp           not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table "user";
