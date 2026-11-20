# Reusable Go CI

Workflow: `.github/workflows/go-ci.yml`

The workflow checks one Go module on `ubuntu-24.04` using three parallel jobs:

1. **Code Quality** — gofmt validation, `go vet ./...`, `go build ./...`, and golangci-lint.
2. **Unit Tests and Coverage** — `go test -count=1`, race detection by default, atomic coverage over `./...`, coverage summary, and a seven-day artifact if a profile was produced.
3. **Dependency Security** — a version-pinned govulncheck source scan.

No checks need secrets or write permissions. Caches are disabled.

## Inputs

| Input | Type | Default | Notes |
| --- | --- | --- | --- |
| `working-directory` | string | `.` | Module directory relative to repository root. |
| `go-version-file` | string | `go.mod` | `go.mod` or `go.work` relative to `working-directory`. |
| `race` | boolean | `true` | Disable only if the module cannot be built/tested with the race detector. |
| `test-timeout` | string | `3m` | Go test timeout per package test binary; separate from the 15-minute job timeout. |
| `golangci-lint-version` | string | `v2.15.0` | Exact version passed to the pinned golangci-lint action. |
| `coverage-artifact-name` | string | `go-coverage` | Use a distinct name if the caller invokes this workflow more than once. |

No outputs or secrets are defined.

## Caller example

```yaml
jobs:
  go:
    uses: Delta-a-Sierra/workflows/.github/workflows/go-ci.yml@<FULL_COMMIT_SHA>
    with:
      working-directory: services/api
      go-version-file: go.mod
      race: true
      test-timeout: 3m
```

Replace the placeholder SHA with a reviewed, full 40-character commit SHA before use.

## Behavior and limitations

- The workflow checks the caller's code and uses the caller's golangci-lint config, if present.
- Formatting runs `go fmt ./...` and fails if it changes tracked Go files.
- The coverage profile is named `coverage.out` in the module directory and uploaded when present, including after a test failure.
- No coverage threshold is enforced.
- Dependency security runs `go run golang.org/x/vuln/cmd/govulncheck@v1.1.4 ./...`. A module may start failing when new vulnerability data is published even if its source is unchanged; vulnerability intelligence is live.
- The Go setup action reads the requested version file. The caller must use a version supported by the selected runner/toolchain.
- One invocation covers one module. For a repository with multiple modules, call once per module and provide unique artifact names.
- Private Go dependencies are not authenticated by this workflow.
