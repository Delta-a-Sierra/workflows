# Versioning and maintenance

Reusable workflows are executable dependencies. A change can affect every caller, so consumer pins and release changes are reviewed deliberately.

## Consumer pinning

Pin a caller to the full 40-character commit SHA:

```yaml
uses: Delta-a-Sierra/workflows/.github/workflows/go-ci.yml@0123456789abcdef0123456789abcdef01234567 # v1.0.0
```

The SHA above is illustrative only. Use the actual SHA of a reviewed release commit. Avoid `main`, floating tags, and invented placeholder SHAs in active workflows.

## Releases

- Publish semantic tags only after hub validation and pilot adoption.
- Treat tags as immutable; use a new version rather than moving a published tag.
- Breaking input/output/artifact/check-name behavior requires a major version and migration notes.
- Compatible improvements can use minor versions; fixes use patch versions.
- Tool or policy updates can expose new lint/vulnerability failures even if the workflow interface is unchanged. State these behavioral changes in release notes.
- A moving `v1` convenience tag is optional, but full-SHA consumer pins remain the production recommendation.

## Dependency maintenance

Dependabot should keep GitHub Actions references current through reviewed pull requests. Explicit tool versions installed by shell commands (for example govulncheck) and action inputs (for example Trivy and golangci-lint versions) need deliberate version updates; do not assume the Actions ecosystem updater maintains all such versions.

Validate action references, tool compatibility with supported Go versions, and scanner behavior before merging updates. Do not enable auto-merge for security or policy changes by default.

## Rollback

If a hub revision breaks a consumer, pin that consumer back to the previous known-good full SHA or revert its migration. Do not disable genuine vulnerability findings as a rollback shortcut. Retain prior release commits and migration notes.
