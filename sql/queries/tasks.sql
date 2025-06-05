-- name: GetTask :one

select * from tasks
where id = $1 limit 1;

-- name: ListTasks :many

select * from tasks
order by created_at;

-- name: CreateTask :one

insert into tasks(title, description, priority)
values($1, $2, $3)
returning *;

-- name: UpdateTask :one

update tasks set
    title = $2,
    description = $3,
    priority = $4,
    updated_at = current_timestamp
where id = $1
returning *;

-- name: DeleteTask :exec

delete from tasks
where id = $1;