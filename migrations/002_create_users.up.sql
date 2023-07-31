create table if not exists users
(
    id uuid primary key default gen_random_uuid(),
    /* owner_id int not null, */
    username varchar(100) not null unique,
    pass varchar(200) not null 
);