FROM golang:1.4
MAINTAINER tobe tobeg3oogle@gmail.com

RUN apt-get update -y

ADD . /archci
WORKDIR /archci

RUN go get
RUN go build

CMD /bin/bash

