-- +goose Up
CREATE TABLE lists (
	id UUID PRIMARY KEY,
	board_id UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
	list_title TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE lists;

