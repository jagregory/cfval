all: get-deps test build

get-deps:
	go get ./...

build:
	go build

test:
	go test . ./constraints ./resources/*/ ./schema
