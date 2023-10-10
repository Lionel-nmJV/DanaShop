include .env

run :
	go run cmd/api/main.go

up :
	docker-compose up -d

up-dev :
	docker-compose -f docker-compose-dev.yaml up -d

stop : 
	docker-compose stop

stop-dev :
	docker-compose -f docker-compose-dev.yaml stop

create-migration :
	migrate create -ext sql -dir migrations $(name)

migrate-up :
	migrate -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path migrations up

migrate-down :
	migrate -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path migrations down