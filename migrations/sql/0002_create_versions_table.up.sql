CREATE TABLE IF NOT EXISTS versions (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES services(id) ON DELETE CASCADE,
    version_number VARCHAR(50),
    created_at TIMESTAMP DEFAULT NOW()
);