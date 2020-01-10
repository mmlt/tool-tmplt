#.PHONY all build push install-linux

VERSION?=v0.6.0

all: build 

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-linux-amd64 github.com/mmlt/tool-tmplt
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-windows-amd64.exe github.com/mmlt/tool-tmplt
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-darwin-amd64 github.com/mmlt/tool-tmplt

install-linux:
	sudo cp bin/tmplt-$(VERSION)-linux-amd64 /usr/local/bin/
	sudo ln -sfr /usr/local/bin/tmplt-$(VERSION)-linux-amd64 /usr/local/bin/tmplt

