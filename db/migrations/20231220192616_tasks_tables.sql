-- migrate:up
create type task_priority as enum (
    'LOW',
    'MEDIUM',
    'HIGH',
    'EXTRA'
);

create type task_status as enum (
    'CREATED',
    'IN_PROGRESS',
    'TESTING',
    'DONE'
);

create table workspaces (
    id uuid not null default uuid_generate_v4()
        primary key,
    title varchar(80) not null
);

create table task_groups (
    id uuid not null default uuid_generate_v4()
        primary key,
    title varchar(80) not null
);

create table tasks (
    id uuid not null default uuid_generate_v4()
            primary key,
    workspace_id uuid references workspaces(id),
    parent_task_id uuid references tasks(id),
    creater_id uuid not null references users(id),
    responsible_user_id uuid references users(id),
    title varchar(80) not null,
    description text,
    status task_status not null default 'CREATED',
    task_group_id uuid references task_groups (id),
    priority task_priority,
    estimate_time integer,
    time_spent integer,
    deleted_at timestamptz,
    archived boolean
);
-- migrate:down
drop table tasks;
drop table workspaces;
drop table task_groups;
drop type task_status;
drop type task_priority;
