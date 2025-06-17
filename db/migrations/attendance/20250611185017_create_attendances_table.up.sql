BEGIN;

CREATE TABLE IF NOT EXISTS attendance.attendances
(
    id                      uuid,
    employee_id             uuid,
    shift_id               uuid,
    shift_start             VARCHAR(50),
    shift_end               VARCHAR(50),
    attendance_date         TIMESTAMP,
    checkin                 TIMESTAMP,
    checkout                TIMESTAMP,
    created_at              TIMESTAMP,
    created_by              VARCHAR(200),
    updated_at              TIMESTAMP,
    updated_by              VARCHAR(200),
    deleted_at              TIMESTAMP,
    deleted_by              VARCHAR(200),
    PRIMARY KEY(id),
    CONSTRAINT fk_employee FOREIGN KEY (employee_id) REFERENCES auth.users(id) ON DELETE CASCADE,
    CONSTRAINT fk_shift FOREIGN KEY (shift_id) REFERENCES master.shift_details(id) ON DELETE CASCADE
);

COMMIT;
