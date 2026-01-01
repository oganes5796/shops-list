-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_status AS ENUM (
    'created',
    'in_progress',
    'success',
    'fail'
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id INT NOT NULL,
    shop_id INT NOT NULL,

    cart JSONB NOT NULL,
    -- cart = [{ sku, quantity, price }]

    summary_volume INT NOT NULL,
    price DOUBLE PRECISION NOT NULL,

    status order_status NOT NULL DEFAULT 'created',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
DROP TYPE IF EXISTS order_status;
-- +goose StatementEnd
