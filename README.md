# GitHub Actions workflow hub

Reusable, versioned CI workflows for Go projects and container images.

## Workflows

- [Go CI](docs/go-ci.md): formatting, vet, build, golangci-lint, race-aware tests and coverage, and govulncheck.
- [Container build and scan](docs/container-build-scan.md): build a local Docker image, scan it with Trivy, and optionally export the exact scanned image for a caller-owned smoke test.
- [Getting started and private-repository access](docs/getting-started.md).
- [Versioning and pinning policy](docs/versioning.md).

Reusable workflow files live directly under `.github/workflows/`, as required by GitHub Actions. Caller repositories retain triggers, permissions, concurrency, and application-specific smoke tests.

## Quick start

1. Confirm the hub repository is accessible to the caller. This hub is public, so private callers can use it without configuring private-workflow sharing.
2. Copy an example from [`examples/`](examples/) into the caller's `.github/workflows/`.
3. Replace `<FULL_COMMIT_SHA>` with the full, reviewed hub commit SHA. Do not use a branch name or placeholder in a production caller.
4. Keep `permissions: contents: read` unless the caller has a separately justified need.
5. Run the workflow and configure branch protection/rulesets using the actual check names reported by GitHub.

The example caller names are templates. `examples/go-container-ci.yml` intentionally leaves the smoke-test command to each project.

## Security and support boundaries

- CI only: no image pushes, releases, deployment, cloud credentials, or registry login.
- No secrets are required by either workflow.
- Caching is disabled initially.
- One Go module is checked per Go workflow invocation.
- Private Go module and private base-image authentication are not configured by default.
- GitHub-hosted Ubuntu runners, outbound network access, and public Go vulnerability/Trivy feeds are expected.
- Workflow and tool updates are reviewed and released from this hub; callers pin an immutable commit.
