TEST_DIR = ./internal/hello ./internal/metric

deploy:
	clear
	docker compose down
	docker compose up --build -d

test:
	clear
	go test -v $(TEST_DIR)
