# Список выполняющихся конейтеров
docker ps

# Список всех контейнеров
docker ps -a

# Остановить выполнение контейнера
docker stop [id] (первые 5 символов айдишника)
docker stop [container_name]

# Запуск контейнера 
docker start [container_name]

# Перезапустить контейнер
docker restart [container_name]

# Удаление существующего конейнера
docker rm [container_name]

# Логи контейнера
docker logs [container_name]
