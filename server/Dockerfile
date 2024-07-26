FROM golang:1.21.1

ADD . /server

WORKDIR /server

COPY . .

RUN go mod tidy

RUN go build -o main ./api/main.go

EXPOSE 8080

CMD ["./main"]
