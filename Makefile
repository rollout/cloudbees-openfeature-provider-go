build:
	go build ./...
test:
	go test -cover ./...
test-junit:
	go install github.com/jstemmer/go-junit-report@v1.0.0
	go test -cover -v ./... | go-junit-report -set-exit-code >junit.xml
lint:
	go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3
	${GOPATH}/bin/golangci-lint run --deadline=3m --timeout=3m ./... # Run linters
