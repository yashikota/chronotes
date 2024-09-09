# Go
.PHONY: lint
lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go vet ./...
	staticcheck ./...

.PHONY: test
test:
	go build ./...
	go test -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: dev
dev:
	docker compose up --build

.PHONY: build
build:
	docker build -t 58hack .

# API
.PHONY: api-lint
api-lint: api-bundle
	docker run --rm -v ${PWD}:/spec redocly/cli lint --config docs/api/redoc.yaml docs/api/bundled.yaml

.PHONY: bundle
bundle:
	docker run --rm -v ${PWD}:/spec redocly/cli bundle docs/api/openapi.yaml -o docs/api/bundled.yaml

.PHONY: docs
docs:
	docker run --rm -v ${PWD}:/spec redocly/cli build-docs docs/api/openapi.yaml --o docs/api/redoc.html
