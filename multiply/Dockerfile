FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold
COPY . /go/src/github.com/MrMjauh/faas-scaffold/multiply
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/multiply

EXPOSE 8082

RUN govendor init
RUN govendor fetch -v +m
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/multiply/cmd/multiplyd
RUN go build main.go
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/multiply

CMD ./cmd/multiplyd/main -port=8082
