# drop snippets db
```bash
DROP DATABASE snippets;
```
# create snippets db
```bash
CREATE DATABASE snippets
    WITH OWNER gopher
    TEMPLATE = 'template0'
    ENCODING = 'utf-8'
    LC_COLLATE = 'C.UTF-8'
    LC_CTYPE = 'C.UTF-8';
```

# setup migrations
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 $GOPATH/bin/migrate
```
# create first migration
```bash
migrate create -seq -ext sql -dir migrations init_schema
```
# create second migration
```bash
migrate create -seq -ext sql -dir migrations add_column_main_theme
```
# run migrations
```bash
migrate -database "postgresql://gopher:P@ssw0rd@localhost:5432/snippets?sslmode=disable" -path migrations -verbose up
```

# init integration test
```bash
DATABASE_URL=postgres://gopher:P@ssw0rd@localhost:5432/snippets go test --tags=integration ./pkg/storage
```
