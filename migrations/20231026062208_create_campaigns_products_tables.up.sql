CREATE TABLE IF NOT EXISTS "campaigns_products" (
    campaign_id uuid REFERENCES campaigns(id) NOT NULL,
    product_id uuid REFERENCES products(id) NOT NULL,
    campaign_price NUMERIC NOT NULL,
    PRIMARY KEY (campaign_id, product_id)
);