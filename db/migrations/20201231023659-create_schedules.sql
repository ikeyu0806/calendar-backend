
-- +migrate Up
CREATE TABLE IF NOT EXISTS schedules (
    id integer AUTO_INCREMENT,
    user_id integer,
    title varchar(255),
    content varchar(255),
    memo varchar(255),
    start_at datetime,
    end_at datetime,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY(id)
);

-- +migrate Down
DROP TABLE IF EXISTS schedules;
