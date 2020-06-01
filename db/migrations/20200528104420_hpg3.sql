
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE u_table ALTER COLUMN u_created_on SET DEFAULT NOW();

ALTER TABLE t_table ALTER COLUMN t_created_on SET DEFAULT NOW();


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

