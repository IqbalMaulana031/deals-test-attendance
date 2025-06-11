BEGIN;
    ALTER TABLE auth.users 
    ADD COLUMN company_id uuid NULL;
COMMIT;