# Tasks — adopt-go-ci

- [x] From the meta session, run `/alemax:update-skills` so the class-M set (incl. `ci.yml`) is staged for this repo. — broadcast `meta-broadcast/deliver-class-m-templates-ciyml--hygiene`.
- [x] In this repo's session, run `/alemax:complete-update` to apply the update branch onto the working branch.
- [x] Confirm `.github/workflows/ci.yml` present and its jobs gate on `go.mod`. — go jobs gate on `detect.outputs.go == 'true'`.
- [x] Trial push the branch; confirm `go-build`/`go-vet`/`go-test`/`golangci-lint` are green. — PR #1 CI green (golangci v2.12.2; `go-build` uses `-o /dev/null` for the cmd-collision).
- [x] Confirm the rest of class-M landed: `.editorconfig`, `.gitattributes`, `.github/*`, `dependabot.yml`, `.pre-commit-config.yaml`, `bin/set-secret.sh`. — all present.
