# Getting started

## Private consumer repositories

The workflow hub is public, so a private consumer can call it using the normal reusable-workflow syntax. A PAT is not needed just to read and invoke this public workflow.

If the hub is made private in the future, configure GitHub Actions access on the hub under **Settings → Actions → General → Access** and allow the consumer repositories. The owner/account boundary matters: personal repositories and organization repositories have different sharing settings. Public repositories cannot call a private hub. Do not enable access more broadly than needed.

Private workflow access does not authenticate private Go modules, private Docker base images, or private registries. Those require a separate, narrowly scoped design.

## Add Go CI

Copy [`../examples/go-ci.yml`](../examples/go-ci.yml) to the consuming repository's `.github/workflows/`. Pin the `uses` line to the full SHA of a reviewed hub commit. Configure the branch trigger to match that repository's default branch.

The caller should retain ownership of:

- Events (`pull_request`, pushes, manual/scheduled runs).
- Concurrency and cancellation behavior.
- Explicit token permissions.
- The hub commit SHA.
- Branch protection/ruleset settings.

## Add container CI

Copy [`../examples/go-container-ci.yml`](../examples/go-container-ci.yml) and configure its image tag, context, Dockerfile, and artifact names as needed.

The container reusable workflow builds and scans but does not push. If the caller needs a smoke test, set `upload-image: true`, then use a dependent caller-owned job to download the workflow's image artifact, load `image.tar`, and test the resulting image. Do not rebuild the image for the smoke test: test the exact artifact that passed scanning.

The smoke-test job must depend on the container workflow job with `needs`. Do not use `always()` to bypass a failed build or scan.

## Check names and branch protection

Reusable workflows create GitHub Actions checks in the caller repository. Their displayed job/check names may differ from the old inline jobs. After a real run, record the exact names and coordinate any required-check/ruleset changes. Do not guess new required check names before GitHub reports them.
