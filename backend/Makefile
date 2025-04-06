BUILD_OUTPUT_DIR=build/release
SERVER_MAIN_FILE=cmd/server/main.go
SERVER_BINARY=server
ENV_FILE=.env

.PHONY: setup
setup:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install go.uber.org/mock/mockgen@latest

.PHONY: run
run:
	export $$(cat ${ENV_FILE} | xargs) \
	&& go run ${SERVER_MAIN_FILE}

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o ${BUILD_OUTPUT_DIR}/${SERVER_BINARY} ${SERVER_MAIN_FILE}

.PHONY: clean
clean:
	go clean
	rm -f ${BUILD_OUTPUT_DIR}/${SERVER_BINARY}

.PHONY: generate
generate-code:
	go generate ./...
