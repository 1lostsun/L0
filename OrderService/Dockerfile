FROM golang:1.24.6

WORKDIR /app

COPY OrderService/go.mod OrderService/go.sum ./
RUN go mod download
COPY OrderService/ .

ARG SERVICE=OrderService
RUN go build -o server ./cmd/server

EXPOSE 8081
CMD ["./server"]