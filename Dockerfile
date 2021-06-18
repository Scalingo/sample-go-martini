FROM golang:1.16

RUN go get github.com/cespare/reflex
ADD . /go/src/github.com/Scalingo/sample-go-martini
WORKDIR /go/src/github.com/Scalingo/sample-go-martini
EXPOSE 3000
RUN go install
CMD /go/bin/sample-go-martini

