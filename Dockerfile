FROM golang:alpine3.6

MAINTAINER defermat <defermat@defermat.net>

RUN apk add --update git
ADD . /go/src/github.com/SmallAffairCollective/jsonchart
RUN go get github.com/SmallAffairCollective/jsonchart

WORKDIR /go/src/github.com/SmallAffairCollective/jsonchart

ENTRYPOINT ["/go/bin/jsonchart"]
CMD ["http://genit/genit", "2", "15"]
