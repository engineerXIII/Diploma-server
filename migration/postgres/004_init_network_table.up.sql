create table IF NOT EXISTS hv_network
(
    id              serial primary key,
    network_name    text                                                      not null,
    network_address cidr default '0.0.0.0/0',
    network_vlan    int,
    router_id       int
        constraint hv_router_hv_router_id_fk
            references hv_network (id)
            on update cascade
);