FROM golang:1.21.1

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./api/main.go

EXPOSE 8080

CMD ["./main"]
