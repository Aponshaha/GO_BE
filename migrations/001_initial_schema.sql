--
-- PostgreSQL database dump
--

\restrict kI5AyHEjxQh09rXXY9aNhsZakhA0DpZdyVY4tmciyPgEmvuauqk9hKhJbtbpAUd

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.11 (Homebrew)

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
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    description text
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: coupons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.coupons (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    code text NOT NULL,
    discount numeric NOT NULL,
    expiry_date text NOT NULL
);


ALTER TABLE public.coupons OWNER TO postgres;

--
-- Name: coupons_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.coupons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.coupons_id_seq OWNER TO postgres;

--
-- Name: coupons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.coupons_id_seq OWNED BY public.coupons.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    phone_number text,
    address text
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.customers_id_seq OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: invoices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoices (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint NOT NULL,
    invoice_date text NOT NULL,
    total_amount numeric NOT NULL
);


ALTER TABLE public.invoices OWNER TO postgres;

--
-- Name: invoices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoices_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.invoices_id_seq OWNER TO postgres;

--
-- Name: invoices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invoices_id_seq OWNED BY public.invoices.id;


--
-- Name: order_coupons; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_coupons (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint NOT NULL,
    coupon_id bigint NOT NULL
);


ALTER TABLE public.order_coupons OWNER TO postgres;

--
-- Name: order_coupons_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_coupons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_coupons_id_seq OWNER TO postgres;

--
-- Name: order_coupons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_coupons_id_seq OWNED BY public.order_coupons.id;


--
-- Name: order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_items (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id bigint NOT NULL,
    product_id bigint NOT NULL,
    quantity bigint NOT NULL,
    unit_price numeric NOT NULL
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
    order_number text NOT NULL,
    customer_id bigint NOT NULL,
    total_amount numeric NOT NULL,
    status text NOT NULL,
    payment_id bigint,
    shipping_id bigint
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
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    method text NOT NULL,
    amount numeric NOT NULL
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: product_images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_images (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    url text NOT NULL,
    product_id bigint NOT NULL
);


ALTER TABLE public.product_images OWNER TO postgres;

--
-- Name: product_images_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_images_id_seq OWNER TO postgres;

--
-- Name: product_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_images_id_seq OWNED BY public.product_images.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    description text,
    price numeric NOT NULL,
    stock_quantity bigint NOT NULL,
    category_id bigint NOT NULL
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
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    product_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    rating bigint NOT NULL,
    comment text
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reviews_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO postgres;

--
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- Name: shipping_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shipping_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    address text NOT NULL,
    city text NOT NULL,
    postal_code text NOT NULL,
    country text NOT NULL
);


ALTER TABLE public.shipping_details OWNER TO postgres;

--
-- Name: shipping_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shipping_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shipping_details_id_seq OWNER TO postgres;

--
-- Name: shipping_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shipping_details_id_seq OWNED BY public.shipping_details.id;


--
-- Name: wishlist_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wishlist_items (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    wishlist_id bigint NOT NULL,
    product_id bigint NOT NULL
);


ALTER TABLE public.wishlist_items OWNER TO postgres;

--
-- Name: wishlist_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wishlist_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wishlist_items_id_seq OWNER TO postgres;

--
-- Name: wishlist_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wishlist_items_id_seq OWNED BY public.wishlist_items.id;


--
-- Name: wishlists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wishlists (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    customer_id bigint NOT NULL
);


ALTER TABLE public.wishlists OWNER TO postgres;

--
-- Name: wishlists_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wishlists_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wishlists_id_seq OWNER TO postgres;

