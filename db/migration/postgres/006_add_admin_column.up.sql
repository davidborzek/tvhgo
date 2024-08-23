ALTER TABLE "user"
ADD COLUMN is_admin BOOLEAN DEFAULT false;

UPDATE "user"
SET
    is_admin = true;
