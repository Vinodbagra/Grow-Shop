CREATE TABLE IF NOT EXISTS tokens (
    user_id uuid PRIMARY KEY,                         -- Primary key and foreign key
    token uuid NOT NULL UNIQUE,                       -- The token value
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP, -- Token creation time
    expires_at timestamptz NOT NULL,                  -- Token expiration time
    FOREIGN KEY (user_id) REFERENCES user_tokens(id) ON DELETE CASCADE
);