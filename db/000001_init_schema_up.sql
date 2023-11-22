-- Set timezone
set timezone = "Africa/Lagos";

-- Table: todo
create table if not exists todo (
    todo_id serial not null,
    title varchar(225) not null,
    description text not null,
    completed bool not null,
    start_at timestamptz not null,
    end_at timestamptz not null,
    PRIMARY KEY(todo_id)
);

alter table todo alter column completed set default false;

-- Add indexes
create index active_todo on todo (title) where completed = true;
