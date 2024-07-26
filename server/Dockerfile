FROM golang:1.21.1

WORKDIR /server

COPY . .

RUN ls -la 

RUN ls -la /server

RUN go mod tidy

RUN go build -o main ./api/main.go

EXPOSE 8080

CMD ["./main"]
