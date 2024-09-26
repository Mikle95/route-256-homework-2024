-- +goose Up
-- +goose StatementBegin
create table item (
    sku INTEGER,
    order_id INTEGER,
    count INTEGER,
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table item;
-- +goose StatementEnd
