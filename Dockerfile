FROM golang:1.24.2-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 5000
CMD ["./main"]