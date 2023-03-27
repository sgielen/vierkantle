CREATE SCHEMA IF NOT EXISTS vierkantle;

CREATE TABLE vierkantle.users (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	username TEXT UNIQUE NOT NULL,
	email TEXT UNIQUE NULL,
	last_login_at TIMESTAMPTZ NULL
);

CREATE TABLE vierkantle.scores (
	board_name TEXT NOT NULL,
	anonymous_id BIGINT NOT NULL,
	PRIMARY KEY (board_name, anonymous_id),
	user_id BIGINT NULL REFERENCES vierkantle.users (id),
	team_size INT NOT NULL,
	words INT NOT NULL,
	seconds INT NOT NULL,
	started_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);
