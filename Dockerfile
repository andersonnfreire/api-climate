FROM golang:1.19.0

WORKDIR /app

COPY .env .

COPY . .

RUN go mod tidy && go mod vendor


RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
