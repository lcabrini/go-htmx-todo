-- +goose Up

alter table tasks add column completed boolean not null default false;

-- +goose Down

alter table tasks drop column completed;