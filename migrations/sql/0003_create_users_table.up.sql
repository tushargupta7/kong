CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(255) NOT NULL UNIQUE,
                                     password VARCHAR(255) NOT NULL,  -- You can store a hashed password here
                                     role VARCHAR(50) NOT NULL,        -- Role: 'admin' or 'user'
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);