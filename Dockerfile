### ----------------- ###
### Development image ###
### ----------------- ###
ARG GO_VERSION=latest
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x
# hadolint ignore=DL3059
RUN go install github.com/air-verse/air@latest

COPY . .
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

### ---------------- ###
### Production image ###
### ---------------- ###
ARG GO_VERSION=latest
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build

WORKDIR /app

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage bind mounts to go.sum and go.mod to avoid having to copy them into the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# This is the architecture you're building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
ARG TARGETARCH

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# Leverage a bind mount to the current directory to avoid having to copy the
# source code into the container.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/chronotes

### ---------------- ###
### Production image ###
### ---------------- ###
FROM gcr.io/distroless/static-debian12:nonroot AS final

# Copy the executable from the "build" stage.
COPY --from=build /bin/server /bin/

WORKDIR /app

# Timezone
ENV TZ=Asia/Tokyo

# Expose the port that the application listens on.
EXPOSE 8080

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
