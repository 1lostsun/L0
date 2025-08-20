FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG SERVICE=ProducerService
RUN go build -o server ./ProducerService/cmd/server

EXPOSE 8080
CMD ["./server"]