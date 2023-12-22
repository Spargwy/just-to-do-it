-- migrate:up
alter table tasks add created_at timestamptz not null default now();

-- migrate:down
alter table tasks drop created_at;
