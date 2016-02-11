build:
	go build

test:
	go test ./constraints ./resources ./resources/*/ ./schema
