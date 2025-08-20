FROM golang:latest

WORKDIR /order

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server