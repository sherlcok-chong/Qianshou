-- name: GetAllUser :many
select *
from "user";

-- name: CreateUsers :one
insert into "user" (name, user_type)
values ($1, $2) returning * ;
