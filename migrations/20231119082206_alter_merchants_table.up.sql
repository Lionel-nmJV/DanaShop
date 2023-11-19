ALTER TABLE merchants
    ADD COLUMN focused_on varchar(100) NOT NULL,
    ADD COLUMN address varchar(255) NOT NULL;