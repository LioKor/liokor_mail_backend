CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE SET NULL,
    file_path VARCHAR(256) NOT NULL,
    created_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(file_path)
);