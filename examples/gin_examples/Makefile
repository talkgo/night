run:
	@echo "=============starting server============="
	go run cmd/server/main.go

test:
	@echo "=============running test============="
	go test ./...

docker-build:
	@echo "=============building image============="
	docker build . -t ginexamples-backend:`git rev-parse HEAD`

compose-up:
	@echo "=============starting gollery locally============="
	docker-compose -f docker-compose-dev.yml up

compose-logs:
	docker-compose logs -f

compose-down:
	docker-compose down

pg-up:
	@echo "=============running a temporary postgres============="
	docker run --rm --name pg-docker -e POSTGRES_USER=postgres -e POSTGRES_DB=ginexamples -d -p 5432:5432 postgres

pg-down:
	@echo "=============stopping the temporary postgres============="
	docker stop pg-docker

docker-reset:
	@echo "=============resetting docker============="
	docker system prune -af
