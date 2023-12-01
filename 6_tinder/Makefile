clean:
	[ ! -e $(CURDIR)/build/bin ] || rm -R $(CURDIR)/build/bin
	docker rmi -f tinder/matching
	docker stop matching || true
	docker rm matching || true

build: clean
	docker build -t tinder/matching -f $(CURDIR)/build/Dockerfile .

run: build
	docker run -d --name matching -p 8080:8080 tinder/matching

e2e: run
	$(GOROOT)/bin/go test $(CURDIR)/test/e2e

.PHONY: build