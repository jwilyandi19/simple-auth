migrate-up:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/auth_db" -verbose up
	go run cmd/seed/app.go

migrate-down:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/auth_db" -verbose down

start:
	go run cmd/main/app.go

run-db:
	docker pull mysql/mysql-server:latest
	docker run --name simple-auth-1 -e MYSQL_USER=root -e MYSQL_DATABASE=auth_db -e MYSQL_ROOT_PASSWORD=root -e MYSQL_PASSWORD=root -p 3306:3306 -d mysql/mysql-server:latest