.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux env SLS_DEBUG=* go build -ldflags="-s -w" -o bin/sum sum/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
