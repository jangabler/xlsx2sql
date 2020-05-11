.PHONY: clean test coverage

%: ./cmd/%/main.go
	@go build -o $@ $<

clean:
	@rm -f ./xlsx2sql ./xlsxgen ./test/data/*.xlsx

test:
	@go test -covermode=count -coverprofile=count.out ./...

coverage: test
	@go tool cover -html=count.out
