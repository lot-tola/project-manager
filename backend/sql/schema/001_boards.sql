-- +goose Up


CREATE TABLE boards (
	id UUID PRIMARY KEY,
	board_title TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE boards;
