-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
COMMENT ON DATABASE trade_bot IS 'Data associated with users actions on trade-bot';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    uid           uuid not null unique default uuid_generate_v4(),
    name          text not null,
    username      text not null unique,
    password_hash text not null
);
COMMENT ON TABLE users IS 'Stands for storing users credentials and metadata';

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
