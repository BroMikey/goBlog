APP_NAME := server
BIN_DIR  := bin
CONFIG   ?= configs/dev.yaml
# allow overriding the actual env the binary reads (default to CONFIG)
APP_CONFIG ?= $(CONFIG)
export APP_CONFIG

.PHONY: build run clean ls help

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

run: build
	$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)

ls:
	@echo "Listing project files (excluding $(BIN_DIR)):"
	@find . -path ./$(BIN_DIR) -prune -o -type f -print

help:
	@printf "Available targets:\n  make build    # build binary into $(BIN_DIR)\n  make run      # build then run (use CONFIG=path or APP_CONFIG=path to override)\n  make clean    # remove $(BIN_DIR)\n  make ls       # list project files\n"
