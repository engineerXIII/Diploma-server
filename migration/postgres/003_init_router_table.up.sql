create table IF NOT EXISTS
    hv_router
(
    id              serial
        primary key,
    router_type     text not null,
    router_address  text not null
        unique,
    router_hostname text,
    router_port     int
);
