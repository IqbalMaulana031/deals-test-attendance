BEGIN;
CREATE TABLE transaction.orders (
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    amount INT DEFAULT 0
    PRIMARY KEY (id)
);
COMMIT;