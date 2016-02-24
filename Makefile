all: get-deps test build test_templates

get-deps:
	go get ./...

build:
	go build

test:
	go test . ./constraints ./parse ./resources/*/ ./schema

test_templates:
	find ./spec_templates -name \*.json | xargs -I{} -n1 sh -c "echo; echo Validating: {}; ./cfval validate {}"

package: all
	cat version.go | sed -n 's/.*"\(.*\).*"/\1/p' | xargs -n1 -I{} tar -czf cfval-{}.tar.gz cfval
