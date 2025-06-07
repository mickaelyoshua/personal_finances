COMPOSE = docker compose
EXEC = docker exec -it
DATABASE_URL = postgresql://postgres:postgres@localhost:5432/personal_finance?sslmode=disable

# Run database container
.PHONY: run_db
run_db:
	$(COMPOSE) up postgres -d && \
	$(COMPOSE) logs -f postgres

# Drop database
.PHONY: drop_db
drop_db:
	$(EXEC) postgres psql -U postgres -c "DROP DATABASE IF EXISTS personal_finance;"

.PHONY: create_migrations
create_migrations:
	mkdir app/migrations && \
	migrate create -ext sql -dir app/migrations -seq schema

# Run migrations
.PHONY: migrate_up
migrate_up:
	migrate -path app/migrations -database "$(DATABASE_URL)" -verbose up

# Rollback migrations
.PHONY: migrate_down
migrate_down:
	migrate -path app/migrations -database "$(DATABASE_URL)" -verbose down

# Down containers
.PHONY: down
down:
	$(COMPOSE) down

# Delete volumes
.PHONY: delete_volumes
delete_volumes:
	sudo rm -rf ./postgres_data