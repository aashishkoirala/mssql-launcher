SHELL := /bin/bash

all: generateDocs

.PHONY: clean
clean:
	@if [ -d bin ]; then rm -fR bin; fi
	@mkdir bin
	@mkdir bin/linux
	@mkdir bin/windows

.PHONY: buildlinux
buildlinux: clean
	@cp -f config.yaml bin/linux/mssql-launcher-config.yaml
	@go build -o bin/linux/mssql-launcher

.PHONY: buildwindows
buildwindows: clean
	@cp -f config.yaml bin/windows/mssql-launcher-config.yaml
	@GOOS=windows go build -o bin/windows/mssql-launcher.exe

.PHONY: runTests
runTests: buildlinux buildwindows
	@go test ./...

.PHONY: generateDocs
generateDocs: runTests
	@go doc -all > bin/doc
	@for d in config command logger menu; do go doc -all $$d >> bin/doc; done