FROM golang:latest

MAINTAINER liyao.miao@yeepay.com 

ADD webapp-test /go/bin/

RUN ulimit -n 10240

EXPOSE 9999

ENTRYPOINT ["/go/bin/webapp-test"]

