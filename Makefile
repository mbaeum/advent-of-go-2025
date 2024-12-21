
test:
	mkdir -p coverage
	go test -v -coverprofile=coverage/coverage.out ./...

build:
	go build -v ./...
