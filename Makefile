.PHONY: up
up:
	docker compose run --rm app
	docker compose down
