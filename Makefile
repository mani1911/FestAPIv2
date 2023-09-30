build:
	go run /app/main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

seed_admin:
	go run ./scripts/seed_db/main.go admin

seed_database:
	go run ./scripts/seed_db/main.go db