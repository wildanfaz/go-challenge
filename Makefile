install:
	go mod tidy

start:
	go run main.go start

migrate:
	go run main.go migrate

rollback:
	go run main.go rollback