# build stage
FROM golang:alpine AS build-env
ADD main.go main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goblin

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/goblin /app/
ENTRYPOINT ./goblin
