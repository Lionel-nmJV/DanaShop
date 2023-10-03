CREATE TABLE IF NOT EXISTS merchants (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users(id) NOT NULL,
    name varchar(100) UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    image_url text DEFAULT 'https://iconlogovector.com/uploads/images/2023/03/lg-8b694ad63ac093b84ec0972a7c442d7c61.jpg'
)