alter table if exists questions
    add 
    column user_id uuid references users(id);
