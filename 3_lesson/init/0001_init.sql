CREATE USER gopher
WITH PASSWORD 'P@ssw0rd';

CREATE DATABASE gopher_corp
    WITH OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

\c gopher_corp

SET ROLE gopher;

DROP TABLE IF EXISTS departments CASCADE;
CREATE TABLE departments (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    parent_id INT,
    name VARCHAR(200)
);

DROP TABLE IF EXISTS employees CASCADE;
CREATE TABLE employees (
    id INT GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(200),
    last_name VARCHAR(200),
    phone TEXT,
    email TEXT,
    salary MONEY,
    manager_id INT,
    department_id INT,
    position INT
);

DROP TABLE IF EXISTS positions CASCADE;
CREATE TABLE positions (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(200)
);
