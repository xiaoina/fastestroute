.PHONY: lint test vendor_clean vendor_get vendor_update

GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH

projname = FastestRoute

default: vendor_update lint test
	# go2xunit -fail -input test.out -output test.xml
	# gocov convert cover.out | gocov-xml > coverage.xml
	gb build all

lint:
	golint ./src/${projname} > lint.txt

test:
	go test ./src/${projname} -cover -v -coverprofile=cover.out > test.out

clean:
	rm -rf pkg/*
	rm -rf bin/*
	rm lint.txt
	rm test.out
	rm cover.out
	# rm -rf tmp
	# rm tests.xml
	# rm coverage.xml

vendor_clean:
	rm -dRf ./vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/vendor go get -d -u -v \
	github.com/kardianos/govendor
	govendor fetch \ 
	github.com/constabulary/gb/...

vendor_update: vendor_get
	rm -rf `find ./vendor/src -type d -name .git` \
	&& rm -rf `find ./vendor/src -type d -name .hg` \
	&& rm -rf `find ./vendor/src -type d -name .bzr` \
	&& rm -rf `find ./vendor/src -type d -name .svn`

commit:
	default
	clean
