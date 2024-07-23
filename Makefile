.PHONY: help
help:
	@echo "list of available commands:"
	@echo "- make -> same with 'make help'"
	@echo "- make help -> show this help"
	@echo "- make run -> run the project"
	@echo "- make build -> build the project"
	@echo "- make test -> run unit test"

.PHONY: run
run:
	go run .

.PHONY: build
build:
	go build .

.PHONY: test
test:
	go test ./... -v
