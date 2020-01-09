#.PHONY all build push login

VERSION?=v0.6.0
ARTIFACTORY_PATH?=development/delivery
ARTIFACTORY_APIKEY?=<your API Key>

all: build release

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-linux-x86_64 tfs.bb.delivery/tool-tmplt
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-windows-x86_64.exe tfs.bb.delivery/tool-tmplt
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o bin/tmplt-$(VERSION)-darwin-x86_64 tfs.bb.delivery/tool-tmplt

install-linux:
	sudo cp bin/tmplt-$(VERSION)-linux-x86_64 /usr/local/bin/
	sudo ln -sfr /usr/local/bin/tmplt-$(VERSION)-linux-x86_64 /usr/local/bin/tmplt

release:
	curl -ki -H "X-JFrog-Art-Api: $(ARTIFACTORY_APIKEY)" -T bin/tmplt-$(VERSION)-linux-x86_64 "https://artifactory.binckbank.nv/artifactory/$(ARTIFACTORY_PATH)/"
	curl -ki -H "X-JFrog-Art-Api: $(ARTIFACTORY_APIKEY)" -T bin/tmplt-$(VERSION)-windows-x86_64.exe "https://artifactory.binckbank.nv/artifactory/$(ARTIFACTORY_PATH)/"
	curl -ki -H "X-JFrog-Art-Api: $(ARTIFACTORY_APIKEY)" -T bin/tmplt-$(VERSION)-darwin-x86_64 "https://artifactory.binckbank.nv/artifactory/$(ARTIFACTORY_PATH)/"
	echo
	git tag $(VERSION)
	echo "remember to git push --tags"