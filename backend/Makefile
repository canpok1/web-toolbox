BUILD_OUTPUT_DIR=build/release
SERVER_MAIN_FILE=cmd/server/main.go
SERVER_BINARY=server

.PHONY: setup
setup:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: run
run:
	go run ${SERVER_MAIN_FILE}

.PHONY: build
build:
	go build -o ${BUILD_OUTPUT_DIR}/${SERVER_BINARY} ${SERVER_MAIN_FILE}

.PHONY: clean
clean:
	go clean
	rm -f ${BUILD_OUTPUT_DIR}/${SERVER_BINARY}

.PHONY: generate-code
generate-code:
	go generate ./...
