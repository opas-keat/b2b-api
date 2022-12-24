create table t_users (
    id serial primary key,
    username varchar(50) not null,
    pwd varchar(255) not null,
    created_by           varchar,
    created_at           timestamp with time zone not null,
    updated_by           varchar,
    updated_at           timestamp with time zone not null,
    deleted_by           varchar,
    deleted_at           timestamp with time zone
);

create unique index users_username_uindex
    on t_users (username);

create table t_authentication_tokens (
    token_id serial primary key,
    user_id integer not null,
    auth_token varchar(255),
    generated_at timestamp with time zone not null,
    expires_at   timestamp with time zone not null
);
