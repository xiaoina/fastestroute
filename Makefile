export GOPATH=$(pwd)/vendor

projname = FastestRoute

make:
		# golint ./src/cmd/chatbot/... > lint.txt
		# go test ./cmd/... -cover -v -coverprofile=cover.out > test.out
		# go2xunit -fail -input test.out -output test.xml
		# gocov convert cover.out | gocov-xml > coverage.xml
		gb build all

clean:
		rm -rf pkg/*
		rm -rf bin/*
		# rm -rf tmp
		# rm lint.txt
		# rm test.out
		# rm cover.out
		# rm tests.xml
		# rm coverage.xml