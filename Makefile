COMPOSE = docker compose

# Run only db
.PHONY: run_db
run_db:
	$(COMPOSE) up postgres

# Run pgadmin
.PHONY: run_pgadmin
run_pgadmin:
	$(COMPOSE) up pgadmin -d && \
	$(COMPOSE) logs -f postgres

# Down containers
.PHONY: down
down:
	$(COMPOSE) down

# Delete volumes
.PHONY: delete_volumes
delete_volumes:
	sudo rm -rf ./postgres_data