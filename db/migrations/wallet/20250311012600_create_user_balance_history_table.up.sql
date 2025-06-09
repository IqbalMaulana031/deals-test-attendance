BEGIN;
CREATE TABLE wallet.balance_history (
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    id UUID NOT NULL,
    wallet_id UUID NOT NULL,
    balance INT DEFAULT 0
    detail_action VARCHAR(128) NOT NULL,
    PRIMARY KEY (id)
);
COMMIT;