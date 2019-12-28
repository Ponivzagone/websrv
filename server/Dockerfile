FROM golang:1.8

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD ./src/controllers /go/src/app/controllers
ADD ./src/handler /go/src/app/handler
ADD ./src/models /go/src/app/models


RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/rs/cors
RUN go get github.com/gorilla/handlers