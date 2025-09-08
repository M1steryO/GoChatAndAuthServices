-- +goose Up
create table message
(
    id         serial primary key,
    chat_id    int                     not null,
    "from"     varchar(255)            not null,
    text       text                    not null,
    timestamp  timestamp               not null,
    created_at timestamp default now() not null,
    foreign key (chat_id) references chat (id) on delete cascade on update cascade
);

-- +goose Down
drop table message;
