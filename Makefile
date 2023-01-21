SHALL=/bin/bash

export CGO_ENABLED=0
export DSN=postgres://gouser:gopassword@localhost:4321/gotest?sslmode=disable

default: build
.PHONY: default

build:
	@ echo "-> build binary ..."
	@ go build -ldflags "-X main.HashCommit=`git rev-parse HEAD` -X main.BuildStamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`" -o ./calendar/cmd/server .
.PHONY: build

test:
	@ echo "-> running tests ..."
	@ CGO_ENABLED=1 go test -race ./...
.PHONY: test

lint:
	@ echo "-> running linters ..."
	@ golangci-lint run ./...
.PHONY: lint

migrate:
	@ echo "-> running migration ..."
	@ migrate -path ./migrations -database "postgres://gouser:gopassword@localhost:4321/gotest?sslmode=disable" -verbose up
.PHONY: migrate

migrate-down:
	@ echo "-> running migration ..."
	@ migrate -path ./migrations -database "postgres://gouser:gopassword@localhost:4321/gotest?sslmode=disable" -verbose down
.PHONY: migrate-down
