clean:
	[ ! -e $(CURDIR)/build/bin ] || rm -R $(CURDIR)/build/bin
	docker rmi -f tinder/matching

build: clean
	docker build -t tinder/matching -f $(CURDIR)/build/Dockerfile .

.PHONY: build