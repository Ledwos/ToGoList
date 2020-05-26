
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE u_table (
    u_id serial PRIMARY KEY,
    u_name VARCHAR (50) UNIQUE NOT NULL,
    u_pass VARCHAR (50) NOT NULL,
    u_created_on TIMESTAMP NOT NULL
);

CREATE TABLE t_table (
    t_id serial PRIMARY KEY,
    t_name VARCHAR (50) NOT NULL,
    t_desc VARCHAR (200),
    t_created_on TIMESTAMP NOT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE u_table;
DROP TABLE t_table;

