-------- create t_logs --------
create table t_logs (
    id varchar (36) primary key,
    detail varchar (1000) not null,
    created_by varchar,
    created_at timestamp with time zone not null
);

create unique index idx_t_logs_id
    on t_logs (id);