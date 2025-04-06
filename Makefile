.PHONY: setup
setup:
	make -C frontend setup
	make -C backend setup

.PHONY: run
run:
	make -C frontend build
	make -C backend run

.PHONY: test
test:
	make -C backend test

.PHONY: check
check:
	make -C frontend check

.PHONY: build
build:
	make -C frontend build
	make -C backend build

.PHONY: format
format:
	make -C frontend format

.PHONY: generate
generate:
	make -C frontend generate
	make -C backend generate
