ALTER TABLE campaigns_products DROP CONSTRAINT IF EXISTS campaigns_products_campaign_id_fkey;
ALTER TABLE campaigns_products DROP CONSTRAINT IF EXISTS campaigns_products_product_id_fkey;

DROP TABLE IF EXISTS campaigns_products;