--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.6
-- Dumped by pg_dump version 9.6.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

ALTER TABLE IF EXISTS ONLY public.example DROP CONSTRAINT IF EXISTS example_pkey;
ALTER TABLE IF EXISTS public.example ALTER COLUMN pk DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.example_pk_seq;
DROP TABLE IF EXISTS public.example;
DROP SCHEMA IF EXISTS public;
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: example; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE example (
    pk integer NOT NULL,
    text_field text
);


ALTER TABLE example OWNER TO postgres;

--
-- Name: example_pk_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE example_pk_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE example_pk_seq OWNER TO postgres;

--
-- Name: example_pk_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE example_pk_seq OWNED BY example.pk;


--
-- Name: example pk; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY example ALTER COLUMN pk SET DEFAULT nextval('example_pk_seq'::regclass);


--
-- Data for Name: example; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO example (pk, text_field) VALUES (1, 'ex_1');
INSERT INTO example (pk, text_field) VALUES (2, 'ex_2');


--
-- Name: example_pk_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('example_pk_seq', 2, true);


--
-- Name: example example_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY example
    ADD CONSTRAINT example_pkey PRIMARY KEY (pk);


--
-- PostgreSQL database dump complete
--