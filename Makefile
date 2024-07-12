VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build: format
	go get
	CGO_ENABLED=0 GOOS=${shell dpkg --print-architecture} go build -v -o kbot -ldflags "-X="kbot/cmd.appVersion=${VERSION}

clean:
	rm -rf kbot