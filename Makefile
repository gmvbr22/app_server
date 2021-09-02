run:
	go run api/main.go

build:
	go build api/main.go

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out