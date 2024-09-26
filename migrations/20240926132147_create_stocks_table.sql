-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock (
    sku INTEGER PRIMARY KEY,
    total_count int NOT NULL,
    reserved int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock;
-- +goose StatementEnd
