build:
	go build ./...
test:
	go test -cover ./...
lint:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
	${GOPATH}/bin/golangci-lint run --deadline=3m --timeout=3m ./... # Run linters
