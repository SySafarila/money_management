--
-- PostgreSQL database cluster dump
--

-- Started on 2026-09-12 21:55:50

\restrict pglzViiEhcHrkhRFLbZQu3nLMKHlLaj9iSCbp22sNxJpvnDhUu9DJxCbl03DYrf

SET default_transaction_read_only = off;

SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;

--
-- User Configurations
--








\unrestrict pglzViiEhcHrkhRFLbZQu3nLMKHlLaj9iSCbp22sNxJpvnDhUu9DJxCbl03DYrf

--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

\restrict sCw9Kvgvi2ztWKoEVMjJQKkkiW7jLHYuJLz74r4uj3wFms5tW6teotQRLWKP0ki

-- Dumped from database version 18.4
-- Dumped by pg_dump version 18.4

-- Started on 2026-09-12 21:55:50

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- Completed on 2026-09-12 21:55:51

--
-- PostgreSQL database dump complete
--

\unrestrict sCw9Kvgvi2ztWKoEVMjJQKkkiW7jLHYuJLz74r4uj3wFms5tW6teotQRLWKP0ki

--
-- Database "money_management" dump
--

--
-- PostgreSQL database dump
--

\restrict h5Hff0PWNWOjspgm0TZ4BqymI3HTzdlYJYT3GiBrkCP7ZBbzUpMVt5YFS95UPvZ

-- Dumped from database version 18.4
-- Dumped by pg_dump version 18.4

-- Started on 2026-09-12 21:55:51

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4931 (class 1262 OID 16384)
-- Name: money_management; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE money_management WITH TEMPLATE = template0 ENCODING = 'WIN1252' LOCALE_PROVIDER = libc LOCALE = 'en_US';


ALTER DATABASE money_management OWNER TO postgres;

\unrestrict h5Hff0PWNWOjspgm0TZ4BqymI3HTzdlYJYT3GiBrkCP7ZBbzUpMVt5YFS95UPvZ
\connect money_management
\restrict h5Hff0PWNWOjspgm0TZ4BqymI3HTzdlYJYT3GiBrkCP7ZBbzUpMVt5YFS95UPvZ

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'WIN1252';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 221 (class 1259 OID 16460)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying CONSTRAINT categories_namme_not_null NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16424)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL,
    amount bigint NOT NULL,
    is_income boolean DEFAULT true NOT NULL,
    description character varying NOT NULL,
    date timestamp with time zone DEFAULT now() NOT NULL,
    category_id uuid
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16388)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    username character varying NOT NULL,
    full_name character varying NOT NULL,
    password character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    verified_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4778 (class 2606 OID 16440)
-- Name: transactions transactions_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_unique UNIQUE (id);


--
-- TOC entry 4774 (class 2606 OID 16403)
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 4776 (class 2606 OID 16405)
-- Name: users users_username_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_unique UNIQUE (username);


-- Completed on 2026-09-12 21:55:52

--
-- PostgreSQL database dump complete
--

\unrestrict h5Hff0PWNWOjspgm0TZ4BqymI3HTzdlYJYT3GiBrkCP7ZBbzUpMVt5YFS95UPvZ

-- Completed on 2026-09-12 21:55:52

--
-- PostgreSQL database cluster dump complete
--

