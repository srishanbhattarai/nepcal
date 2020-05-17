test:
	- go test -v ./...

cross:
	- go build -o bin/cross cmd/cross/main.go
	- ./bin/cross cmd/cross/reference.json

cover:
	- go test -v -covermode=count -coverprofile=coverage.out ./...

