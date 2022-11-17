
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
CREATE INDEX ON snippets(header text_pattern_ops); /* for index scan when filtering header with LIKE operator*/


DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    tag VARCHAR(200) UNIQUE
);
CREATE INDEX ON tags(tag text_pattern_ops);


DROP TABLE IF EXISTS snippets_tags;
CREATE TABLE snippets_tags (
    snippet_id INT,
    tag_id INT,
    FOREIGN KEY (snippet_id) REFERENCES snippets (id),
    FOREIGN KEY (tag_id) REFERENCES tags (id), 
    PRIMARY KEY (snippet_id, tag_id)
);


INSERT INTO snippets (header, content)
VALUES 
    ('CAP теорема', 'Consistensy, Availability, Partition Tolerance. Можно выбрать только 2 из 3 свойств. CA - нераспределенный сервер, масштабируемый вертикально. CP/AP - распределенная система, масштабируемая горизонтально.');

INSERT INTO tags (tag)
VALUES ('CAP теорема'), 
    ('Базы данных');

INSERT INTO snippets_tags (snippet_id, tag_id)
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

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'lesson 2 Recap'
    and t.tag in ('DDL', 'DML', 'CTE', 'CONSTRAINTS', 'Insert Overriding system value');

INSERT INTO snippets (header, content)    
VALUES
    ('Create Index', '1) Создание индекса. Добавляется командой CREATE INDEX
При создании индекса можно явно указать его тип, по умолчанию BTree (balanced tree). 
На время создания индекса данные в таблице доступны только для чтения.
Чтобы избежать этого можно использовать CREATE INDEX CONCURRENTLY
Пример:
CREATE INDEX ON employees(email)

2) Доступные классы операторов
select am.amname as index_method,
    opc.opcname as opclass_name,
    opc.opcintype::regtype as indexed_type,
    opc.opcdefault as is_default
from pg_am am, pg_opclass opc
where opc.opcmethod = am.oid
order by index_method, opclass_name;

3) Чтоб работал index scan при операторе like 
нужно указать класс оператора (например text_pattern_ops);
CREATE INDEX ON employees(email text_pattern_ops);

4) Covering index (покрывающий индекс)
CREATE INDEX ON employees(email text_pattern_ops)
INCLUDE(phone);
При таком индексе запрос вида:
select email, phone from employees where email = bla-bla;
получит данные из индекса и не будет заглядывать в таблицу.
expain analyze выведет Index Only Scan, вместо Index Scan

5) Порядок индексации
Если при использовании сортировки (order by) expain выводит Seq Scan, 
то можно добавить индекс 
CREATE INDEX first_name_last_name_desc
ON employees (first_name, last_name);
Тогда expain будет выводить Index Scan

6) Частичная индексация
Подмножество, поиск по которому будет работать быстрее
CREATE INDEX alice_employees
ON employees (manager) WHERE (manager = 3);

7) Индексация выражений
CREATE INDEX first_name_lower
ON employees(lower(first_name));
Такой индекс создается для запросов вида:
select last_name from employees
where lower(first_name) = blabla;');

INSERT INTO tags (tag)
VALUES
    ('Create Index'), ('Создание индекса'), ('Доступные классы операторов'),
    ('text_pattern_ops'), ('Покрывающий индекс'), ('Порядок индексации'),
    ('Частичная индексация'), ('Индексация выражений');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Create Index'
    and t.tag in ('Create Index', 'Создание индекса', 
        'Доступные классы операторов', 'text_pattern_ops', 
        'Покрывающий индекс', 'Порядок индексации',
        'Частичная индексация', 'Индексация выражений');

INSERT INTO snippets (header, content)
VALUES
    ('Analyze', 'Статистика хранится в представлении pg_stats
Обновляет статистику vacuum, но можно пересчитать статистику
командой analyze [tablename];');

INSERT INTO tags (tag)
VALUES ('analyze'), ('pg_stats'), ('обновление статистики');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Analyze'
    and t.tag in ('analyze', 'pg_stats', 'обновление статистики');

INSERT INTO snippets (header, content)
VALUES('Explain', '1) Получение плана запроса (без фактического выполнения запроса)
EXPLAIN
SELECT COUNT(*) FROM employees;
2) Построение плана и выполнение запроса
EXPLAIN ANALYZE
SELECT COUNT(*) FROM employees');

INSERT INTO tags (tag)
VALUES('explain'), ('explain analyze');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Explain'
    and t.tag in ('explain', 'explain analyze');

INSERT INTO snippets (header, content)
VALUES ('Random', '1) Получение случайной строки
select * from employees
order by random()
limit 1;');

INSERT INTO tags (tag)
VALUES ('Получение случайной строки'), ('random');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Random'
    and t.tag in ('Получение случайной строки', 'random');

INSERT INTO snippets (header, content)
VALUES ('Типы индексов', 'Типы индексов
1) BTree - выбор по умолчанию, подходит под все типы данных.
2) GIN - позволяет реализовать качественный полнотекстовый
поиск, умеет индексировать JSONB и временные интервалы.
3) GiST + SP-GiST - индексация геометрических объектов. Это 
позволяет быстро выполнять запросы, как поиск ближайшей к 
пользователю машины такси или нахождение всех точек в 
заданной области.
4) BRIN - создавался для работы с огромными наборами данных.
Основная цель индекса - возможность пропустить просмотр
заведомо ненужных страниц. Подходит для данных, которые лежат 
в таблице в уже почти упорядоченном виде, например, 
timeseries данные.
5) Hash - позволяет очень быстро выполнять выражения на 
поиск полного совпадения.');

INSERT INTO tags (tag)
VALUES ('Типы индексов'), ('BTree'), ('GIN'), ('GiST'), ('BRIN');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Типы индексов'
    and t.tag in ('Типы индексов', 'BTree', 'GIN', 'GiST', 'BRIN');

INSERT INTO snippets (header, content)
VALUES ('Способы соединения таблиц', 'Способы соединения таблиц.
1) Nested Loops - на каждую строку из одной таблицы
происходит полное сканирование второй таблицы.
2) Merge Join - обе таблицы сортируются перед началом
сканирования, а затем происходит сканирование обеих
таблиц за один проход: используется дополнительная 
память. При этом данные, участвующие в сортировке, 
могут быть вытеснены на диск, если не помещаются в 
оперативную память.
3) Has Join - перед началом работы в оперативной
памяти строится hash-таблица, по которой в последующем
за один проход получаются данные.');

INSERT INTO tags (tag)
VALUES ('Способы соединения таблиц'), ('Nested Loops'), ('Merge Join'),
    ('Has Join');

INSERT INTO snippets_tags (snippet_id, tag_id)
SELECT 
    s.id snippet_id, t.id tag_id
FROM snippets s, tags t
WHERE
    s.header = 'Способы соединения таблиц'
    and t.tag in ('Способы соединения таблиц', 'Nested Loops', 
    'Merge Join', 'Has Join');