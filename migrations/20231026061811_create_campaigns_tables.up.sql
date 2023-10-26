CREATE TABLE IF NOT EXISTS "campaigns" (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	merchant_id uuid REFERENCES merchants(id) NOT NULL,
	name varchar(100) NOT NULL,
	start_date timestamp NOT NULL,
	end_date timestamp NOT NULL,
	is_active boolean NOT NULL DEFAULT true,
	video_url varchar(255) NOT NULL,
	description TEXT NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp
);