--
-- Name: wishlists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wishlists_id_seq OWNED BY public.wishlists.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: coupons id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.coupons ALTER COLUMN id SET DEFAULT nextval('public.coupons_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: invoices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices ALTER COLUMN id SET DEFAULT nextval('public.invoices_id_seq'::regclass);


--
-- Name: order_coupons id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_coupons ALTER COLUMN id SET DEFAULT nextval('public.order_coupons_id_seq'::regclass);


--
-- Name: order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items ALTER COLUMN id SET DEFAULT nextval('public.order_items_id_seq'::regclass);


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: product_images id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images ALTER COLUMN id SET DEFAULT nextval('public.product_images_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- Name: shipping_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shipping_details ALTER COLUMN id SET DEFAULT nextval('public.shipping_details_id_seq'::regclass);


--
-- Name: wishlist_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlist_items ALTER COLUMN id SET DEFAULT nextval('public.wishlist_items_id_seq'::regclass);


--
-- Name: wishlists id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists ALTER COLUMN id SET DEFAULT nextval('public.wishlists_id_seq'::regclass);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: coupons coupons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.coupons
    ADD CONSTRAINT coupons_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: invoices invoices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_pkey PRIMARY KEY (id);


--
-- Name: order_coupons order_coupons_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_coupons
    ADD CONSTRAINT order_coupons_pkey PRIMARY KEY (id);


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
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: product_images product_images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT product_images_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- Name: shipping_details shipping_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shipping_details
    ADD CONSTRAINT shipping_details_pkey PRIMARY KEY (id);


--
-- Name: coupons uni_coupons_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.coupons
    ADD CONSTRAINT uni_coupons_code UNIQUE (code);


--
-- Name: customers uni_customers_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT uni_customers_email UNIQUE (email);


--
-- Name: orders uni_orders_order_number; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT uni_orders_order_number UNIQUE (order_number);


--
-- Name: wishlist_items wishlist_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlist_items
    ADD CONSTRAINT wishlist_items_pkey PRIMARY KEY (id);


--
-- Name: wishlists wishlists_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT wishlists_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_coupons_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_coupons_deleted_at ON public.coupons USING btree (deleted_at);


--
-- Name: idx_customers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_customers_deleted_at ON public.customers USING btree (deleted_at);


--
-- Name: idx_invoices_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_invoices_deleted_at ON public.invoices USING btree (deleted_at);


--
-- Name: idx_order_coupons_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_coupons_deleted_at ON public.order_coupons USING btree (deleted_at);


--
-- Name: idx_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_items_deleted_at ON public.order_items USING btree (deleted_at);


--
-- Name: idx_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_orders_deleted_at ON public.orders USING btree (deleted_at);


--
-- Name: idx_payments_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_payments_deleted_at ON public.payments USING btree (deleted_at);


--
-- Name: idx_product_images_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_product_images_deleted_at ON public.product_images USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_reviews_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_reviews_deleted_at ON public.reviews USING btree (deleted_at);


--
-- Name: idx_shipping_details_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shipping_details_deleted_at ON public.shipping_details USING btree (deleted_at);


--
-- Name: idx_wishlist_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wishlist_items_deleted_at ON public.wishlist_items USING btree (deleted_at);


--
-- Name: idx_wishlists_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wishlists_deleted_at ON public.wishlists USING btree (deleted_at);


--
-- Name: products fk_categories_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_categories_products FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- Name: order_coupons fk_coupons_orders; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_coupons
    ADD CONSTRAINT fk_coupons_orders FOREIGN KEY (coupon_id) REFERENCES public.coupons(id);


--
-- Name: orders fk_customers_orders; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_customers_orders FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- Name: invoices fk_invoices_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT fk_invoices_order FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: order_coupons fk_order_coupons_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_coupons
    ADD CONSTRAINT fk_order_coupons_order FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: order_items fk_order_items_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: order_items fk_orders_order_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_items
    ADD CONSTRAINT fk_orders_order_items FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: orders fk_payments_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_payments_order FOREIGN KEY (payment_id) REFERENCES public.payments(id);


--
-- Name: product_images fk_products_images; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_images
    ADD CONSTRAINT fk_products_images FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: reviews fk_products_reviews; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_products_reviews FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: reviews fk_reviews_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT fk_reviews_customer FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- Name: orders fk_shipping_details_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_shipping_details_order FOREIGN KEY (shipping_id) REFERENCES public.shipping_details(id);


--
-- Name: wishlist_items fk_wishlist_items_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlist_items
    ADD CONSTRAINT fk_wishlist_items_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: wishlists fk_wishlists_customer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT fk_wishlists_customer FOREIGN KEY (customer_id) REFERENCES public.customers(id);


--
-- Name: wishlist_items fk_wishlists_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlist_items
    ADD CONSTRAINT fk_wishlists_items FOREIGN KEY (wishlist_id) REFERENCES public.wishlists(id);


--
-- PostgreSQL database dump complete
--

\unrestrict kI5AyHEjxQh09rXXY9aNhsZakhA0DpZdyVY4tmciyPgEmvuauqk9hKhJbtbpAUd

