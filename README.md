# Go Garbage Collection Example

## Local Development

Install [air](https://github.com/cosmtrek/air) for live reloading:

```sh
go install github.com/cosmtrek/air@latest
```

Use the "dev" command in the Makefile:

```sh
make dev
```

### Generate Docs

Use the "docs" command in the Makefile to generate reference documentation for each package as a README.md file:

```sh
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
make docs
```

### Commit Hooks

Set up Git hooks to do things like generate docs automatically.

```sh
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
make hooks
```

### Linting

Lint with [golangci-lint](https://golangci-lint.run/):

```sh
make lint

# Run lint with automatic fixes
make lint/fix
```

### Testing

To run unit tests:

```sh
make test
```

To test individual packagesm or an individual test within a package:

```sh
go test ./internal/handlers -v
go test ./internal/handlers -v -run="TestMemoryHandlers"
```

For packages that use testify suites, to run an individual test in the suite:

```sh
go test ./internal/handlers -v
go test ./internal/handlers -v -run="TestMemoryHandlers" -testify.m "TestAllocateSuccess"
```
