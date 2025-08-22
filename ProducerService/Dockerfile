FROM golang:1.24.6

WORKDIR /app

COPY ProducerService/go.mod ProducerService/go.sum ./
RUN go mod download
COPY ProducerService/ .

ARG SERVICE=ProducerService
RUN go build -o server ./cmd/server

EXPOSE 8080
CMD ["./server"]