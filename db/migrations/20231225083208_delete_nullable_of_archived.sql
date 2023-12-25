-- migrate:up
update tasks set archived = false where archived is null;
alter table tasks alter column archived set not null;
alter table tasks alter column archived set default false;

-- migrate:down
alter table tasks alter column archived drop not null;
