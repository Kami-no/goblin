# build stage
FROM golang:1.18-alpine AS build-env
WORKDIR /opt/goblin
# hadolint ignore=DL3018
RUN apk add --no-cache git
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags "netgo" -ldflags '-s -w' -o goblin

# final stage
# hadolint ignore=DL3007
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /opt/goblin/goblin /app/
CMD ["./goblin"]
