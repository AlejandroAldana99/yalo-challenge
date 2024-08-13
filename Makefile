run:
	go run main.go

build:
	go build -o bin/main main.go

mocks:
	mockery --all

test:
	go test ./...