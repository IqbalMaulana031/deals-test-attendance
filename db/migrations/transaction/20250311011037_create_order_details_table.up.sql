BEGIN;
CREATE TABLE transaction.order_details (
    created_by          VARCHAR(128) NOT NULL,
    updated_by          VARCHAR(128) NOT NULL,
    deleted_by          VARCHAR(128),
    created_at          TIMESTAMPTZ NOT NULL,
    updated_at          TIMESTAMPTZ NOT NULL,
    deleted_at          TIMESTAMPTZ,
    id                  UUID NOT NULL,
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    qty INT DEFAULT 0,
    amount INT DEFAULT 0
    PRIMARY KEY (id)
);
COMMIT;