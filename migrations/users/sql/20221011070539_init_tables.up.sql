create table t_users (
    id varchar (36) primary key,
    email varchar (255) unique not null,
    password_hash varchar (255) not null,
    dealer_code varchar(50) unique not null,
    role_name varchar(100) not null,
    created_by varchar,
    created_at timestamp with time zone not null,
    updated_by varchar,
    updated_at timestamp with time zone not null,
    deleted_by varchar,
    deleted_at timestamp with time zone
);

create unique index idx_user_id
    on t_users (id);
CREATE UNIQUE INDEX idx_user_email
    ON t_users (email);

