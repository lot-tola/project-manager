-- name: GetLists :many
SELECT 
    l.id   AS list_id,
    l.list_title,
    t.id   AS task_id,
    t.task_title,
    t.description,
    t.due_date,
    t.assigned_to,
		t.status,
    t.labels
FROM lists l
LEFT JOIN tasks t ON t.list_id = l.id
WHERE l.board_id = $1;

-- name: GetAllLists :many
SELECT * FROM lists;

-- name: CreateList :one
INSERT INTO lists (id, board_id, list_title)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateList :one
UPDATE lists SET list_title = $1, updated_at = $2 WHERE id = $3 RETURNING *;

-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1;


