# Chronotes

[![codecov](https://codecov.io/github/yashikota/chronotes/graph/badge.svg?token=8LK1D9KWN5)](https://codecov.io/github/yashikota/chronotes)
[![Go Report Card](https://goreportcard.com/badge/github.com/yashikota/chronotes)](https://goreportcard.com/report/github.com/yashikota/chronotes)

> [!NOTE]
> フロントエンドのリポジトリは [GenichiMaruo/chronotes-front](https://github.com/GenichiMaruo/chronotes-front) にあります。  

## Development

```sh
go install github.com/go-task/task/v3/cmd/task@latest
```

```txt
task: Available tasks for this project:
* all:                Run all tasks
* build:              Build Docker image
* default:            Display this help message
* dev:                Run development environment
* actions:lint:       Lint GitHub Actions
* api:fmt:            Format OpenAPI specification
* api:lint:           Lint API documentation
* api:split:          Split OpenAPI specification
* api:tsp:            Generate Open API from TypeSpec
* dev:re:             Rebuild service
* docker:lint:        Lint Dockerfile
* go:fmt:             Format Go code
* go:lint:            Lint Go code
* go:test:            Run Go tests
* md:lint:            Lint Markdown files
```
