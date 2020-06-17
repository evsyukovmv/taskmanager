CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR (500) NOT NULL,
    description VARCHAR (5000),
    position INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    column_id INTEGER NOT NULL REFERENCES columns(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (position, column_id)
);
