PROJECT?=github.com/reddaemon/apiproject
APP?=apiproject
PORT?=8080

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=reddaemon/${APP}

GOOS?=linux
GOARCH?=amd64

precommit:
	gofmt -w -s -d .
	go vet .
	golangci-lint run --enable-all
	go mod tidy
	go mod verify

test:
	go test -race -cover ./handlers/...
	go test -cover ./handlers/...

clean:
	rm -f ${APP}

build: test clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
        -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
        -o ${APP}

up: build
	docker-compose -f docker/docker-compose/docker-compose.yml up
down:
	docker-compose -f docker/docker-compose/docker-compose.yml down
up-build: build
	docker-compose -f docker/docker-compose/docker-compose.yml up --build -d

restart: down up




