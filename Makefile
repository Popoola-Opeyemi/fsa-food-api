run:
	go run cmd/main.go

test:
	go test ./...

build:
	go build -o hygiene cmd/main.go

install:
	go mod download
