FROM golang:1.23-alpine3.20 AS builder

ARG WD_NAME=/app
WORKDIR $WD_NAME

# build a static binary
ENV CGO_ENABLED=0
ENV GOROOT=/usr/local/go

# download and cache the dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOROOT/pkg \
    go mod download && go mod verify

COPY cmd cmd
COPY internal internal
COPY docs docs
RUN mkdir ./bin && go build -x -o ./bin ./cmd/... ./internal/... ./docs/...

FROM alpine:3.20 AS final

ARG WD_NAME=/app
WORKDIR $WD_NAME

COPY config config
COPY --from=builder $WD_NAME/bin/app $WD_NAME/bin/app

ENTRYPOINT ["./bin/app"]
