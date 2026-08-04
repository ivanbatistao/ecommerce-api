-- +goose Up
ALTER TABLE order_items
ADD COLUMN price_in_cents INTEGER NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE order_items
DROP COLUMN price_in_cents;