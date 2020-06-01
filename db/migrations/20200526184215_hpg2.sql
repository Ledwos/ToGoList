
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE ut_table (
    ut_id SERIAL PRIMARY KEY,
    utu_id INTEGER REFERENCES u_table(u_id),
    utt_id INTEGER REFERENCES t_table(t_id) ON DELETE CASCADE
);

ALTER TABLE u_table ALTER COLUMN u_created_on TYPE TIMESTAMPTZ;
ALTER TABLE U_TABLE ADD COLUMN u_email CITEXT UNIQUE;

ALTER TABLE t_table ALTER COLUMN t_created_on TYPE TIMESTAMPTZ;
ALTER TABLE t_table ADD COLUMN t_date DATE;
ALTER TABLE t_table ADD COLUMN t_time TIME;
ALTER TABLE t_table ADD COLUMN t_comp BOOLEAN;






-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE ut_table;

ALTER TABLE u_table ALTER COLUMN u_created_on TIMESTAMP NOT NULL;
ALTER TABLE u_table DROP COLUMN u_email;

ALTER TABLE t_table ALTER COLUMN t_created_on TIMESTAMP NOT NULL;
ALTER TABLE t_table DROP COLUMN t_date; 
ALTER TABLE t_table DROP COLUMN t_time;
ALTER TABLE t_table DROP COLUMN t_comp;