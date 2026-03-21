FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/worker
CMD ["./main"]
