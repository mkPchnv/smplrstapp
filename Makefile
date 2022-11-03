BINARY_NAME=smplrstapp

build:
	GOARCH=amd64 GOOS=darwin go build -o ./.bin/${BINARY_NAME} ./cmd/smplrstapp/main.go

run:
	./.bin/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ./.bin/${BINARY_NAME}

dep:
	go mod download

swag:
	swag init -g ./cmd/smplrstapp/main.go

lint:
	golangci-lint run --fast