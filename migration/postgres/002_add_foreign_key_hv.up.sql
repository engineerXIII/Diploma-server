ALTER TABLE public.hv_node
    ADD hv_type TEXT;

ALTER TABLE public.hv_auth_token
    ADD CONSTRAINT hv_auth_token_pk UNIQUE (hv_type);

ALTER TABLE public.hv_node
    ADD CONSTRAINT hv_node_hv_auth_token_hv_type_fk
        FOREIGN KEY (hv_type) REFERENCES public.hv_auth_token (hv_type);
