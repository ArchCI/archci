FROM golang:1.4
MAINTAINER tobe tobeg3oogle@gmail.com

ADD . /go/src/github.com/ArchCI/archci
WORKDIR /go/src/github.com/ArchCI/archci

RUN go get
RUN go build

EXPOSE 10010

CMD ./archci

