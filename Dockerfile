FROM golang:1.12.7

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go get -d ./...
RUN go build -o main .

CMD ["/app/main"]
