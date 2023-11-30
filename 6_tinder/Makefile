clean:
	[ ! -e $(CURDIR)/build/bin ] || rm -R $(CURDIR)/build/bin
	docker rmi -f tinder/matching

build: clean
	go mod download
	go build -o $(CURDIR)/build/bin/matching $(CURDIR)/cmd/matching
	docker build -t tinder/matching -f $(CURDIR)/build/Dockerfile .

.PHONY: build