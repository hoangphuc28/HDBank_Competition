.PHONY: createdb migrate_up migrate_down

# Go migrate
create_mysql:
	docker run --name go01-mysql -e MYSQL_ROOT_PASSWORD=Abc123! -p 3307:3306 -d mysql:latest
migrate_up:
	migrate -path migrations -database "mysql://root:Abc123!@tcp(127.0.0.1:3307)/HDbank" up 1

migrate_down:
	migrate -path migrations -database "mysql://root:Abc123!@tcp(127.0.0.1:3307)/HDbank" down 1

# Main
run:
	go run main.go

build:
	go build main.go