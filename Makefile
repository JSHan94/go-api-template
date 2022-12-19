BUILDDIR ?= $(CURDIR)/build
DOCSDIR ?= $(CURDIR)/docs
BINARY_NAME = initia-apis

all: docs build run

build:
	@go build -o build/${BINARY_NAME} main.go

run:
	./build/${BINARY_NAME}

test:
	@go test ./...

docs:
	@swag init -g main.go

clean:
	rm -rf \
    $(BUILDDIR)/ \
	$(DOCSDIR)
	
 
.PHONY: all build run test clean 