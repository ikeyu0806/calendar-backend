
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id integer AUTO_INCREMENT,
    name varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
