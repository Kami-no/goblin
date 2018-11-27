# build stage
FROM golang:alpine AS build-env
COPY main.go main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goblin

# final stage
# hadolint ignore=DL3007
FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/goblin /app/
ENTRYPOINT ["./goblin"]
