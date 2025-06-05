-- +goose Up

create type priorities as enum('low', 'medium', 'high');

create table tasks(
    id uuid primary key,
    title text not null,
    description text not null,
    priority priorities not null default 'medium'
);

-- +goose Down

drop table tasks;
drop type priorities;