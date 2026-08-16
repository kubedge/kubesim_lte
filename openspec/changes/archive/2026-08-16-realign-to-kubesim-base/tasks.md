# Tasks — realign-to-kubesim-base

- [x] Wait for `kubesim_base` to publish the new tag (its multimodule-tag-realign change). — v0.1.25 published.
- [x] `go get` each kubesim_base sub-module this repo requires (`config`/`connected`/`grpc/go`) at v0.1.25.
- [x] `go mod tidy`; confirm every kubesim_base require moved to v0.1.25. — all three at v0.1.25.
- [x] `go build ./... && go vet ./... && go test ./... -race` green. — `go build -o /dev/null ./...` (cmd-collision), vet + test green (no test files).
