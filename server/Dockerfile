FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD ./src /go/src/app

RUN go get -v
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq