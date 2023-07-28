CREATE TABLE IF NOT EXISTS two_factor_settings (
    user_id INTEGER PRIMARY KEY,
    secret TEXT NOT NULL,
    enabled BOOLEAN NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user(id) ON DELETE CASCADE
);
