.PHONY: default
default:
	echo "do nothing"

.PHONY: run
run:
	go run ./cmd/url-shortener/main.go

.PHONY: docker-build
docker-build:
	docker build -t shinunandakumar/url-shortener .

.PHONY: docker-run
docker-run:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down
