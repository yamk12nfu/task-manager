create table if not exists tasks (
    id varchar(16) primary key,
    name text not null,
    description text not null,
    due_date datetime not null,
    status bool not null default false,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp
);
