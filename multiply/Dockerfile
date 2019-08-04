FROM golang:1.11

# go stuff
RUN go get -u github.com/kardianos/govendor && \
    go get -u -v github.com/codeskyblue/fswatch

RUN mkdir -p /go/src/faas-scaffold

WORKDIR /go/src/faas-scaffold/multiply

EXPOSE 8080

CMD govendor init && govendor fetch -v +m && go run /go/src/faas-scaffold/multiply/cmd/multiplyd/main.go
