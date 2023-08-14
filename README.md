### package

https://github.com/golang-migrate/migrate

### add database

go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### Create Migrate

migrate create -ext sql -dir db/migrations create_table_name

### Run Migrate

#### Up

migrate -database "mysql://db_user@tcp(db_host:db_port)/db_name" -path db/migrations up

#### Down

migrate -database "mysql://db_user@tcp(db_host:db_port)/db_name" -path db/migrations down

// use number to migrate each version //
