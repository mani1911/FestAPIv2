build:
	go build -o server ./main.go 

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

lint:
	golangci-lint run --fix --config .golangci.yml

fix:
	golangci-lint run --fix

seed_admin:
	go run ./scripts/seed_db/main.go admin

seed_database:
	go run ./scripts/seed_db/main.go db

swag_init:
	swag init

swag_fmt:
	swag fmt
