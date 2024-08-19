build:
@go build -o bin/ecommerce_backend cmd/main.go


test:
@go test -v ./..

run: build
@./bin/ecommerce_backend