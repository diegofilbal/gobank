FROM golang:1.21.1

# Instala o tree
RUN apt-get update && apt-get install -y tree

WORKDIR /server

COPY . .

RUN tree

RUN cd server && go mod tidy

RUN cd server/api && go build -o main main.go

EXPOSE 8080

CMD ["./main"]
