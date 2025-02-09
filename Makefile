# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down
	
# для удобства добавим команду run, которая будет запускать наше приложение

migrateUsers:
	$(MIGRATE) up

migrateUsers-down:
	$(MIGRATE) down
 
run:
	go run cmd/app/main.go # Теперь при вызове make run мы запустим наш сервер

gen:
	oapi-codegen -config openapi/.openapi -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

genUsers:
	oapi-codegen -config openapi/.openapi -package Users openapi/openapiUsers.yaml > ./internal/web/Users/api.genUsers.go



