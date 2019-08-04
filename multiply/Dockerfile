FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u -v github.com/codeskyblue/fswatch

RUN mkdir -p /go/src/github.com/MrMjauh/faas-scaffold
COPY . /go/src/github.com/MrMjauh/faas-scaffold
WORKDIR /go/src/github.com/MrMjauh/faas-scaffold/multiply

EXPOSE 8080

RUN govendor init
RUN govendor fetch -v +m

CMD go run /go/src/github.com/MrMjauh/faas-scaffold/multiply/cmd/multiplyd/main.go -port=8080
