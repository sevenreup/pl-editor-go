BINARY_NAME=pl-editor
GO=go

build:
	go build -o bin/${BINARY_NAME} src/main.go

run:
	go run src/main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	$(GO) test ./src/...