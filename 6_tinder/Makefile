build:
	go mod download
	go build -o $(CURDIR)/build/bin/matching $(CURDIR)/cmd/matching
	docker build -t tinder/matching -f $(CURDIR)/build/Dockerfile .

.PHONY: build