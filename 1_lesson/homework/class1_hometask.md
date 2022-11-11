# 2. Create user and DB
```sql
CREATE USER gopher
WITH PASSWORD 'TheP@ssw0rd';
```

```sql
CREATE DATABASE snippets
    OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

psql --host 127.0.0.1 --port 5432 --username gopher --dbname snippets
```

# 3. Create tables
```sql
CREATE TABLE snippets (
    id INT GENERATED ALWAYS AS IDENTITY,
    header VARCHAR(200) UNIQUE,
    content text
);

CREATE TABLE tags (
    id INT GENERATED ALWAYS AS IDENTITY,
    tag VARCHAR(200) UNIQUE,
    snippet_id INT
);
```

# 4. Insert some data into created tables
```sql
INSERT INTO snippets (header, content)
VALUES 
    ('CAP теорема', 'Consistensy, Availability, Partition Tolerance. Можно выбрать только 2 из 3 свойств. CA - нераспределенный сервер, масштабируемый вертикально. CP/AP - распределенная система, масштабируемая горизонтально.');

INSERT INTO tags (tag, snippet_id)
VALUES 
    ('CAP теорема', 1),
    ('Базы данных', 1);
```

# 5. List entities using psql metacommands
```psql
/d
List of relations
 Schema |      Name       |   Type   | Owner  
--------+-----------------+----------+--------
 public | snippets        | table    | gopher
 public | snippets_id_seq | sequence | gopher
 public | tags            | table    | gopher
 public | tags_id_seq     | sequence | gopher
```

# 6.7. Project description
```txt
Проект для создания кратких заметок. Пока схема содержит
две таблицы: заметки и тэги. По мере развития проекта 
возможно добавление таблиц
```

# 8. Create scheme
```txt
Реализовано в пунктах 3,4
```