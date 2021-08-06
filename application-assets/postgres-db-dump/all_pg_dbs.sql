--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'md53175bce1d3201d16594cebf9d7eb3f9d';






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

-- Dumped from database version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)

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
-- PostgreSQL database dump complete
--

--
-- Database "go_movies" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)

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
-- Name: go_movies; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE go_movies WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_IN';


ALTER DATABASE go_movies OWNER TO postgres;

\connect go_movies

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.genres (
    id integer NOT NULL,
    genre_name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.genres OWNER TO postgres;

--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.genres_id_seq OWNER TO postgres;

--
-- Name: genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.genres_id_seq OWNED BY public.genres.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying,
    description text,
    year integer,
    release_date date,
    runtime integer,
    rating integer,
    cbfc_rating character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    poster character varying
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- Name: movies_genres; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies_genres (
    id integer NOT NULL,
    movie_id integer,
    genre_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.movies_genres OWNER TO postgres;

--
-- Name: movies_genres_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_genres_id_seq OWNER TO postgres;

--
-- Name: movies_genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_genres_id_seq OWNED BY public.movies_genres.id;


--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: genres id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.genres ALTER COLUMN id SET DEFAULT nextval('public.genres_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: movies_genres id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies_genres ALTER COLUMN id SET DEFAULT nextval('public.movies_genres_id_seq'::regclass);


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.genres (id, genre_name, created_at, updated_at) FROM stdin;
1	Drama	2021-05-17 00:00:00	2021-05-17 00:00:00
2	Crime	2021-05-17 00:00:00	2021-05-17 00:00:00
3	Action	2021-05-17 00:00:00	2021-05-17 00:00:00
4	Comic Book	2021-05-17 00:00:00	2021-05-17 00:00:00
5	Sci-Fi	2021-05-17 00:00:00	2021-05-17 00:00:00
6	Mystery	2021-05-17 00:00:00	2021-05-17 00:00:00
7	Adventure	2021-05-17 00:00:00	2021-05-17 00:00:00
8	Comedy	2021-05-17 00:00:00	2021-05-17 00:00:00
9	Romance	2021-05-17 00:00:00	2021-05-17 00:00:00
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, title, description, year, release_date, runtime, rating, cbfc_rating, created_at, updated_at, poster) FROM stdin;
5	Theory of EveryThing	The Theory of Everything is a 2014 biographical romantic drama film directed by James Marsh.	2014	2014-03-16	222	5	U/A	2021-08-02 20:28:09.727854	2021-08-06 00:36:56.194886	/kJuL37NTE51zVP3eG5aGMyKAIlh.jpg
2	The Godfather	The aging patriarch of an organized crime dynasty transfers control to his son	1972	1972-03-24	175	5	U/A	2021-05-17 00:00:00	2021-08-06 00:37:53.910835	/3bhkrj58Vtu7enYsRolD1fZdja1.jpg
4	American Psycho	 A wealthy New York investment banking executive hides his alternate psychopathic ego.	2000	2000-04-14	102	5	A	2021-05-17 00:00:00	2021-08-06 20:50:33.929311	/3ddHhfMlZHZCefHDeaP8FzSoH4Y.jpg
1	The Shawshank Redemption	Two imprisoned men bond over a number of years.	1994	1994-10-14	142	5	U	2021-05-17 00:00:00	2021-08-06 21:22:41.459558	/q6y0Go1tsGEsmtFryDOJo3dEmqu.jpg
3	The Dark Knight	The menace known as the Joker wreaks havoc on Gotham City.	2008	2008-07-18	152	5	A	2021-05-17 00:00:00	2021-08-06 21:45:07.596417	/qJ2tW6WMUDux911r6m7haRef0WH.jpg
\.


--
-- Data for Name: movies_genres; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies_genres (id, movie_id, genre_id, created_at, updated_at) FROM stdin;
1	1	1	2021-08-01 00:00:00	2021-08-01 00:00:00
2	1	2	2021-08-01 00:00:00	2021-08-01 00:00:00
\.


--
-- Name: genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.genres_id_seq', 9, true);


--
-- Name: movies_genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_genres_id_seq', 1, false);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 15, true);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);


--
-- Name: movies_genres movies_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: movies_genres fk_movie_genries_genre_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT fk_movie_genries_genre_id FOREIGN KEY (genre_id) REFERENCES public.genres(id);


--
-- Name: movies_genres fk_movie_genries_movie_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT fk_movie_genries_movie_id FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

\connect postgres

--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.3 (Ubuntu 13.3-1.pgdg20.04+1)

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
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

