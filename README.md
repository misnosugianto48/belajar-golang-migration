### package

https://github.com/golang-migrate/migrate

### add database

go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## Create Migrate

migrate create -ext sql -dir db/migrations create_table_name

## Run Migrate

migrate -database "mysql://root@tcp(localhost:3306)/belajar-golang-migration" -path db/migrations up
