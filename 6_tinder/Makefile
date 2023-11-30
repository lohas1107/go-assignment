clean:
	[ ! -e $(CURDIR)/build/bin ] || rm -R $(CURDIR)/build/bin
	docker rmi -f tinder/matching

build: clean
	docker build -t tinder/matching -f $(CURDIR)/build/Dockerfile .

run: build
	docker run -p 8080:8080 tinder/matching

e2e:
	go test $(CURDIR)/test/e2e

.PHONY: build