CONSTRAINTS

1) CHECK(логическое выражение)
CONSTRAINT employees_salary_positive_check check(salary > 0)

2) NOT NULL или NULL
manager_id INT NOT NULL

3) FOREIGN KEY
CONSTRAINT employees_fk_manager_id foreign key (manager_id) references
    employees (id) on delete restrict
-- on delete restrict указывает, что при удалении менеджера, у которого есть
-- подчиненные, произойдет ошибка    

4) UNIQUE

5) PRIMARY KEY

6) EXCLUSION 
работает на уровне таблицы и позволяет убедиться что при
сравнении любых двух строк на указанных значениях столбцов
с использованием указанных операторов хотя бы один 
результат - false
CREATE TABLE circles (
    c circle,
    EXCLUDE USING gist(c WITH &&)
);
