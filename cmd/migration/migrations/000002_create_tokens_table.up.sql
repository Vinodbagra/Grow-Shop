CREATE TABLE tokens (
    token UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Token as UUID with auto-generated value
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE, -- Foreign key referencing users
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Token creation timestamp
    expires_at TIMESTAMP NOT NULL -- Token expiration timestamp
);