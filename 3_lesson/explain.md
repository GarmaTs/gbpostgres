# Получение плана запроса (без фактического выполнения запроса)
# перед запросом нужно ввести ключевое слово explain
EXPLAIN
SELECT COUNT(*) FROM employees;

# Получение случайной строки
select * from employees
order by random()
limit 1;

# Построение плана и выполнение запроса
EXPLAIN ANALYZE
SELECT COUNT(*) FROM employees;