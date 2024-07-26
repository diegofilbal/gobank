FROM golang:1.21.1

WORKDIR /server

COPY . .

RUN cd server && go mod tidy

RUN cd server/api && go build -o main main.go

EXPOSE 8080

CMD ["./main"]
