create table if not exists usr(
    id bigserial primary key ,
    created_at timestamp default now(),
    telegram_id bigint not null unique ,
    username varchar(256) not null
);