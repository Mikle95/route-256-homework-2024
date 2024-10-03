-- +goose Up
-- +goose StatementBegin
CREATE TABLE item (
    sku INTEGER NOT NULL REFERENCES stock (sku),
    order_id INTEGER NOT NULL REFERENCES user_order (order_id),
    count INTEGER NOT NULL,
    unique (sku, order_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE item;
-- +goose StatementEnd
