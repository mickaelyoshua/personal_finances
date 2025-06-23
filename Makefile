COMPOSE = docker compose
EXEC = docker exec -it
DATABASE_URL = postgresql://postgres:postgres@localhost:5432/personal_finance?sslmode=disable

# Run database container
.PHONY: run_db
run_db:
	$(COMPOSE) up postgres -d

# Drop database
.PHONY: drop_db
drop_db:
	$(EXEC) postgres psql -U postgres -c "DROP DATABASE IF EXISTS personal_finance;"

# Create migrations directory and files
.PHONY: create_migrations
create_migrations:
	mkdir -p app/db/migrations && \
	migrate create -ext sql -dir app/db/migrations -seq schema

# Run migrations
.PHONY: migrate_up
migrate_up:
	migrate -path app/db/migrations -database "$(DATABASE_URL)" -verbose up

# Rollback migrations
.PHONY: migrate_down
migrate_down:
	migrate -path app/db/migrations -database "$(DATABASE_URL)" -verbose down

# Down containers
.PHONY: down
down:
	$(COMPOSE) down

# Delete volumes
.PHONY: delete_volumes
delete_volumes:
	sudo rm -rf ./postgres_data

# Run tests
.PHONY: test
test:
	cd app && \
	go test ./... -v -cover

# Generate code templ
.PHONY: generate_templ
generate_templ:
	cd app && \
	templ generate

# Generate mock
.PHONY: generate_mock
generate_mock:
	cd app && \
	mockgen -source=db/sqlc/agent.go -destination=db/mock/agent.go -package=mock

# Run application
.PHONY: run_app
run_app:
	cd app && \
	sqlc generate && \
	templ generate && \
	go run .