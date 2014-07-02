FROM ubuntu:14.04

ENV HOME /root
ENV GOPATH /root/go
ENV PATH /root/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN apt-get update
RUN apt-get install -y git-core wget

RUN wget http://golang.org/dl/go1.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.3.linux-amd64.tar.gz

ADD . /root/go/src/github.com/stevenle/shortn
RUN cd /root/go/src/github.com/stevenle/shortn
RUN go get -d -v github.com/stevenle/shortn/...
RUN go install -v github.com/stevenle/shortn
