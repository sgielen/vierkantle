-- migrate:up

ALTER TABLE vierkantle.users
	ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT false;

-- migrate:down

ALTER TABLE vierkantle.users
	DROP COLUMN is_admin;
