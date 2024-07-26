FROM golang:1.21.1

RUN apt-get update && apt-get install -y tree

WORKDIR /server

COPY . ./server

RUN ls -la 

RUN ls -la /server

RUN tree

RUN go mod tidy ./server

RUN go build -o main ./api/main.go

EXPOSE 8080

CMD ["./main"]
