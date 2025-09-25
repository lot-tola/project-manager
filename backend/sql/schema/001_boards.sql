-- +goose Up


CREATE TABLE boards (
	id UUID PRIMARY KEY,
	board_title TEXT NOT NULL,
	created_at TIME NOT NULL,
	updated_at TIME NOT NULL
);

-- +goose Down
DROP TABLE boards;
