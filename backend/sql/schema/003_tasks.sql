-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    list_id UUID NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
    task_title TEXT NOT NULL,
    description TEXT NOT NULL,
    due_date TIMESTAMP NOT NULL,
    assigned_to TEXT NOT NULL,
		status TEXT NOT NULL,
    labels TEXT[] NOT NULL,
		created_at TIME NOT NULL,
		updated_at TIME NOT NULL
);
-- +goose Down
DROP TABLE tasks;
