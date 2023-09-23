create table if not exists unique_ids (
    id bigint(20) not null auto_increment,
    unique_id varchar(8) not null,
    primary key (`id`)
);
