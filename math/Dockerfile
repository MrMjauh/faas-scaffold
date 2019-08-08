FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold
COPY . /go/src/github.com/MrMjauh/faas-scaffold/math
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/math

EXPOSE 8082

RUN govendor init
RUN govendor fetch -v +m
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/math/cmd/mathd
RUN go build main.go
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/math

CMD ./cmd/mathd/main -port=8082 -mode=$mode