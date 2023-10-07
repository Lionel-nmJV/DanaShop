CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    merchant_id uuid REFERENCES merchants(id) NOT NULL,
    name varchar(100) NOT NULL,
    category varchar(100) NOT NULL,
    price numeric NOT NULL,
    stock integer NOT NULL,
    image_url varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp
)