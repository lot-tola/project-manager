-- +goose Up
CREATE TABLE lists (
	id UUID PRIMARY KEY,
	board_id UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
	list_title TEXT NOT NULL,
	created_at TIME NOT NULL,
	updated_at TIME NOT NULL
);


-- +goose Down

DROP TABLE lists;

