GO ?= go
CMD_PKG := ./cmd/reflector
BINARY ?= reflector
BUILD_DIR ?= bin

.PHONY: all build build-x64 clean test docker docker-push

all: build

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY) $(CMD_PKG)

build-x64:
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY)-linux-amd64 $(CMD_PKG)

test:
	$(GO) test ./...

docker:
	docker build --platform=linux/amd64 -t $(BINARY):latest .

clean:
	rm -rf $(BUILD_DIR)
