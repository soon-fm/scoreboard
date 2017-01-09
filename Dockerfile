# Lightweight Hardended Linux Distro
FROM alpine:3.3

# Update and Install OS level packages
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

# Default build arguments
ARG BIN_DIR=./.build
ARG BIN_NAME=linux64_scoreboard
ARG BUILD_DEST=/usr/local/bin/scoreboard
ARG CONFIGPATH=/etc/scoreboard

# Copy Binary
COPY ${BIN_DIR}/${BIN_NAME} ${BUILD_DEST}

# Volumes
VOLUME ["/etc/scoreboard", "/var/log/scoreboard"]

# Set our Application Entrypoint
ENTRYPOINT ["scoreboard"]
