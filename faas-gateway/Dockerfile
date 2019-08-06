FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
COPY . /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway

EXPOSE 8081

RUN govendor init
RUN govendor fetch -v +m

CMD go run /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway/cmd/gatewayd/main.go -port=8081