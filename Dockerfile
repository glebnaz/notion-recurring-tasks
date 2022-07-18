#---Build stage---
FROM golang:1.18 AS builder
COPY . /go/src/notion-recurring-tasks
WORKDIR /go/src/notion-recurring-tasks/cmd

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
ENTRYPOINT ["go/bin/service"]