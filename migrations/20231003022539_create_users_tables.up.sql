CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE roles_enum AS ENUM ('customer', 'merchant');

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email varchar(100) UNIQUE,
    password varchar(100) UNIQUE,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    roles roles_enum
);