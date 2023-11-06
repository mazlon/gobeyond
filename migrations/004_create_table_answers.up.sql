create table if not exists answers
(
    id uuid primary key default gen_random_uuid(),
    /* owner_id int not null, */
    answer json not null,
    question_id uuid references questions(id),
    date_create timestamptz not null default current_timestamp
);