-------- create t_users --------
create table t_users (
    id varchar (36) primary key,
    email varchar (255) unique not null,
    password_hash varchar (255) not null,
    dealer_code varchar(50) unique not null,
    role_name varchar(100) not null,
    verified boolean not null default false,
    verification_code varchar (255),
    created_by varchar,
    created_at timestamp with time zone not null,
    updated_by varchar,
    updated_at timestamp with time zone not null,
    deleted_by varchar,
    deleted_at timestamp with time zone
);

create unique index idx_t_users_id
    on t_users (id);
create unique index idx_t_users_email
    on t_users (email);

insert into t_users (created_at, updated_at, created_by, updated_by, email, dealer_code, password_hash, role_name, verification_code, verified, id)
values ('2022-12-28 09:21:42.395','2022-12-28 09:21:42.395','u_system','u_system','admin@ppsw.com','999999999','$2a$10$kpzqCBFQjFjPQHBA43h/y.7p6ZBd7pyVF8IxEGUjudc9W5oYwyKJ6','admin','',true,'e2bb7391-9a6f-4b36-ad54-11df22c2355d');

-------- create t_dealers --------
create table t_dealers (
    id varchar (36) primary key,
    name varchar (255) not null,
    address varchar (1000) not null,
    email varchar (255),
    phone varchar (100),
    link_id varchar (100),
    dealer_code varchar(50) unique not null,
    set_default boolean default false,
    disabled boolean default false,
    created_by varchar,
    created_at timestamp with time zone not null,
    updated_by varchar,
    updated_at timestamp with time zone not null,
    deleted_by varchar,
    deleted_at timestamp with time zone
);

create unique index idx_t_dealer_id
    on t_dealers (id);