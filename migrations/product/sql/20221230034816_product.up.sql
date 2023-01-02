-------- create t_products --------
create table t_products (
    id varchar (36) primary key,
    model varchar(100) not null,
    offset_name varchar(100),
    name varchar(2000) not null,
    mat_size varchar(200),
    color varchar(100),
    brand varchar(100) not null,
    code varchar(30) not null,
    prod_grp_code varchar(100),
    pitch_circle_code varchar(100) not null,
    tread_ware varchar(100) not null,
    width_name varchar(100) not null,
    link_id varchar(100) unique not null,
    price numeric(18, 5) default 0,
    dealer_price numeric(18, 5) default 0,
    quantity_balance numeric(21, 5) default 0,
    quantity_balance_b numeric(21, 5) default 0,
    created_by varchar,
    created_at timestamp with time zone not null,
    updated_by varchar,
    updated_at timestamp with time zone not null,
    deleted_by varchar,
    deleted_at timestamp with time zone
);

create unique index idx_t_products_id
    on t_products (id);
