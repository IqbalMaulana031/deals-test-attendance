BEGIN;
ALTER TABLE auth.users 
    DROP COLUMN company_id;
COMMIT;