CREATE TABLE IF NOT EXISTS "session" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    hashed_token TEXT UNIQUE NOT NULL,
    client_ip TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    last_used_at INTEGER,
    rotated_at INTEGER NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES "user"(id) ON DELETE CASCADE
)
