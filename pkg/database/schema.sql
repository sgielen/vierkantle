CREATE SCHEMA IF NOT EXISTS vierkantle;

CREATE TABLE vierkantle.scores (
	board_name TEXT NOT NULL,
	anonymous_id BIGINT NOT NULL,
	PRIMARY KEY (board_name, anonymous_id),
	team_size INT NOT NULL,
	words INT NOT NULL,
	seconds INT NOT NULL,
	started_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);
