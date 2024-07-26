FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./api/main.go

EXPOSE 8080

CMD ["./main"]
