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

create unique index idx_user_id
    on t_users (id);
CREATE UNIQUE INDEX idx_user_email
    ON t_users (email);

INSERT INTO t_users (created_at,updated_at,created_by,updated_by,email,dealer_code,password_hash,role_name,verification_code,verified,id)
VALUES ('2022-12-28 09:21:42.395','2022-12-28 09:21:42.395','u_system','u_system','admin@ppsw.com','123456','$2a$10$kpzqCBFQjFjPQHBA43h/y.7p6ZBd7pyVF8IxEGUjudc9W5oYwyKJ6','admin','',false,'e2bb7391-9a6f-4b36-ad54-11df22c2355d')
