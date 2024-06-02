-- +goose Up
DROP TABLE IF EXISTS authors, posts;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) DEFAULT 0,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO authors (name) VALUES ('Автор 1');
INSERT INTO authors (name) VALUES ('Автор 2');
INSERT INTO authors (name) VALUES ('Автор 3');

INSERT INTO posts (author_id, title, content) VALUES (1, 'Заголовок поста 1', 'Содержание поста 1');
INSERT INTO posts (author_id, title, content) VALUES (2, 'Заголовок поста 2', 'Содержание поста 2');
INSERT INTO posts (author_id, title, content) VALUES (3, 'Заголовок поста 3', 'Содержание поста 3');


-- +goose Down
DROP TABLE posts;
DROP TABLE authors;
