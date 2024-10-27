--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Ubuntu 16.4-1.pgdg22.04+2)
-- Dumped by pg_dump version 17.0 (Ubuntu 17.0-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: basket_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.basket_items (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    basket_id bigint,
    product_id bigint,
    quantity bigint
);


ALTER TABLE public.basket_items OWNER TO postgres;

--
-- Name: basket_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.basket_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.basket_items_id_seq OWNER TO postgres;

--
-- Name: basket_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.basket_items_id_seq OWNED BY public.basket_items.id;


--
-- Name: baskets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.baskets (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint
);


ALTER TABLE public.baskets OWNER TO postgres;

--
-- Name: baskets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.baskets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.baskets_id_seq OWNER TO postgres;

--
-- Name: baskets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.baskets_id_seq OWNED BY public.baskets.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint,
    product_id bigint,
    quantity bigint,
    price bigint
);


ALTER TABLE public.order_items OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_items_id_seq OWNER TO postgres;

--
-- Name: order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_items_id_seq OWNED BY public.order_items.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    order_status character varying(100) NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    quantity bigint,
    description text,
    price bigint,
    image text,
    category text
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    email text,
    password text,
    role character varying(100) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: basket_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket_items ALTER COLUMN id SET DEFAULT nextval('public.basket_items_id_seq'::regclass);


