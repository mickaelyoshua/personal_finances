COMPOSE = docker compose

# Run only db
.PHONY: run_db
run_db:
	$(COMPOSE) run --rm postgres