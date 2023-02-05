build_path := build/bin
binary := epr

binfile := $(build_path)/$(binary)
repo := "github.com/adanilovich/epr"

GO_PACKAGES    := $(shell $(BUILDER) go list ./... | grep -v "vendor/")

CURDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build
build:
	@go build -race -gcflags='-N -l' -o "$(binfile)" "$(repo)/cmd/$(binary)" 2>&1

deploy: build
	scp "$(binfile)" main:~/


run: build
	@cat urls.txt | build/bin/epr

debug: build
	@echo "[!] Starting debug process"
	@echo -e "[*] Debug target: \e[33m${build_path}/${binary}\e[0m"
	@echo "**********************************************************"
	@dlv exec "${build_path}/${binary}" -- 

debug-test:
	@go clean -testcache
	@go test -tags=unit `go list ./... | grep -i "${package}"` -gcflags='-l -N' -v -c -o "${build_path}/${binary}.test"

	dlv exec "${build_path}/${binary}.test" -- -test.run "${test_name}"

clean-out:
	rm -rf ./out/*
