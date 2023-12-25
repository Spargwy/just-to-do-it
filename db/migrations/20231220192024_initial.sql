-- migrate:up
set timezone = "Asia/Bishkek";
create extension if not exists "uuid-ossp" with schema public;
comment on extension "uuid-ossp" is 'generate universally unique identifiers (UUIDs)';

-- migrate:down
set timezone = 'UTC';
drop extension if exists "uuid-ossp";

