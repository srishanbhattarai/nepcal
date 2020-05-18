test:
	- go test -v ./...

cross: reference.json
	- go build -o bin/cross cmd/cross/main.go
	- ./bin/cross cmd/cross/reference.json

reference.json:
	- curl -o cmd/cross/reference.json https://raw.githubusercontent.com/mesaugat/bikram-sambat-anno-domini-fixtures/master/export-minified.json

cover:
	- go test -v -covermode=count -coverprofile=coverage.out ./...

