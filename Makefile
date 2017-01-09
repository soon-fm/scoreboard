BIN_DIR					?= ./.bin
BIN_NAME				?= scoreboard
BUILD_TIME				?= $(shell date +%s)
VERSION					?= $(shell git rev-parse HEAD)
DOCKER_IMAGE			?= quay.io/thisissoon/fm-scoreboard
DOCKER_TAG				?= latest
VERSION_BUILD_FLAG		?= -X scoreboard/version.buildTime=${BUILD_TIME}
BUILDTIME_BUILD_FLAG 	?= -X scoreboard/version.version=${VERSION}
BUILD_FLAGS				?= -ldflags "-s ${VERSION_BUILD_FLAG} ${BUILDTIME_BUILD_FLAG}"

all: bin_dir linux darwin

build:
	go build ${BUILD_FLAGS}

bin_dir:
	mkdir -p ${BIN_DIR}

#
# Linux
#

linux: linux64
linux64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/linux64_$(BIN_NAME) ${BUILD_FLAGS}

#
# Darwin
#

darwin: darwin64
darwin64:
	mkdir -p ${BIN_DIR}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/osx64_$(BIN_NAME) ${BUILD_FLAGS}

#
# Docker Image
#

image: all
	docker build --build-arg BIN_DIR=$(BIN_DIR) --build-arg BIN_NAME=linux64_$(BIN_NAME) --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
