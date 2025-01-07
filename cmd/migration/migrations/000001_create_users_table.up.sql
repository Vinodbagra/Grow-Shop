CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- Enables UUID generation

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- User ID as UUID
    user_name VARCHAR(100), -- User name (optional)
    email VARCHAR(255) NOT NULL UNIQUE, -- Email must be unique
    mobile_no VARCHAR(15), 
    address TEXT, -- Address (optional)
    password VARCHAR(255) NOT NULL, -- Encrypted password
    business_name VARCHAR(255), -- Business name (optional)
    gender VARCHAR(10) DEFAULT 'Other', -- Gender (optional)
    shops UUID[] DEFAULT '{}', -- List of shop UUIDs
    license_id UUID, -- Single UUID for license
    images TEXT[] DEFAULT '{}', -- List of image URLs
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation time
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Last update time
);