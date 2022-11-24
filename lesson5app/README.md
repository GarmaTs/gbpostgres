# init integration test
```bash
DATABASE_URL=postgres://gopher:P@ssw0rd@localhost:5432/gopher_corp go test ./... --tags=integration ./pkg/emailhint/storage
```

# setup migrations
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 $GOPATH/bin/migrate

# create first migration
migrate create -seq -ext sql -dir migrations init_schema

# create test db for migrations
CREATE DATABASE lesson5db
    WITH OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';

# run migration
migrate -database "postgresql://gopher:P@ssw0rd@localhost:5432/lesson5db?sslmode=disable" -path migrations -verbose up

# create second migration
migrate create -seq -ext sql -dir migrations add_column_main_theme

