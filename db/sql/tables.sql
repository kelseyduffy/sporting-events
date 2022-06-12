
-- create the table for teams and the corresponding table for their name history

CREATE TABLE IF NOT EXISTS teams(
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
founded_year TEXT NOT NULL,
dissolved_year TEXT,
sport TEXT NOT NULL
);