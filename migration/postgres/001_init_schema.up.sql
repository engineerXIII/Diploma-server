CREATE TABLE IF NOT EXISTS public.hv_auth_token
(
    id SERIAL,
    hv_type text COLLATE pg_catalog."default" NOT NULL,
    token_name text COLLATE pg_catalog."default",
    CONSTRAINT hv_auth1_token_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.hv_node
(
    id SERIAL,
    node_address text COLLATE pg_catalog."default" NOT NULL,
    secret_name text COLLATE pg_catalog."default",
    secret_value text COLLATE pg_catalog."default",
    node_name text COLLATE pg_catalog."default",
    CONSTRAINT hv_node_pkey PRIMARY KEY (id)
);