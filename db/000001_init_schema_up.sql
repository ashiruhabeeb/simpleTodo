-- Set timezone
set timezone = "Africa/Lagos";

-- Table: user
create table if not exists users (
    user_id serial,
    username varchar(30) unique not null,
    fullname varchar(225) not null,
    email varchar(225) unique not null,
    password varchar(225) not null,
    address_id serial,
    house_number int not null,
    street_name varchar(225) not null,
    local_area varchar(100) not null,
    state varchar(225) not null,
    country varchar(225) not null,
    avatar varchar(225),
    dob date not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    primary key(user_id)
);

-- Table: todo
create table if not exists todo (
    todo_id serial,
    user_id int not null,
    title varchar(225) not null,
    description text not null,
    completed bool not null,
    start_at timestamptz not null,
    end_at timestamptz not null,
    PRIMARY KEY(todo_id)
);

-- alter statements
alter table todo alter column completed set default false;

alter table "todo" add foreign key ("user_id") references "users" ("user_id");

-- Add indexes
create index active_todo on todo (title) where completed = true;
