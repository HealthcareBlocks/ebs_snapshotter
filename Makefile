.DEFAULT_GOAL := osx

osx:
	go build -o bin/ebs_snapshotter

linux:
	docker run --rm -it -v $(PWD):/src healthcareblocks/gobuild -o linux

deps:
	godep restore

docker:
	docker build -t healthcareblocks/ebs_snapshotter .
	@docker images -f "dangling=true" -q | xargs docker rmi

push_to_docker: tag_version
	version=$(shell docker run --rm healthcareblocks/ebs_snapshotter -v); \
	docker push healthcareblocks/ebs_snapshotter:latest; \
	docker push healthcareblocks/ebs_snapshotter:$$version;

tag_version:
	version=$(shell docker run --rm healthcareblocks/ebs_snapshotter -v); \
	docker tag healthcareblocks/ebs_snapshotter healthcareblocks/ebs_snapshotter:$$version

.PHONY: build build_all clean docker push_to_docker tag_version
