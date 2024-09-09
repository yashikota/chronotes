# 58hack

## Go

### Dev

```sh
docker compose up --build
```

or

```sh
make dev
```

### Build

```sh
docker build -t 58hack .
```

or

```sh
make build
```

### Lint

```sh
go install honnef.co/go/tools/cmd/staticcheck@latest
go vet ./...
staticcheck ./...
```

or

```sh
make lint
```

### Test

```sh
go build ./...
go test -v ./...
```

or

```sh
make test
```

### Fmt

```sh
go fmt ./...
```

or

```sh
make fmt
```

## API

### Lint

```sh
docker run --rm -v ${PWD}:/spec redocly/cli lint --config docs/api/redoc.yaml docs/api/bundled.yaml
```

or

```sh
make api-lint
```

### Bundle

```sh
docker run --rm -v ${PWD}:/spec redocly/cli bundle docs/api/openapi.yaml -o docs/api/bundled.yaml
```

or

```sh
make bundle
```

### Docs

```sh
docker run --rm -v ${PWD}:/spec redocly/cli build-docs docs/api/openapi.yaml --o docs/api/redoc.html
```

or

```sh
make docs
```
