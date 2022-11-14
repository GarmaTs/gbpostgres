# /bin/bash

# Запуск:
# chmod +x restart.sh
# ./restart.sh

set -e
(docker stop postgres && docker rm postgres) || true
sudo rm -rf data

docker run \
    -d \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=P@ssw0rd \
    -e PGDATA=/var/lib/postgresql/data \
    -v $(pwd)/data:/var/lib/postgresql/data \
    -v $(pwd)/init:/docker-entrypoint-initdb.d \
    postgres:14.4