ALTER TABLE public.hv_node
  DROP CONSTRAINT hv_node_hv_auth_token_hv_type_fk;

ALTER TABLE public.hv_node
    DROP CONSTRAINT hv_auth_token_pk;