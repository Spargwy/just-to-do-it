-- migrate:up
create table users(
    id uuid not null default uuid_generate_v4()
        primary key,
    name varchar(80),
    hashed_password varchar(72) not null
)

-- migrate:down
drop table users;
