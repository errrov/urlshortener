FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go
EXPOSE 8080
CMD ["/cmd"]