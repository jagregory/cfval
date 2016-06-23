all: get-deps test build test_templates

get-deps:
	go get ./...

build:
	go build

test:
	go test . ./constraints ./parse ./resources/*/ ./schema

test_templates:
	find ./spec_templates -name \*.json -exec sh -c "echo; echo Validating: {}; ./cfval validate {}" \;

package: all
	cat version.go | sed -n 's/.*"\(.*\).*"/\1/p' | xargs -n1 -I{} tar -czf cfval-{}.tar.gz cfval

fetch_sample_templates:
	s3cmd sync s3://cloudformation-templates-ap-southeast-2 ./specs/sample_templates

test_sample_templates:
	find ./specs/sample_templates -exec sh -c "echo; echo Validating: {}; ./cfval validate -format=machine {}" \; -name \*.template -or -name \*.json

update_docs_last_modified:
	@grep --include \*.go -RPohe '(?<=// see: )(.*)$$' . | \
	cut -d'#' -f1 | \
	sort | \
	uniq | \
	xargs -P50 -n1 -I{} sh -c "export url={}; curl -s -I -X HEAD \$${url} | grep Last-Modified | xargs echo \$${url}" | \
	sort > docs/aws_modifications.txt
