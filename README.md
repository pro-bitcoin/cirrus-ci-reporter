# Bitcoin cirrus ci reporter

[![Test](https://github.com/pro-bitcoin/cirrus-ci-reporter/actions/workflows/test.yml/badge.svg)](https://github.com/pro-bitcoin/cirrus-ci-reporter/actions/workflows/test.yml) [![golangci-lint](https://github.com/pro-bitcoin/cirrus-ci-reporter/actions/workflows/lint.yml/badge.svg)](https://github.com/pro-bitcoin/cirrus-ci-reporter/actions/workflows/lint.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/pro-bitcoin/cirrus-ci-reporter)](https://goreportcard.com/report/github.com/pro-bitcoin/cirrus-ci-reporter) [![Go Reference](https://pkg.go.dev/badge/github.com/pro-bitcoin/cirrus-ci-reporter.svg)](https://pkg.go.dev/github.com/pro-bitcoin/cirrus-ci-reporter) [![codecov](https://codecov.io/gh/pro-bitcoin/cirrus-ci-reporter/branch/main/graph/badge.svg?token=Y5K4SID71F)](https://codecov.io/gh/pro-bitcoin/cirrus-ci-reporter)

# Makefile Targets
```sh
$> make
build                          build golang binary
clean                          clean up environment
cover                          display test coverage
docker-build                   dockerize golang application
fmtcheck                       run gofmt and print detected files
fmt                            format go files
help                           list makefile targets
install                        install golang binary
lint-fix                       fix
lint                           lint go files
pre-commit                     run pre-commit hooks
run                            run the app
test                           run go tests
```

# Contribute
If you find issues in that setup or have some nice features / improvements, I would welcome an issue or a PR :)