--
-- Name: baskets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.baskets ALTER COLUMN id SET DEFAULT nextval('public.baskets_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: basket_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.basket_items (id, created_at, updated_at, deleted_at, basket_id, product_id, quantity) FROM stdin;
15	2024-10-28 00:07:54.707785+05	2024-10-28 00:07:54.707785+05	2024-10-28 00:11:08.2504+05	1	4	1
14	2024-10-28 00:07:53.31824+05	2024-10-28 00:07:53.31824+05	2024-10-28 00:13:31.422335+05	1	5	1
16	2024-10-28 00:18:37.443863+05	2024-10-28 00:18:37.443863+05	2024-10-28 00:19:44.401085+05	1	2	1
26	2024-10-28 00:51:29.578805+05	2024-10-28 00:51:51.02989+05	\N	3	9	2
5	2024-10-27 23:49:44.234005+05	2024-10-28 00:07:43.792093+05	2024-10-28 00:07:43.990037+05	1	9	1
1	2024-10-27 23:13:37.256531+05	2024-10-27 23:13:37.256531+05	2024-10-27 23:14:00.974687+05	1	2	1
19	2024-10-28 00:33:41.018295+05	2024-10-28 01:00:19.84715+05	2024-10-28 01:00:19.955004+05	2	6	1
18	2024-10-28 00:33:39.548028+05	2024-10-28 00:33:39.548028+05	2024-10-28 00:33:42.591268+05	2	2	1
27	2024-10-28 01:10:00.133397+05	2024-10-28 01:10:00.133397+05	2024-10-28 01:10:38.496682+05	2	13	1
24	2024-10-28 00:36:36.105884+05	2024-10-28 01:10:40.348228+05	2024-10-28 01:10:40.636273+05	2	3	1
12	2024-10-28 00:07:51.16268+05	2024-10-28 00:07:51.16268+05	2024-10-28 00:07:55.847737+05	1	1	1
20	2024-10-28 00:33:43.62158+05	2024-10-28 01:10:41.505166+05	\N	2	1	6
21	2024-10-28 00:33:45.126794+05	2024-10-28 00:34:11.970506+05	2024-10-28 00:34:12.077405+05	2	3	1
17	2024-10-28 00:19:45.045337+05	2024-10-28 00:35:22.109936+05	\N	1	2	3
10	2024-10-27 23:49:48.454854+05	2024-10-27 23:49:48.454854+05	2024-10-27 23:49:53.959314+05	1	12	1
9	2024-10-27 23:49:47.880732+05	2024-10-27 23:49:47.880732+05	2024-10-27 23:49:54.368882+05	1	10	1
8	2024-10-27 23:49:47.367181+05	2024-10-27 23:49:47.367181+05	2024-10-27 23:49:55.124914+05	1	11	1
7	2024-10-27 23:49:46.102956+05	2024-10-27 23:49:46.102956+05	2024-10-27 23:51:31.767512+05	1	7	1
13	2024-10-28 00:07:52.44863+05	2024-10-28 00:14:40.512692+05	2024-10-28 00:35:22.674573+05	1	6	1
2	2024-10-27 23:13:39.853408+05	2024-10-28 00:35:24.2263+05	\N	1	3	20
11	2024-10-27 23:51:32.664907+05	2024-10-27 23:59:21.335147+05	2024-10-27 23:59:21.480084+05	1	7	1
6	2024-10-27 23:49:45.592062+05	2024-10-27 23:49:45.592062+05	2024-10-27 23:59:22.234025+05	1	8	1
23	2024-10-28 00:36:01.908873+05	2024-10-28 00:36:02.654035+05	\N	1	6	5
4	2024-10-27 23:14:37.976199+05	2024-10-27 23:14:42.107468+05	2024-10-27 23:14:42.387395+05	1	6	1
22	2024-10-28 00:34:12.439783+05	2024-10-28 00:36:35.528121+05	2024-10-28 00:36:35.648208+05	2	3	1
3	2024-10-27 23:14:02.695239+05	2024-10-28 00:09:08.569995+05	2024-10-28 00:09:08.992955+05	1	2	1
28	2024-10-28 01:10:42.606183+05	2024-10-28 01:10:42.606183+05	2024-10-28 01:10:47.817597+05	2	11	1
29	2024-10-28 01:10:43.041109+05	2024-10-28 01:10:48.562436+05	\N	2	10	8
25	2024-10-28 00:49:37.404439+05	2024-10-28 00:50:23.077297+05	\N	1	12	6
30	2024-10-28 01:10:43.947449+05	2024-10-28 01:10:49.826012+05	\N	2	13	12
\.


--
-- Data for Name: baskets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.baskets (id, created_at, updated_at, deleted_at, user_id) FROM stdin;
1	2024-10-27 23:13:34.913876+05	2024-10-27 23:13:34.913876+05	\N	1
2	2024-10-27 23:26:25.331414+05	2024-10-27 23:26:25.331414+05	\N	3
3	2024-10-28 00:51:25.424478+05	2024-10-28 00:51:25.424478+05	\N	4
\.


--
-- Data for Name: order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_items (id, created_at, updated_at, deleted_at, order_id, product_id, quantity, price) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, created_at, updated_at, deleted_at, user_id, order_status) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, created_at, updated_at, deleted_at, name, quantity, description, price, image, category) FROM stdin;
1	2024-10-22 10:28:12.162831+05	2024-10-22 10:28:12.162831+05	\N	burger	100	burger desc	1200	https://grab-n-go.s3.amazonaws.com/images/1729574890788244098.jpg	\N
2	2024-10-22 10:36:28.418131+05	2024-10-22 10:36:28.418131+05	\N	lavash	100	lavash desc	1390	https://grab-n-go.s3.amazonaws.com/images/1729575388309811086.jpg	\N
3	2024-10-22 10:48:53.1308+05	2024-10-22 10:48:53.1308+05	\N	Beef Burger Set	100		1990	https://grab-n-go.s3.amazonaws.com/images/1729576133041695742.jpg	\N
4	2024-10-22 10:49:54.239714+05	2024-10-22 10:49:54.239714+05	\N	Chiken doner	120		1890	https://grab-n-go.s3.amazonaws.com/images/1729576193976800419.jpg	\N
5	2024-10-22 10:50:22.246037+05	2024-10-22 10:50:22.246037+05	\N	Burger Combo	150		2590	https://grab-n-go.s3.amazonaws.com/images/1729576222119651969.jpg	\N
6	2024-10-22 10:51:10.448589+05	2024-10-22 10:51:10.448589+05	\N	Doner na batone	110		1690	https://grab-n-go.s3.amazonaws.com/images/1729576270390545331.jpg	\N
7	2024-10-22 10:52:04.995998+05	2024-10-22 10:52:04.995998+05	\N	Hotdogs	160		1490	https://grab-n-go.s3.amazonaws.com/images/1729576324864814559.jpg	\N
8	2024-10-22 10:52:45.740413+05	2024-10-22 10:52:45.740413+05	\N	Cheesecake	30		2590	https://grab-n-go.s3.amazonaws.com/images/1729576365663045033.jpg	\N
9	2024-10-22 10:53:18.502464+05	2024-10-22 10:53:18.502464+05	\N	Breakfast	50		1800	https://grab-n-go.s3.amazonaws.com/images/1729576398318085834.jpg	\N
10	2024-10-22 10:53:47.92549+05	2024-10-22 10:53:47.92549+05	\N	Trifle	40		2000	https://grab-n-go.s3.amazonaws.com/images/1729576427769361450.jpg	\N
11	2024-10-22 10:55:13.581697+05	2024-10-22 10:55:13.581697+05	\N	Omelette	25		1100	https://grab-n-go.s3.amazonaws.com/images/1729576513424676076.jpeg	\N
12	2024-10-22 11:30:40.270288+05	2024-10-22 11:30:40.270288+05	\N	tea	1		350	https://grab-n-go.s3.amazonaws.com/images/1729578640194051289.webp	
13	2024-10-22 12:02:00.928399+05	2024-10-22 12:02:00.928399+05	\N	drinks	2		590	https://grab-n-go.s3.amazonaws.com/images/1729580520835782796.jfif	
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, created_at, updated_at, deleted_at, name, email, password, role) FROM stdin;
1	2024-10-22 02:43:22.412015+05	2024-10-22 02:43:22.412015+05	\N	admin	admin@gmail.com	$2a$10$wIt55Jul1LCOnfHTuQvDkOEXGacgu6dYCzBV0syYZXlBVi6pf5qTG	admin
2	2024-10-22 09:54:27.414755+05	2024-10-22 09:54:27.414755+05	\N	shadi	shadi@gmail.com	$2a$10$4rg28pDCkkPeC7KLmvFYdefx9.Yvdoj/okJebAlVpJpxK0XElZZw2	customer
3	2024-10-27 16:44:15.425182+05	2024-10-27 16:44:15.425182+05	\N	iliyas	iliyas@gmail.com	$2a$10$Ox.hZknPLLxFTad0tHG.v.3.pTV7cHmGyCld7cJBQg3JkI.SVpMX.	customer
4	2024-10-27 22:49:33.519451+05	2024-10-27 22:49:33.519451+05	\N	one	one@gmail.com	$2a$10$wSWoQAC4e7dwvfEM.Svhp.HkORr/X09FYTENGfma0LQidAXWjwkhO	customer
\.


