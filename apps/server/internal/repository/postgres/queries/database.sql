

-- name: CreateSource :one
INSERT into database_sources  (name, type , host , port , username, password_encrypted , enabled, created_by)
VALUES
($1, $2, $3 , $4 , $5 , $6 , $7, $8)
RETURNING name, type , host , port, username, enabled ;

-- name: GetSourceById :one
SELECT *
FROM database_sources where id = $1;

-- name: GetAllSources: many
SELECT * FROM database_sources;

-- name: ListAllSouces :many
SELECT * FROM database_sources
WHERE created_by = $1;


-- name: DeleteSource :exec
DELETE * from database_sources
WHERE id = $1;




