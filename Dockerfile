FROM golang:1.19.0

WORKDIR /app

COPY . .

RUN go mod tidy && go mod vendor

RUN go build ./cmd/main.go

EXPOSE 9000

CMD ["./main"]
