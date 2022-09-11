.PHONY: checklint
checklint:
ifeq (, $(shell which golangci-lint))
	@echo 'error: golangci-lint is not installed, please exec `brew install golangci-lint`'
	@exit 1
endif

.PHONY: lint
lint: checklint
	golangci-lint run --skip-dirs-use-default


.PHONY: checkbuf
checkbuf:
ifeq (,$(shell which buf))
	@echo 'error:bufbuild/buf is not installed ,please exec `brew install bufbuild/buf/buf`'
	@echo 1
endif

.PHONY: proto
proto: checkbuf
	buf generate

.PHONY: build
build:
	go build cmd/main.go