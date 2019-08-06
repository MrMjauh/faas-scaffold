FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
COPY . /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway

EXPOSE 8081

# Make sure we prebuild as much as possible
# So we dont get any nasty suprises
RUN govendor init
RUN govendor fetch -v +m
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway/cmd/gatewayd
RUN go build main.go
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/faas-gateway

CMD ./cmd/gatewayd/main -port=8081