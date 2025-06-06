#name app
APP_NAME  = server

dev:
	air

run: 
	docker-compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker-compose kill

up:
	docker-compose up -d

down:
	docker-compose down

.PHONY: run