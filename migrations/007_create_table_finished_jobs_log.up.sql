create table if not exists finished_jobs_log
(
    queue_id uuid,
    queue_body TEXT,
    run_at timestamptz not null default current_timestamp
);