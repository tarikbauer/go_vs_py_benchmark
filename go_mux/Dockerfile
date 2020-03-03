FROM golang:1.14.0

WORKDIR /app

RUN go get -u github.com/gorilla/mux

ADD main.go /app
ADD helper.go /app

RUN go build main.go helper.go

CMD /app/main

