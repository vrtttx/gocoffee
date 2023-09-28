DSN="host=localhost port=5432 user=root password=secret dbname=gocoffeedb sslmode=disable timezone=UTC connect_timeout=5"
PORT=8080

DB_DOCKER_CONTAINER=gocoffee_db
BINARY_NAME=gocoffee

# creating the container with postgres software
postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# creating the database inside the postgres container
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root gocoffeedb

# stop other docker containers
stop-containers:
	@echo "stopping other docker containers"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no active containers found..."; \
	fi

# start docker container
start-docker:
	docker start ${DB_DOCKER_CONTAINER}

create-migrations:
	sqlx migrate add -r init

migrate-up:
	sqlx migrate run --database-url "postgres://root:secret@localhost:5432/gocoffeedb?sslmode=disable"

migrate-down:
	sqlx migrate revert --database-url "postgres://root:secret@localhost:5432/gocoffeedb?sslmode=disable"

build:
	@echo "building backend api binary"
	go build -o ${BINARY_NAME} cmd/server/*.go
	@echo "binary built!"

run: build stop-containers start-docker
	@echo "starting api"
	@env PORT=${PORT} DSN=${DSN} ./${BINARY_NAME} &
	@echo "api started!"

stop:
	@echo "stopping api"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "api stopped!"

start: run

restart: stop start