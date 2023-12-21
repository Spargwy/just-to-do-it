-- migrate:up
create table users(
    id uuid not null default uuid_generate_v4()
        primary key,
    email varchar(255) not null unique,
    name varchar(255),
    hashed_password varchar(72) not null
)

-- migrate:down
drop table users;
