test:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

tidy:
	go mod tidy