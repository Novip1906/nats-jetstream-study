FROM golang:alpine
WORKDIR /app
COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/api
CMD ["./main"]
