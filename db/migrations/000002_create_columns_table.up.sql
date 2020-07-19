CREATE TABLE IF NOT EXISTS columns(
    id SERIAL PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    position INTEGER DEFAULT 0,
    project_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (name, project_id),
    UNIQUE (position, project_id)
);
