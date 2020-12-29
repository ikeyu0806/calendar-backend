
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE schedules (
    id integer AUTO_INCREMENT,
    user_id integer,
    start_at datetime,
    end_at datetime,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE schedules;
