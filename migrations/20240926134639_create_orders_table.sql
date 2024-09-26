-- +goose Up
-- +goose StatementBegin
create table order (
    order_id SERIAL,
    user_id INTEGER,
    current_status TEXT NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table order;
-- +goose StatementEnd
