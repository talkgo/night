up:
	@echo "=============starting gollery locally============="
	docker-compose -f docker-compose-dev.yml up

logs:
	docker-compose logs -f

down:
	docker-compose down

pg-up:
	@echo "=============running a temporary postgres============="
	docker run --rm --name pg-docker -e POSTGRES_USER=postgres -e POSTGRES_DB=ginexamples -d -p 5432:5432 postgres

pg-down:
	@echo "=============stopping the temporary postgres============="
	docker stop pg-docker

docker-rest:
	@echo "=============resetting docker============="
	docker system prune -af
