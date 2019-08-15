start: start-db start-server

start-db:
	docker-compose up

start-server:
	go run main.go
