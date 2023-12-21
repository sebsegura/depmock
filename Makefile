BIN= $(CURDIR)/bin
FUNCTION=credit

.PHONY: build

build: fmt
	@env GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap cmd/main.go

lint:
	@go vet ./...

fmt:
	@go fmt ./...

clean:
	@rm -rf $(BIN)

zip: build
	@zip -j $(FUNCTION).zip bin/bootstrap

.PHONY: test
test:
	@go test -v -cover ./...