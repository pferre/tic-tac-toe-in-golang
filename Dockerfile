FROM golang:1.14-alpine

WORKDIR /go/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV GOPATH="/go/"
ENV PATH="$PATH:$GOPATH/bin"