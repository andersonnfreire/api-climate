FROM golang:1.19.0

WORKDIR /app

# Copy the .env file
COPY .env .

# Copy the entire project (including the cmd directory)
COPY . .

# Copie a imagem para o diretório de trabalho (certifique-se de ajustar o caminho conforme necessário)
COPY pkg/handlers/img/imagem.png /app

RUN go mod tidy && go mod vendor

# Compile the code in the cmd directory and generate the main executable
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
