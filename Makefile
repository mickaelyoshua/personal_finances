COMPOSE = docker compose

# Run only db
.SILENT:
.PHONY: run_db
run_db:
	$(COMPOSE) run --rm postgres