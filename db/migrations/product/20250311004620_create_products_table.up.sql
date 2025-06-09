BEGIN;
CREATE TABLE product.products (
    created_by VARCHAR(128) NOT NULL,
    updated_by VARCHAR(128) NOT NULL,
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    id UUID NOT NULL,
    product_name VARCHAR(128) NOT NULL,
    stock INT DEFAULT 0,
    price INT DEFAULT 0,
    PRIMARY KEY (id)
);
COMMIT;