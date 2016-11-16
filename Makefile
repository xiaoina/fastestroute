.PHONY: default build lint test clean_build clean_all vendor_clean vendor_get vendor_update vet all

GOPATH := ${PWD}/vendor
export GOPATH

projname = fastestroute

default: lint test build
	
build: vet
	gb build ${projname}

lint:
	golint ./src/${projname} > lint.txt

test:
	go test ./src/tests -cover -v -coverprofile=cover.out > test.out

clean_build:
	rm -rf pkg/*
	rm -rf bin/*
	rm lint.txt
	rm test.out
	rm cover.out

clean_all: vendor_clean
	rm -rf pkg/*
	rm -rf bin/*
	rm lint.txt
	rm test.out
	rm cover.out

vendor_clean:
	rm -dRf ./vendor/*

vendor_get: vendor_clean
	gb vendor fetch github.com/BurntSushi/toml
	gb vendor fetch github.com/go-kit/kit

vendor_update: vendor_get
	rm -rf `find ./vendor/src -type d -name .git` \
	&& rm -rf `find ./vendor/src -type d -name .hg` \
	&& rm -rf `find ./vendor/src -type d -name .bzr` \
	&& rm -rf `find ./vendor/src -type d -name .svn`

vet:
	go vet ./src/...

all: clean vendor_update default

# go2xunit -fail -input test.out -output test.xml
# gocov convert cover.out | gocov-xml > coverage.xml
# rm -rf tmp
# rm tests.xml
# rm coverage.xml