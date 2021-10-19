# build stage
FROM golang:alpine AS build-env
WORKDIR /opt/goblin
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags "netgo" -ldflags '-s -w' -o goblin

# final stage
# hadolint ignore=DL3007
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /opt/goblin/goblin /app/
ENTRYPOINT ["./goblin"]
