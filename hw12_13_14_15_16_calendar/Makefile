BIN := "./bin/calendar"
BIN_SCHEDULER := "./bin/calendar_scheduler"
BIN_SENDER := "./bin/calendar_sender"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

build-scheduler:
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/calendar_scheduler

build-sender:
	go build -v -o $(BIN_SENDER) -ldflags "$(LDFLAGS)" ./cmd/calendar_sender

run: build
	$(BIN) -config ./configs

run-scheduler: build-scheduler
	$(BIN_SCHEDULER) -config ./configs

run-sender: build-sender
	$(BIN_SENDER) -config ./configs

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/storage/memory/ ./internal/app/

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.64.8

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build run build-img run-img version test lint

migrate:
	goose up -dir migrations


migrate-down:
	goose down -dir migrations

generate:
	cd api && protoc --go_out=../internal/server/grpc --go_opt=paths=source_relative --go-grpc_out=../internal/server/grpc --go-grpc_opt=paths=source_relative EventService.proto