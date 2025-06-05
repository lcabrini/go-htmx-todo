-- +goose Up

create extension if not exists "uuid-ossp";

create type priorities as enum('low', 'medium', 'high');

create table tasks(
    id uuid primary key default uuid_generate_v4(),
    title text not null,
    description text not null,
    priority priorities not null default 'medium',
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

-- +goose Down

drop table tasks;
drop type priorities;
drop extension "uuid-ossp";