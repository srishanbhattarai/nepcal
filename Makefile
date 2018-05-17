build:
	- go build -o bin/nepcal .

test:
	- go test -v ./...
