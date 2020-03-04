FROM golang:1.14.0

WORKDIR /app

RUN go get -u google.golang.org/grpc

ADD main.go /app
ADD api /go/src/github.com/tarikbauer/go_vs_py_benchmark/go_grpc/api

RUN go build main.go

CMD /app/main
