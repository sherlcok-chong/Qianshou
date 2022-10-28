-- name: GetRelationType :one
select *
from relationship
where fid = $1
  and tid = $2;

-- name: UpdateRelationType :exec
update relationship
set relation_type = $1
where fid = $2
  and tid = $3;

-- name: GetAllRelations :many
select *
from relationship
where fid = $1;

-- name: CreateRelations :one

insert into relationship (fid, tid, relation_type)
values ($1, $2, $3) returning *;