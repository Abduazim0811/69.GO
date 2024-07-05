CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    title VARCHAR(59) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN DEFAULT False
);