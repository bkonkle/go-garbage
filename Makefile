.PHONY: fmt
fmt:
	gofmt -w -s ./

.PHONY: dev
dev:
	air

.PHONY: hooks
hooks:
	cp hooks/pre-commit .git/hooks/pre-commit

.PHONY: docs
docs:
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest && \
	gomarkdoc ./cmd/... ./internal/...

# Linting
# -------

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: lint/fix
lint/fix:
	golangci-lint run --fix

# Testing
# -------

.PHONY: test
test:
	go test -v -race ./...

.PHONY: test/ci
test/ci:
	go test -v -race ./... -coverprofile c.out

.PHONY: test/clean
test/clean:
	go clean -testcache

# Go Build
# --------

define go_build
    export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/$(1) ./cmd/$(1)/main.go
endef

.PHONY: build/clean
build/clean:
	if [ -d bin ]; then rm -r bin; fi
	mkdir -p bin

.PHONY: build/example
build/example:
	$(call go_build,example)
