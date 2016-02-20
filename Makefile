all: get-deps test build

get-deps:
	go get ./...

build:
	go build

test:
	go test . ./constraints ./parse ./resources/*/ ./schema

package: all
	cat version.go | sed -n 's/.*"\(.*\).*"/\1/p' | xargs -n1 -I{} tar -czf cfval-{}.tar.gz cfval
