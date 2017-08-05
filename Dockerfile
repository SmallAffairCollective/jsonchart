FROM golang:alpine3.6

MAINTAINER Charlie Lewis <defermat@gmail.com>

ADD . /go/src/github.com/SmallAffairCollective/jsonchart
RUN go get github.com/SmallAffairCollective/jsonchart

WORKDIR /go/src/github.com/SmallAffairCollective/jsonchart

ENTRYPOINT ["/go/bin/jsonchart"]
CMD ["-h"]
