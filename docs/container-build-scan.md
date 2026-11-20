# Reusable container build and scan

Workflow: `.github/workflows/container-build-scan.yml`

Builds one local image archive on `ubuntu-24.04` using Buildx, scans that archive using Trivy, and can upload the successfully scanned archive for a caller-owned smoke-test job. The scan report is uploaded as JSON when Trivy produces it. This workflow does not render a custom Markdown summary. It never logs into a registry or pushes an image.

Default scan policy matches the initial `f1_williams_assessment` CI: OS and library vulnerabilities, HIGH and CRITICAL severities, including unfixed findings. Scanner/database errors fail the job. Trivy caching is disabled.

## Inputs

| Input | Type | Default | Notes |
| --- | --- | --- | --- |
| `context` | string | `.` | Local Docker build context inside the checked-out repository. |
| `dockerfile` | string | `Dockerfile` | Dockerfile path relative to repository root. |
| `image-tag` | string | `ci-image:local` | Local tag used consistently for build, scan, and exported image. |
| `upload-image` | boolean | `false` | Export the image only after a successful build and scan. |
| `image-artifact-name` | string | `container-image` | Artifact name for the exported `image.tar`. |
| `scan-artifact-name` | string | `trivy-image-report` | Artifact name for the JSON report when Trivy produced one. |
| `severity` | string | `HIGH,CRITICAL` | Comma-separated Trivy severity list. Supported values: UNKNOWN, LOW, MEDIUM, HIGH, CRITICAL. |

Artifact names must be unique within a workflow run if the caller invokes this workflow more than once. Image artifacts expire after one day; scan reports expire after seven days.

## Outputs

| Output | Meaning |
| --- | --- |
| `image-tag` | Successfully built and scanned local image reference. |
| `image-artifact-name` | Artifact name when image export is enabled; empty when disabled. |

Outputs identify the artifact and tag but do not transfer an image between runners.

## Caller example

```yaml
jobs:
  container:
    uses: Delta-a-Sierra/workflows/.github/workflows/container-build-scan.yml@<FULL_COMMIT_SHA>
    with:
      image-tag: my-service:local
      upload-image: true
      image-artifact-name: container-image
```

For a smoke test, add a caller-owned job with `needs: container`, download `needs.container.outputs.image-artifact-name`, load `image.tar`, and test `needs.container.outputs.image-tag`. The smoke test then exercises the exact scanned image. A failed build or scan prevents the dependent job from running.

## Boundaries

- Local Docker context and Dockerfile only; remote contexts are not accepted.
- No build arguments, build secrets, registry login, image push, multi-platform build, or private-base-image authentication.
- Docker builds use `--pull`, so base images and network access are required.
- Docker Buildx writes an archive for Trivy to scan. The same archive is optionally uploaded for downstream jobs; no separate `docker save` command is needed.
- The scanner's JSON report is the diagnostic artifact. If Trivy fails before producing a report, the workflow keeps the scan failure and emits an artifact warning rather than masking the failure.
- Do not treat an absent report after an early build failure as a clean scan.
