-- name: CreateTask :one
INSERT INTO tasks (id, list_id, task_title, description, due_date, assigned_to, status, labels)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7, 
		$8
)
RETURNING *;


-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=$1;

-- name: GetTask :one
SELECT * FROM tasks WHERE id=$1;

-- name: UpdateTask :exec
UPDATE tasks
SET task_title=$1, description=$2, due_date=$3, assigned_to=$4, status=$5, labels=$6
WHERE id=$7;

