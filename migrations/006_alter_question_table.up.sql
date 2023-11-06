alter table if exists questions
alter column question type varchar(8192);
-- based on last update in chatgpt limitation of input tokens: 8192