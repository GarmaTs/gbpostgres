
CREATE DATABASE snippets
    WITH OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

\c snippets

SET ROLE gopher;

DROP TABLE IF EXISTS snippets;
CREATE TABLE snippets (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    header VARCHAR(200) UNIQUE,
    content text
);

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tag VARCHAR(200) UNIQUE
);

DROP TABLE IF EXISTS snippets_tags;
CREATE TABLE snippets_tags (
    snippet_id INT,
    tag_id INT,
    FOREIGN KEY (snippet_id) REFERENCES snippets (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id)
);

INSERT INTO snippets (header, content)
VALUES 
    ('CAP теорема', 'Consistensy, Availability, Partition Tolerance. Можно выбрать только 2 из 3 свойств. CA - нераспределенный сервер, масштабируемый вертикально. CP/AP - распределенная система, масштабируемая горизонтально.');

INSERT INTO tags (tag)
VALUES ('CAP теорема'), 
    ('Базы данных');

INSERT INTO snippets_tags 
    (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'CAP теорема'
    and t.tag in ('CAP теорема', 'Базы данных');


INSERT INTO snippets (header, content)
VALUES 
    ('lesson 2 Recap', '1) Data Definition Language (DDL) и Data Manipulation Language (DML).\n 2) Insert Overriding system value.\n 
    3) Ограничения целостности (constraints).\n 4) Common table expression (CTE или CTE with queries)\n
    4) Оценка занимаемого места');

INSERT INTO tags (tag)
VALUES ('DDL'), ('DML'), ('CTE'), ('CONSTRAINTS'), ('Insert Overriding system value');

INSERT INTO snippets_tags 
    (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'lesson 2 Recap'
    and t.tag in ('DDL', 'DML', 'CTE', 'CONSTRAINTS', 'Insert Overriding system value');