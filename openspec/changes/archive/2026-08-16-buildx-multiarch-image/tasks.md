# Tasks — buildx-multiarch-image

- [x] Confirm buildx + a running builder (colima on Apple-Silicon). — colima + buildx v0.35.0.
- [x] For each built image, add a buildx target `docker buildx build --platform linux/arm64,linux/amd64 -t <image>:<tag>`. — single `docker-buildx` target (one image).
- [x] Make each Dockerfile multi-stage (compile per TARGETOS/TARGETARCH; don't copy a prebuilt amd64 binary). — `build/Dockerfile.buildkit` pinned `FROM --platform=$BUILDPLATFORM golang:1.26`, cross-compiles via GOARCH.
- [x] Remove the `_AMD64/_ARM64V8/_ARM32V7` image-name variables. — retired (commented for rollback).
- [~] `docker buildx imagetools inspect` each image → manifest list incl. linux/arm64. — cross-compile validated locally (both arches built, manifest list `sha256:4d9fdb79…` exported without push); registry `--push` + inspect intentionally held (would overwrite live `v0.2.20`).
- [n/a] If deployed by an operator, update that operator's example CR. — standalone sim, no paired operator/CR.
