BUILDDIR ?= $(CURDIR)/build
DOCSDIR ?= $(CURDIR)/docs
BINARY_NAME = initia-apis

all: clean docs build run

build:
	@go build -o build/${BINARY_NAME} main.go

run:
	@./build/${BINARY_NAME}

test:
	@go test ./... -v 

docs:
	@swag init -g main.go

clean:
	rm -rf \
    $(BUILDDIR)/ \
	$(DOCSDIR)
	
 
.PHONY: all test build run clean 


#### docker ####

docker-build: clean docs
	@docker rmi -f initia-apis
	@docker build --no-cache --tag initia-apis ./

docker-run:
	@docker run -d --rm -p 8999:8999 initia-apis  

docker:
	@make docker-build
	@make docker-run

.PHONY: docker-build docker-run
