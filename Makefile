schema:
	go run github.com/99designs/gqlgen generate

stop-local-docker:
	docker-compose down

start-local-docker:
	docker-compose up --build --remove-orphans

stop-start:
	make stop-local-docker && make start-local-docker
