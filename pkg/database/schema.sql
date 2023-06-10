SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: vierkantle; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA vierkantle;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: board_queue; Type: TABLE; Schema: vierkantle; Owner: -
--

CREATE TABLE vierkantle.board_queue (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    board_name text NOT NULL,
    board bytea NOT NULL,
    added_at timestamp with time zone NOT NULL,
    removed_at timestamp with time zone
);


--
-- Name: board_queue_id_seq; Type: SEQUENCE; Schema: vierkantle; Owner: -
--

CREATE SEQUENCE vierkantle.board_queue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: board_queue_id_seq; Type: SEQUENCE OWNED BY; Schema: vierkantle; Owner: -
--

ALTER SEQUENCE vierkantle.board_queue_id_seq OWNED BY vierkantle.board_queue.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: vierkantle; Owner: -
--

CREATE TABLE vierkantle.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: scores; Type: TABLE; Schema: vierkantle; Owner: -
--

CREATE TABLE vierkantle.scores (
    board_name text NOT NULL,
    anonymous_id bigint NOT NULL,
    team_size integer NOT NULL,
    words integer NOT NULL,
    seconds integer NOT NULL,
    started_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    user_id bigint
);


--
-- Name: users; Type: TABLE; Schema: vierkantle; Owner: -
--

CREATE TABLE vierkantle.users (
    id bigint NOT NULL,
    username text NOT NULL,
    email text,
    last_login_at timestamp with time zone,
    registered_at timestamp with time zone NOT NULL,
    is_admin boolean DEFAULT false NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: vierkantle; Owner: -
--

CREATE SEQUENCE vierkantle.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: vierkantle; Owner: -
--

ALTER SEQUENCE vierkantle.users_id_seq OWNED BY vierkantle.users.id;


--
-- Name: board_queue id; Type: DEFAULT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.board_queue ALTER COLUMN id SET DEFAULT nextval('vierkantle.board_queue_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.users ALTER COLUMN id SET DEFAULT nextval('vierkantle.users_id_seq'::regclass);


--
-- Name: board_queue board_queue_pkey; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.board_queue
    ADD CONSTRAINT board_queue_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: scores scores_pkey; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.scores
    ADD CONSTRAINT scores_pkey PRIMARY KEY (board_name, anonymous_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: lower_username_unique_idx; Type: INDEX; Schema: vierkantle; Owner: -
--

CREATE UNIQUE INDEX lower_username_unique_idx ON vierkantle.users USING btree (lower(username));


--
-- Name: board_queue board_queue_user_id_fkey; Type: FK CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.board_queue
    ADD CONSTRAINT board_queue_user_id_fkey FOREIGN KEY (user_id) REFERENCES vierkantle.users(id);


--
-- Name: scores scores_user_id_fkey; Type: FK CONSTRAINT; Schema: vierkantle; Owner: -
--

ALTER TABLE ONLY vierkantle.scores
    ADD CONSTRAINT scores_user_id_fkey FOREIGN KEY (user_id) REFERENCES vierkantle.users(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO vierkantle.schema_migrations (version) VALUES
    ('20230610094921'),
    ('20230610104024'),
    ('20230610105428');
