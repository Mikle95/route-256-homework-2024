-- +goose Up
-- +goose StatementBegin
create table stock (
    sku INTEGER PRIMARY KEY,
    total_count int NOT NULL,
    reserved int NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table stock;
-- +goose StatementEnd
