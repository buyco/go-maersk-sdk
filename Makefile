DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v6.0.0 generate -i sdk/api_files/maersk_track_and_trace.yaml
GO_CLIENT := -g go -o /sdk/${API_TYPE}/maersk \
			--git-repo-id=go-maersk-sdk --git-user-id=buyco \
			--additional-properties=packageName=maersk \
			--additional-properties=isGoSubmodule=true \
			--additional-properties=generateInterfaces=true

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./api

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
