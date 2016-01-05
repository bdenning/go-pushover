build:
	go build ./cmd/pushover

install:
	go install ./cmd/pushover

clean:
	rm -f pushover
	rm -f cmd/pushover/pushover

all: install build

.PHONY: install build clean
