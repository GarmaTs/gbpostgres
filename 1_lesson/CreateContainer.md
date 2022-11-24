# Запуск Postgres в контейнере

```bash
mkdir -p $(pwd)/db/data

# Данные БД будут сохраняться в папку $(pwd)/db/data. Если в  
# ней уже есть данные БД, то при создании контейнера содержимое # папки перезаписываться не будет

docker run \
    -d \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=TheP@ssw0rd \
    -e PGDATA=/var/lib/postgresql/data \
    -v $(pwd)/db/data:/var/lib/postgresql/data \
    postgres:14.4
```

# Выполнение команд внутри контейнера
# bash будет выполняться в контейнере
```bash
docker exec \
    -it \
    --user postgres \
    postgres \
    bash   
```

# Подключение к БД с помощью psql
# Пароль указан при создании контейера
```bash
psql -h localhost -p 5432 -U postgres
```

# PSQL: метакоманды
```bash
# Вывести список БД
\l
# Список пользователей
\du
# Подключиться к БД
\с postgres
# Вывести список таблиц в БД
\d
# Вывести список локальных переменных
\set
# Закрыть psql
\q
# Список индексов
\di+
# Схема таблицы
\d+ employees
# Переключиться на более удобное представление
\x
```

# Создание пользователя
```sql
CREATE USER gopher
WITH PASSWORD 'TheP@ssw0rd';
```

# Создание типа данных semver
```sql
CREATE TYPE semver AS (
    Major integer,
    Minor integer,
    Patch integer
);

CREATE OR REPLACE FUNCTION semver(t text) RETURNS semver as $$
    DECLARE 
        t_parts text[];
    BEGIN
        t_parts = regexp_split_to_array(t, E'\\.');
        RETURN (t_parts[1]::integer, t_parts[2]::integer, t_parts[3]::integer);
    END;
$$ LANGUAGE plpgsql;


Проверка работы созданной функции
select semver('1.2.3');
select * from semver('1.2.3');

# Создание ф-ии для кастования к semver
CREATE CAST(text as semver) WITH FUNCTION semver(t text);
select semver('1.2.3'::text)::semver;
```