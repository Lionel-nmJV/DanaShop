ALTER TABLE products
    ALTER COLUMN image_url SET DATA TYPE VARCHAR(255),
    ALTER COLUMN image_url SET NOT NULL,
    ALTER COLUMN image_url DROP DEFAULT,
    ADD COLUMN weight INTEGER NOT NULL DEFAULT 1,
    ADD COLUMN threshold INTEGER NOT NULL DEFAULT 1,
    ADD COLUMN is_new BOOLEAN NOT NULL DEFAULT true,
    ADD COLUMN description TEXT NOT NULL DEFAULT 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.';
