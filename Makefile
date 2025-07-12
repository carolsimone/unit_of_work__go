INFRA=docker-compose --project-name carnival -f docker-compose.yml
TEST_CONTAINER_NAME=carnival-test

clean:
	docker system prune -f
	docker volume prune -f

build:
	$(INFRA) build --no-cache

up:
	$(INFRA) up -d

ci: build up

test:
	docker exec -i carnival bash -c 'cd src && go test -v -coverprofile=coverage.out $$(go list ./...) && go tool cover -func=coverage.out'

down:
	$(INFRA) down -v --remove-orphans
