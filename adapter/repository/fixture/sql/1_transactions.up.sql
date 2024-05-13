CREATE TABLE transactions (
    id TEXT PRIMARY KEY,
    account TEXT NOT NULL,
    amount REAL NOT NULL,
    status TEXT NOT NULL,
    error_message TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);