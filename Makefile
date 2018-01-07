GOFILES = $(shell find . -name '*.go')
GOPACKAGES = $(shell go list ./...)

default: build

dist:
	mkdir -p dist
	echo "creating dir"

build: dist/props

build-native: $(GOFILES)
	go build -o workdir/native-props .

dist/props: $(GOFILES)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o dist/props .

test: test-all

test-all:
	@go test -v $(GOPACKAGES)