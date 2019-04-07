.DEFAULT_GOAL = all

version  := $(shell git rev-list --count HEAD).$(shell git rev-parse --short HEAD)

name     := wscat
package  := github.com/corpix/$(name)

build       := ./build
build_id    := 0x$(shell echo $(version) | sha1sum | awk '{print $$1}')
ldflags     := -X $(package)/cli.version=$(version) -B $(build_id)
build_flags := -a -ldflags "$(ldflags)" -o build/$(name)

.PHONY: all
all: build

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -bench=. -v ./...

.PHONY: lint
lint:
	go vet -v ./...

.PHONY: check
check: lint test

.PHONY: $(name)
$(name):
	mkdir -p $(build)
	@echo "Build id: $(build_id)"
	go build $(build_flags) -v $(package)/$(name)

.PHONY: build
build: $(name)

.PHONY: clean
clean:
	git clean -xddff
