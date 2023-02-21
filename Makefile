VERSION := v0.1.1
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build:
	go build ./cmd/tmc

.PHONY: package
package: clean build
	zip tmc_$(VERSION)_$(GOOS)_$(GOARCH).zip $(shell ls tmc tmc.exe)

.PHONY: clean
clean:
	rm -f tmc tmc.exe
