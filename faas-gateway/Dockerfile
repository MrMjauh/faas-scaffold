FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u -v github.com/codeskyblue/fswatch

RUN mkdir -p /go/src/faas-scaffold/faas-gateway
RUN mkdir -p /go/src/faas-scaffold/docker
RUN mkdir -p /go/src/faas-scaffold/commons
WORKDIR /go/src/faas-gateway

EXPOSE 80

CMD govendor init && govendor fetch -v +m && go run /go/src/faas-scaffold/faas-gateway/cmd/gatewayd/main.go