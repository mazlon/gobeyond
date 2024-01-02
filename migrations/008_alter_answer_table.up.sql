alter table if exists answers
alter column answer type varchar(8192);
-- based on last update in chatgpt limitation of input tokens: 8192