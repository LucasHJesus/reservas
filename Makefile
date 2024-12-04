build:
	@go build -o bin/reservas cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/reservas