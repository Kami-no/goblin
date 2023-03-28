# build stage
ARG GO_VERSION=1.20
FROM golang:${GO_VERSION}-bullseye AS build

WORKDIR /opt/build

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN go test ./... \
    && CGO_ENABLED=0 go build -a -tags 'netgo' -ldflags '-s -w' -o goblin

# artefact stage
# hadolint ignore=DL3007
FROM gcr.io/distroless/static-debian11:latest

COPY --from=build /opt/build/goblin /usr/local/bin/goblin
USER 1000
CMD ["goblin"]
