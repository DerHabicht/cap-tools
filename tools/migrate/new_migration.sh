# Need the golang-migrate CLI to execute this: https://github.com/golang-migrate/migrate/tree/v4.18.1/cmd/migrate
migrate create -ext sql -dir ../../deployments/migrations/ -seq "$1"