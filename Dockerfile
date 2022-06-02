FROM golang:1.17

RUN go install github.com/cespare/reflex@latest
ADD . /go/src/github.com/Scalingo/sample-go-martini
WORKDIR /go/src/github.com/Scalingo/sample-go-martini
EXPOSE 3000
RUN go install
CMD /go/bin/sample-go-martini

