BEGIN;

CREATE TABLE IF NOT EXISTS master.shifts
(
    id         UUID         NOT NULL,
    shift_name       VARCHAR(128) NOT NULL,
    is_default       BOOLEAN NOT NULL DEFAULT FALSE,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;