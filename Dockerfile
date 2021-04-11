FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o client .

CMD ["./client"]