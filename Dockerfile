FROM golang:1.4
MAINTAINER tobe tobeg3oogle@gmail.com

RUN go get github.com/tools/godep

ADD . /go/src/github.com/ArchCI/archci
WORKDIR /go/src/github.com/ArchCI/archci
RUN godep go build -ldflags "-X main.GitVersion `git rev-parse HEAD` -X main.BuildTime `date -u '+%Y-%m-%d_%I:%M:%S'`"

EXPOSE 10010

CMD ./archci

