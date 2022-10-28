BIN_DIR      := bin
SERVER_DIR   := server
CLIENT_DIR   := client
PROTO_DIR    := proto
PROTO_BIN    := $(shell brew --prefix)/bin/protoc
MODULE       := $(shell  head -1 go.mod | awk '{print $$2}')
OUT_DIR      := .
SERVER_BIN    = ${SERVER_DIR}
CLIENT_BIN    = ${CLIENT_DIR}

CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)

.DEFAULT_GOAL := help
.PHONY: greet calculator
project := greet calculator

all: $(project) ## Generate Pbs and build
greet: $@ ## Generate Pbs and build for greet
calculator: $@ ## Generate Pbs and build for greet

$(project):
	@${CHECK_DIR_CMD}
	@${PROTO_BIN} -I$@/${PROTO_DIR} \
	--go_opt=module=${MODULE} \
	--go_out=${OUT_DIR} \
	--go-grpc_opt=module=${MODULE} \
	--go-grpc_out=. \
	$@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

test: all ## Launch tests
	go test ./...

clean: clean_greet ## Clean generated files
	@rm -f ssl/*.crt
	@rm -f ssl/*.csr
	@rm -f ssl/*.key
	@rm -f ssl/*.pem
	@rm -rf ${BIN_DIR}

clean_greet: ## clean the greet project
	@rm -rf greet/${PROTO_DIR}/*.pb.go

rebuild: clean all ## Rebuild the whole project

bump: all ## Update packages version
	go get -u ./...

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: about
	@grep -E '^[a-zA-Z_-]+:.*?\#\#.*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'