BEGIN;
CREATE TABLE IF NOT EXISTS "auth".users (
    id UUID NOT NULL,
    username VARCHAR(128) NOT NULL,
    password TEXT NOT NULL,
    sallary INTEGER NOT NULL DEFAULT 0,
    role_id UUID NOT NULL,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id), 
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES auth.roles(id) ON DELETE CASCADE
);
COMMIT;