create table t_users (
    uuid varchar (36) primary key,
    name varchar (255) not null,
    email varchar (255) unique not null,
    password_hash varchar (511) not null,
    created_by           varchar,
    created_at           timestamp with time zone not null,
    updated_by           varchar,
    updated_at           timestamp with time zone not null,
    deleted_by           varchar,
    deleted_at           timestamp with time zone
);

create unique index idx_user_uuid
    on t_users (uuid);
CREATE UNIQUE INDEX idx_user_email
    ON t_users (email);

