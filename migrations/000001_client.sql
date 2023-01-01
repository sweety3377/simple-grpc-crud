CREATE TABLE IF NOT EXISTS clients (
    id TEXT UNIQUE,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    age SMALLINT NOT NULL,
    weight SMALLINT NOT NULL,
    height SMALLINT NOT NULL
);

CREATE UNIQUE INDEX idx_clients_id ON clients(id);