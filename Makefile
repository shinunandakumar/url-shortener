.PHONY: default
default:
	echo "do nothing"

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: run
run: # To run in your host machine without using go
	go run ./cmd/url-shortener/main.go

.PHONY: test
test: # To run the tests
	go test -v ./...

.PHONY: docker-build
docker-build: # To build the docker image
	docker build -t shinunandakumar/url-shortener .

.PHONY: docker-run
docker-run: # To start the project
	docker compose up -d

.PHONY: docker-down
docker-down: # To stop the project
	docker compose down
