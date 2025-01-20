CREATE TABLE license (
    license_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Unique identifier for the license
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE, -- Foreign key referencing users
    validity TIMESTAMP NOT NULL, -- Validity timestamp
    license_type VARCHAR(100) DEFAULT 'FREE', -- License type with default value
    shop_limit INT DEFAULT 1, -- Shop limit with default value
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Token creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Automatic update on row modification
);