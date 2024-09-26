-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_order (
    order_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    current_status TEXT NOT NULL DEFAULT 'new'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_order;
-- +goose StatementEnd
