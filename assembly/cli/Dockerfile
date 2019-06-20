FROM golang as base

ENV GO111MODULE=on

COPY . build

WORKDIR build

RUN go mod download

RUN go build ./cmd/discoverysvr-cli

ENTRYPOINT ["./discoverysvr-cli"]

