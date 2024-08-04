APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=ghcr.io/thanxx
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=amd64
TAG=${REGISTRY}/${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

image:
	docker build . -t ${TAG}

push:
	docker push ${TAG}

build: format
	go get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="kbot/cmd.appVersion=${VERSION}

clean:
	rm -rf kbot