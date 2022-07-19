make:
	go run main.go

test:
	go test ./... -short

build:
	go build -o dc main.go

install:
	go install
