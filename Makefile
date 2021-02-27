.PHONY: build
build:
	go build -v ./cmd/main

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

#migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/test" -verbose up
#migrate create -ext sql -dir migrations <name_table>