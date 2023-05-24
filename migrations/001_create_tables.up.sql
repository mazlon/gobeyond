create table if not exists questions
(
    id uuid primary key default gen_random_uuid(),
    /* owner_id int not null, */
    question json not null,
    date_create timestamptz not null default current_timestamp
);