##run gRPC server
.PHONY: run
run: fmt
	go run ./src/main.go

##format project
.PHONY: fmt
fmt:
	go fmt ./...

##build project
.PHONY: build
build: fmt
	go build ./src/main.go 

##install all dependencies
.PHONY: install
install:
	go get -d ./...

##remove unused dependencies
.PHONY: optimize
optimize:
	go mod tidy -v

.DEFAULT_GOAL := run