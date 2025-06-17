BEGIN;

CREATE TABLE IF NOT EXISTS master.shift_details
(
    id         UUID         NOT NULL,
    shift_id   UUID         NOT NULL,
    code       VARCHAR(5) NOT NULL,
    day_type       BOOLEAN NOT NULL DEFAULT TRUE,
    day_in_number    INTEGER NOT NULL DEFAULT 1,
    start_time                  TIME,
    end_time                    TIME,
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ  NOT NULL,
    updated_at TIMESTAMPTZ  NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    CONSTRAINT fk_shift FOREIGN KEY (shift_id) REFERENCES master.shifts(id) ON DELETE CASCADE
    );

COMMIT;