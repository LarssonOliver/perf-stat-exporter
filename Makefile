BIN_DIR := bin
BIN_NAME := perf-stat-exporter

SOURCE_FILES := go.mod go.sum $(shell find cmd pkg -name '*.go')

all: $(BIN_DIR)/$(BIN_NAME)

$(BIN_DIR)/$(BIN_NAME): $(SOURCE_FILES)
	go build -o $@ cmd/perf-stat-exporter.go

.PHONY: clean

clean:
	rm -rf $(BIN_DIR)/$(BIN_NAME)
