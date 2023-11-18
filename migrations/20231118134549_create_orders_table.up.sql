CREATE TYPE order_status AS ENUM (
    'success',
    'failed',
    'pending'
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    merchant_id uuid REFERENCES merchants(id) NOT NULL,
    user_id uuid REFERENCES users(id) NOT NULL,
    user_name varchar(100) NOT NULL,
    user_address varchar(255) NOT NULL,
    user_phone_number varchar(15) NOT NULL,
    created_at timestamp NOT NULL,
    product_id uuid REFERENCES products(id) NOT NULL,
    product_name varchar(100) NOT NULL,
    product_image_url varchar(255),
    quantity int NOT NULL,
    product_price numeric NOT NULL,
    sub_total_price numeric NOT NULL,
    additional_fee numeric NOT NULL,
    shipping_cost numeric NOT NULL,
    total_discount numeric NOT NULL,
    total_price numeric NOT NULL,
    status order_status NOT NULL,
    invoice_id varchar,
    invoice_url varchar
    );