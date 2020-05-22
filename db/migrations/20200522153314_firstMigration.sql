
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE goosetab (
    username VARCHAR (50) UNIQUE NOT NULL,
    info VARCHAR (100),
    created TIMESTAMP NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE goosetab;

