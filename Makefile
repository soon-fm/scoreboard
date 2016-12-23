BUILD_DIR				?= ./.build
BIN_NAME				?= scoreboard
BUILD_TIME				?= $(shell date +%s)
VERSION					?= $(shell git rev-parse HEAD)
DOCKER_IMAGE			?= containers.soon.build/sfm/scoreboard
DOCKER_TAG				?= latest
VERSION_BUILD_FLAG		?= -X scoreboard/version.buildTime=${BUILD_TIME}
BUILDTIME_BUILD_FLAG 	?= -X scoreboard/version.version=${VERSION}
BUILD_FLAGS				?= -ldflags "-s ${VERSION_BUILD_FLAG} ${BUILDTIME_BUILD_FLAG}"

build:
	go build ${BUILD_FLAGS}

all: linux darwin image

#
# Linux
#

linux: linux64
linux64:
	mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/linux64_$(BIN_NAME) ${BUILD_FLAGS}

#
# Darwin
#

darwin: darwin64
darwin64:
	mkdir -p ${BUILD_DIR}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BUILD_DIR}/osx64_$(BIN_NAME) ${BUILD_FLAGS}

#
# Docker Image
#

image:
	docker build --build-arg BUILD_DIR=$(BUILD_DIR) --build-arg BIN_NAME=linux64_$(BIN_NAME) --force-rm -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
