.PHONY: setup
setup:
	make -C frontend setup
	make -C backend setup

.PHONY: setup-e2e
setup-e2e:
	make -C e2e setup

.PHONY: run-front
run-front:
	make -C frontend run

.PHONY: run-back
run-back:
	make -C backend run

.PHONY: test
test:
	make -C backend test

.PHONY: test-e2e
test-e2e: build-front
	make -C e2e test

.PHONY: test-e2e-debug
test-e2e-debug: build-front
	make -C e2e test-debug

.PHONY: show-e2e-report
show-e2e-report:
	make -C e2e show-report

.PHONY: check
check:
	make -C frontend check
	make -C e2e check

.PHONY: build
build:
	make -C frontend build
	make -C backend build

.PHONY: build-front
build-front:
	make -C frontend build

.PHONY: build-back
build-back:
	make -C backend build

.PHONY: format
format:
	make -C frontend format
	make -C e2e format

.PHONY: generate
generate:
	make -C frontend generate
	make -C backend generate
