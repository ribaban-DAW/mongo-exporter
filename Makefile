TEST_DIR = ./internal/hello ./internal/metric

deploy:
	clear
	docker compose down
	docker compose up --build -d

app:
	docker compose up --build -d app

test:
	clear
	go test -v $(TEST_DIR)

.PHONY: deploy app test
