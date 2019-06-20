FROM golang as base

ENV GO111MODULE=on
ENV HTTP=8080
ENV GRPC=8081

EXPOSE $HTTP
EXPOSE $GRPC

COPY . build

RUN ls -al . && ls -al src

WORKDIR build

RUN go mod download

RUN go build ./cmd/discoverysvr

ENTRYPOINT ["./discoverysvr"]

