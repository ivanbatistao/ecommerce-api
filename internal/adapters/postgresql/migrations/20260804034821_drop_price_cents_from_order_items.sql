-- +goose Up
ALTER TABLE order_items
DROP COLUMN price_cents;

-- +goose Down
ALTER TABLE order_items
ADD COLUMN price_cents INTEGER NOT NULL DEFAULT 0;