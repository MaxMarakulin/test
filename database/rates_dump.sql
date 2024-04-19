--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

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
-- Name: rates; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA rates;


ALTER SCHEMA rates OWNER TO postgres;

--
-- Name: SCHEMA rates; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA rates IS 'База данных для курса валют';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: currencies; Type: TABLE; Schema: rates; Owner: postgres
--

CREATE TABLE rates.currencies (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE rates.currencies OWNER TO postgres;

--
-- Name: COLUMN currencies.name; Type: COMMENT; Schema: rates; Owner: postgres
--

COMMENT ON COLUMN rates.currencies.name IS 'Название валюты';


--
-- Name: currencies_id_seq; Type: SEQUENCE; Schema: rates; Owner: postgres
--

CREATE SEQUENCE rates.currencies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE rates.currencies_id_seq OWNER TO postgres;

--
-- Name: currencies_id_seq; Type: SEQUENCE OWNED BY; Schema: rates; Owner: postgres
--

ALTER SEQUENCE rates.currencies_id_seq OWNED BY rates.currencies.id;


--
-- Name: exchange_rate; Type: TABLE; Schema: rates; Owner: postgres
--

CREATE TABLE rates.exchange_rate (
    id integer NOT NULL,
    "from" integer,
    "to" integer,
    value double precision,
    source character varying
);


ALTER TABLE rates.exchange_rate OWNER TO postgres;

--
-- Name: COLUMN exchange_rate."from"; Type: COMMENT; Schema: rates; Owner: postgres
--

COMMENT ON COLUMN rates.exchange_rate."from" IS 'от какой валюты {currencies}';


--
-- Name: COLUMN exchange_rate."to"; Type: COMMENT; Schema: rates; Owner: postgres
--

COMMENT ON COLUMN rates.exchange_rate."to" IS 'к какой валюте {currencies}';


--
-- Name: COLUMN exchange_rate.value; Type: COMMENT; Schema: rates; Owner: postgres
--

COMMENT ON COLUMN rates.exchange_rate.value IS 'значение курса';


--
-- Name: COLUMN exchange_rate.source; Type: COMMENT; Schema: rates; Owner: postgres
--

COMMENT ON COLUMN rates.exchange_rate.source IS 'источник (формат url)';


--
-- Name: exchange_rate_id_seq; Type: SEQUENCE; Schema: rates; Owner: postgres
--

CREATE SEQUENCE rates.exchange_rate_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE rates.exchange_rate_id_seq OWNER TO postgres;

--
-- Name: exchange_rate_id_seq; Type: SEQUENCE OWNED BY; Schema: rates; Owner: postgres
--

ALTER SEQUENCE rates.exchange_rate_id_seq OWNED BY rates.exchange_rate.id;


--
-- Name: currencies id; Type: DEFAULT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.currencies ALTER COLUMN id SET DEFAULT nextval('rates.currencies_id_seq'::regclass);


--
-- Name: exchange_rate id; Type: DEFAULT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.exchange_rate ALTER COLUMN id SET DEFAULT nextval('rates.exchange_rate_id_seq'::regclass);


--
-- Data for Name: currencies; Type: TABLE DATA; Schema: rates; Owner: postgres
--

COPY rates.currencies (id, name) FROM stdin;
2	BYN
3	RUB
\.


--
-- Data for Name: exchange_rate; Type: TABLE DATA; Schema: rates; Owner: postgres
--

COPY rates.exchange_rate (id, "from", "to", value, source) FROM stdin;
11	3	2	28.6488	https://www.cbr-xml-daily.ru/daily_json.js
10	2	3	3.4772	https://api.nbrb.by/exrates/rates/456
\.


--
-- Name: currencies_id_seq; Type: SEQUENCE SET; Schema: rates; Owner: postgres
--

SELECT pg_catalog.setval('rates.currencies_id_seq', 3, true);


--
-- Name: exchange_rate_id_seq; Type: SEQUENCE SET; Schema: rates; Owner: postgres
--

SELECT pg_catalog.setval('rates.exchange_rate_id_seq', 11, true);


--
-- Name: currencies currencies_pk; Type: CONSTRAINT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.currencies
    ADD CONSTRAINT currencies_pk PRIMARY KEY (id);


--
-- Name: exchange_rate id; Type: CONSTRAINT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.exchange_rate
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- Name: exchange_rate exchange_rate_currencies_from_fk; Type: FK CONSTRAINT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.exchange_rate
    ADD CONSTRAINT exchange_rate_currencies_from_fk FOREIGN KEY ("from") REFERENCES rates.currencies(id);


--
-- Name: exchange_rate exchange_rate_currencies_to_fk; Type: FK CONSTRAINT; Schema: rates; Owner: postgres
--

ALTER TABLE ONLY rates.exchange_rate
    ADD CONSTRAINT exchange_rate_currencies_to_fk FOREIGN KEY ("to") REFERENCES rates.currencies(id);


--
-- PostgreSQL database dump complete
--

