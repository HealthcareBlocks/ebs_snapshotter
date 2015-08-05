.DEFAULT_GOAL := release

release:
	@docker run --rm -v "$$PWD":/src centurylink/golang-builder-cross
	@mkdir -p bin && mv ebs_snapshotter* bin/

.PHONY: release
