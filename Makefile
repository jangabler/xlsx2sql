.PHONY: test
test:
	go test -covermode=count -coverprofile=count.out ./...

.PHONY: coverage
coverage: test
	go tool cover -html=count.out

.PHONY: build
build:
	go build -o ./xlsx2sql ./cmd/xlsx2sql/*

.PHONY: run
run: build
	./xlsx2sql -m ./configs/mapping-template.xml
