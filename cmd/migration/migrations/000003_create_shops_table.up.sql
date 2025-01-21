CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Enables UUID generation

CREATE TABLE shops (
    shop_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique shop ID as UUID
    user_id UUID NOT NULL, -- User ID (foreign key reference)
    shop_address TEXT, -- Shop address
    shop_name VARCHAR(100) NOT NULL, -- Shop name
    shop_images TEXT[] DEFAULT '{}', -- Array of image URLs
    shop_description TEXT, -- Description of the shop
    products TEXT[] DEFAULT '{}', -- Array of product names or IDs
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Last update time
);
