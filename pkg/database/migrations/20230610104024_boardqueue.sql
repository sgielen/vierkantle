-- migrate:up

CREATE TABLE vierkantle.board_queue (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	user_id BIGINT NOT NULL REFERENCES vierkantle.users (id),
	board_name TEXT NOT NULL,
	board BYTEA NOT NULL,
	added_at TIMESTAMPTZ NOT NULL,
	removed_at TIMESTAMPTZ NULL
);

-- migrate:down

DROP TABLE vierkantle.board_queue;
