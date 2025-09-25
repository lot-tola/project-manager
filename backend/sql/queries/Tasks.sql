-- name: CreateTask :one
INSERT INTO tasks (id, list_id, task_title, description, due_date, assigned_to, status, labels, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7, 
		$8,
		$9,
		$10
)
RETURNING *;

