-- +goose Up
-- +goose StatementBegin

select 'up SQL query';

create table if not exists users
(
    id            uuid        not null
        default gen_random_uuid()
        constraint user_pk
            primary key,
    first_name    varchar(50) not null,
    second_name   varchar(50) not null,
    birthdate     date        not null,
    biography     text,
    city          text,
    password_hash text        not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

select 'down SQL query';

drop table if exists users;

-- +goose StatementEnd
