-- +goose Up
-- +goose StatementBegin

CREATE TABLE projects (
    id int NOT NULL,
    name text,
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(id)
);
CREATE INDEX index_id 
ON projects (id);

-- //////////////////////////

CREATE TABLE goods (
    id int NOT NULL,
    project_id int NOT NULL,
    name text,
    description text,
    priority int,
    removed bool DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY(id)
);
CREATE INDEX index_goods
ON goods (id, project_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP INDEX index_id;
DROP INDEX index_goods;
DROP TABLE projects;
DROP TABLE goods;
-- +goose StatementEnd
