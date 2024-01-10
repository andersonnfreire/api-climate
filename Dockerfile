FROM golang:1.19.0

WORKDIR /app

# Copy the .env file
COPY .env .

# Copy the entire project (including the cmd directory)
COPY . .

RUN go mod tidy && go mod vendor

# Compile the code in the cmd directory and generate the main executable
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
