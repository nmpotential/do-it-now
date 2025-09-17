-- 001_create_tasks.down.sql
-- Down migration: Drops task_items and tasks tables and related indexes

DROP INDEX IF EXISTS idx_task_items_incomplete;
DROP INDEX IF EXISTS idx_task_items_task_id;
DROP TABLE IF EXISTS task_items;
DROP TABLE IF EXISTS tasks;