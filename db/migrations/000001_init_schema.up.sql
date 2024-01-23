create table public.patients
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text
);

alter table public.patients
    owner to postgres;

create index idx_patients_deleted_at
    on public.patients (deleted_at);


create table public.orders
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    text       text,
    patient_id bigint
);

alter table public.orders
    owner to postgres;

create index idx_orders_deleted_at
    on public.orders (deleted_at);


create table public.order_logs
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id   bigint
        constraint fk_order_logs_order
            references public.orders,
    text       text
);

alter table public.order_logs
    owner to postgres;

create index idx_order_logs_deleted_at
    on public.order_logs (deleted_at);

