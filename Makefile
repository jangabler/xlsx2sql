.PHONY: test coverage

xlsx2sql: ./cmd/xlsx2sql/main.go
	@go build -o $@ $<

xlsxgen: ./cmd/xlsxgen/main.go
	@go build -o $@ $<

test:
	@go test -covermode=count -coverprofile=count.out ./...

coverage: test
	@go tool cover -html=count.out
