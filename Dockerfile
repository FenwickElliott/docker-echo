FROM golang:1.10

RUN mkdir /app && mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go build -o /app/main .

CMD ["/app/main"]
EXPOSE 9090