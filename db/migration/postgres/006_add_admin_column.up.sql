ALTER TABLE "user"
ADD COLUMN is_admin BOOLEAN DEFAULT FALSE;

UPDATE "user"
SET
    is_admin = TRUE;