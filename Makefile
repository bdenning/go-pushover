build:
	go build ./cmd/pushover

install: build
	go install ./cmd/pushover

all: install build

.PHONY: install build
