This is README

## Vanity gap test cases (RTDEV-101378)
Vanity root: `dors.com/testGoMajor` -> `https://github.com/dorsJfrog/testGoMajor`.

| Case | Module / package path | Directory | Tag |
|---|---|---|---|
| A root module (major-branch layout) | `dors.com/testGoMajor/v2` | `/` | `v2.0.2` |
| B sub-package (not a module) | `dors.com/testGoMajor/v2/greet` | `greet/` | (part of A) |
| C major version in subdirectory | `dors.com/testGoMajor/v3` | `v3/` | `v3.0.0` |
| D nested module | `dors.com/testGoMajor/tools` | `tools/` | `tools/v1.0.0` |
| E deep nested module | `dors.com/testGoMajor/cmd/tool` | `cmd/tool/` | `cmd/tool/v1.0.0` |

`v2.0.0` declares the wrong module path (`github.com/dorsJfrog/testGoMajor/v2`) and is kept as a negative case.
