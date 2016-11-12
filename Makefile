.PHONY: lint test vendor_clean vendor_get vendor_update

GOPATH := ${PWD}/vendor
export GOPATH

projname = fastestroute

default: vendor_get lint test build
	
build: vet
	gb build all

lint:
	golint ./src/${projname} > lint.txt

test:
	go test ./src/${projname} -cover -v -coverprofile=cover.out > test.out

clean: vendor_clean
	rm -rf pkg/*
	rm -rf bin/*
	rm lint.txt
	rm test.out
	rm cover.out

vendor_clean:
	rm -dRf ./vendor/*

vendor_get: vendor_clean
	gb vendor fetch \
	github.com/BurntSushi/toml

vendor_update: vendor_get
	rm -rf `find ./vendor/src -type d -name .git` \
	&& rm -rf `find ./vendor/src -type d -name .hg` \
	&& rm -rf `find ./vendor/src -type d -name .bzr` \
	&& rm -rf `find ./vendor/src -type d -name .svn`

vet:
	go vet ./src/...

all: clean default

# go2xunit -fail -input test.out -output test.xml
# gocov convert cover.out | gocov-xml > coverage.xml
# rm -rf tmp
# rm tests.xml
# rm coverage.xml