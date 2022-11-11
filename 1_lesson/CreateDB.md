# Создание БД курса, в кот-й будет храниться инфа 
# о сотрудниках компании gopher_corp

```sql
CREATE DATABASE gopher_corp
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';
```

# Подключение к БД с ранее созданным пользователем gopher_corp
```bash
psql --host 127.0.0.1 --port 5432 --username gopher --dbname gopher_corp
```

# Убрать все права пользователей группы PUBLIC на БД gopher_corp
1) подключиться под супеюзером
```bash
psql --host 127.0.0.1 --port 5432 --username postgres
```
2) убрать права
```sql
REVOKE ALL ON DATABASE gopher_corp FROM public;
```

# Дать права юзеру gopher на БД gopher_corp
1) подключиться под супеюзером
```bash
psql --host 127.0.0.1 --port 5432 --username postgres
```
2) дать права
```sql
ALTER DATABASE gopher_corp OWNER TO gopher;
```

# Создание таблиц
```sql 
CREATE TABLE departments (
    id INT GENERATED ALWAYS AS IDENTITY,
    parent INT,
    name VARCHAR(200)
);

CREATE TABLE employees (
    id INT GENERATED ALWAYS AS IDENTITY,
    first_name VARCHAR(200),
    last_name VARCHAR(200),
    salary MONEY,
    manager INT,
    department INT,
    position INT
);

CREATE TABLE positions (
    id INT GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(200)
);
```

# Добавление данных в созданные таблицы
```sql
INSERT INTO departments(parent, name)
VALUES
    (1, 'root'),
    (1, 'R&D');

INSERT INTO employees (first_name, last_name, salary, manager, department, position)
VALUES
    ('Jane', 'Doe', 75000.00, NULL, 2, 2),
    ('John', 'Doe', 50000.00, NULL, 2, 1);

INSERT INTO positions(title)
VALUES 
    ('Software Engineer II'),
    ('Software Engineer III');
```