-- +goose Up
create table "chat"
(
    id         serial primary key,
    created_at timestamp not null default now(),
    updated_at timestamp
);

create table "chat_member"
(
    id         serial primary key,
    username   varchar(255) not null,
    chat_id    int          not null,
    created_at timestamp    not null default now(),
    foreign key (chat_id)
        references "chat" (id)
        on delete cascade
        on update cascade
);

-- +goose Down
drop table chat;
drop table chat_members
