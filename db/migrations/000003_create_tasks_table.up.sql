CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR (500) NOT NULL,
    description VARCHAR (5000),
    position INTEGER DEFAULT 0,
    column_id INTEGER NOT NULL REFERENCES columns(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (position, column_id)
);