--
-- Name: basket_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.basket_items_id_seq', 30, true);


--
-- Name: baskets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.baskets_id_seq', 3, true);


--
-- Name: order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_items_id_seq', 1, false);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 13, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 4, true);


--
-- Name: basket_items basket_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket_items
    ADD CONSTRAINT basket_items_pkey PRIMARY KEY (id);


--
-- Name: baskets baskets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.baskets
    ADD CONSTRAINT baskets_pkey PRIMARY KEY (id);


--
-- Name: order_items order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT order_items_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: products uni_products_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT uni_products_name UNIQUE (name);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_basket_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_basket_items_deleted_at ON public.basket_items USING btree (deleted_at);


--
-- Name: idx_baskets_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_baskets_deleted_at ON public.baskets USING btree (deleted_at);


--
-- Name: idx_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_items_deleted_at ON public.order_items USING btree (deleted_at);


--
-- Name: idx_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_orders_deleted_at ON public.orders USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: basket_items fk_basket_items_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket_items
    ADD CONSTRAINT fk_basket_items_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: basket_items fk_baskets_basket_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket_items
    ADD CONSTRAINT fk_baskets_basket_items FOREIGN KEY (basket_id) REFERENCES public.baskets(id);


--
-- Name: baskets fk_baskets_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.baskets
    ADD CONSTRAINT fk_baskets_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: order_items fk_order_items_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: order_items fk_order_items_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: orders fk_orders_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_orders_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

