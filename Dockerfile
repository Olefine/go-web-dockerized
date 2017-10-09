FROM golang

ADD . /go/src/github.com/olefine/go-dockerized-web/

RUN go install github.com/olefine/go-dockerized-web/

ENTRYPOINT /go/bin/go-dockerized-web

EXPOSE 8080
