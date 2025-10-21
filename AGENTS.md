# Repository Guidelines

## Project Structure & Module Organization
The core library lives in `version/`, where `banner.go` and `version.go` expose parsing, formatting, and banner helpers, alongside co-located unit tests (`banner_test.go`, `version_test.go`). The `example/` folder hosts runnable examples and doctest-style demos that double as regression coverage. Release metadata is tracked in the root-level `VERSION` file, and the `Makefile` centralises repeatable developer tasks.

## Build, Test, and Development Commands
- `make deps` downloads module dependencies.
- `make build` compiles all Go packages (`go build ./...`).
- `make test` runs the full test suite; use `go test ./example -run Example` when iterating on demonstrations.
- `make lint` runs `golangci-lint`; install it first with `make install-lint` if missing.
- `make scan` invokes `gosec` to emit `gosec.sarif`; pair with `make clean` to delete stale reports.
- `make ci` chains deps, build, test, lint, and security scan for local pre-flight checks.

## Coding Style & Naming Conventions
Follow idiomatic Go style: rely on tabs for indentation and keep files `gofmt`-clean (`go fmt ./version ./example`). Exported types and functions use PascalCase (`BannerOptions`), while unexported helpers stay in camelCase. Maintain descriptive test names (`TestBannerOptionsDefaults`) and prefer short receiver names. Linting expectations mirror `golangci-lint` defaults; address warnings rather than suppressing them.

## Testing Guidelines
Tests use Go's standard `testing` package with table-driven cases where practical. Place unit tests beside their targets and name them `Test*`. Keep examples under `example/` or `Example*` functions so they execute via `go test`. Aim to run `go test ./... -cover` before submitting to watch for unintended coverage regressions.

## Commit & Pull Request Guidelines
Commits follow Conventional Commit semantics observed in history (`feat(version): … (#3)`). Use clear scopes (`feat(version)`, `fix(example)`) and reference linked issues or PR numbers in parentheses. Before opening a PR, ensure `make test`, `make lint`, and `make scan` pass, update documentation if behaviour changes, and include screenshots or terminal output when relevant features affect banner appearance.

## Security & Versioning Tips
Treat `make scan` findings as release blockers. When bumping releases, update the `VERSION` file and confirm the new semver string parses via `go test ./example -run Example`. Avoid embedding secrets in banners or examples; prefer environment variables in consumer code.
