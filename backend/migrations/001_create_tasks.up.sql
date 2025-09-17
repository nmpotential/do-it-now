-- 001_create_tasks.sql
-- Migration: Create tasks and task_items tables for Do It Now backend
-- - tasks: top-level checklist/task
-- - task_items: individual items within a task
-- - Foreign key: task_items.task_id → tasks.id
-- - Indexes for fast lookup
-- Notes:
--   - Using UUID primary keys for scalability and client integration
--   - Consistent timestamps for auditing

-- Create tasks table
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),         -- unique task ID
    user_id UUID NOT NULL,                                 -- owner of the task
    title TEXT NOT NULL,                                   -- task title
    created_at TIMESTAMPTZ DEFAULT now(),                  -- creation timestamp
    updated_at TIMESTAMPTZ DEFAULT now()                   -- last update timestamp
);

-- Create task_items table
CREATE TABLE task_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),         -- unique item ID
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE, -- parent task
    description TEXT NOT NULL,                             -- item description
    completed BOOLEAN DEFAULT FALSE,                       -- completion status
    created_at TIMESTAMPTZ DEFAULT now(),                  -- creation timestamp
    updated_at TIMESTAMPTZ DEFAULT now()                   -- last update timestamp
);

-- Index for fast lookup of items by task
CREATE INDEX idx_task_items_task_id ON task_items(task_id);

-- Partial index for fast lookup of incomplete items
CREATE INDEX idx_task_items_incomplete
ON task_items(task_id)
WHERE completed = FALSE;