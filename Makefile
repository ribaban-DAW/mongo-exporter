TEST_DIR = ./api/hello ./api/metric

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
