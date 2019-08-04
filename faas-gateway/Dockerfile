FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u -v github.com/codeskyblue/fswatch

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
COPY . /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway

EXPOSE 80

RUN govendor init
RUN govendor fetch -v +m

CMD go run /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway/cmd/gatewayd/main.go -port=80