.DEFAULT_GOAL := build

build: clean
	docker run --rm -it -v $(PWD):/src healthcareblocks/gobuild -o darwin -a amd64

build_all: clean
	docker run --rm -it -v $(PWD):/src healthcareblocks/gobuild

clean:
	rm -fr ./bin/*

.PHONY: build clean
