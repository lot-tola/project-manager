-- name: CreateBoard :one
INSERT INTO boards (id, created_at, updated_at, board_title)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetBoards :many
SELECT * FROM boards;

-- name: GetOneBoard :one
SELECT * FROM boards WHERE id=$1;

-- name: UpdateBoard :one
UPDATE boards SET board_title = $1, updated_at = $2 WHERE id = $3 RETURNING *;

-- name: DeleteBoard :exec
DELETE FROM boards WHERE id = $1;
