-- 001_create_todos_table.sql
CREATE TABLE IF NOT EXISTS public.todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


-- Create a procedure that set the updated_at column to the current time when the row is updated

CREATE OR REPLACE FUNCTION update_todo_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';