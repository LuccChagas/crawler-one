run:
	go run main.go --url=https://pt.wikipedia.org/wiki/Wikipédia -s=$HOME/dummyfolder

test:
	go test -v -cover -timeout 0 ./...

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/crawler-one main.go

.PHONY: debug
