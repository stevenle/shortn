FROM golang:1.6

ADD . /go/src/github.com/stevenle/shortn
RUN go install github.com/stevenle/shortn

ENTRYPOINT /go/bin/shortn
EXPOSE 8080
EXPOSE 9090
