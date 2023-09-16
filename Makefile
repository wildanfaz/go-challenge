install:
	go mod tidy

start:
	go run main.go start

migrate:
	go run main.go migrate

rollback:
	go run main.go rollback

add_balance:
	go run main.go add_balance --email wildan123@gmail.com

dumy:
	go run main.go dumy