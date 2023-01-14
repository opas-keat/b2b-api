-------- create t_products --------
create table t_products (
    id varchar (36) primary key,
    code varchar not null,
    product_name varchar not null,
    mat_size varchar,
    color varchar,
    brand varchar not null,
    model varchar not null,
    width varchar,
    offset_name varchar,
    tread_ware varchar,
    pitch_circle_code varchar,
    group_code varchar not null,
    dot varchar,
    speed_index varchar,
    load_index varchar,
    pattern varchar,
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

-------- create t_product_files --------
create table t_product_files (
     id varchar (36) primary key,
     product_id varchar (36),
     file_id varchar (36)
);

create unique index idx_t_product_files_id
    on t_product_files (id);
