start: start-db start-server start-client

start-db:
	docker-compose up

start-server:
	go run main.go

start-client:
	cd client && yarn && yarn start
