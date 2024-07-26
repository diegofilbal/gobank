FROM golang:1.21.1

ADD . /server

WORKDIR /server/api

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
