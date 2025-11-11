# Etapa 1 - Build da aplicação
FROM golang:1.25.2 AS builder

WORKDIR /app

# Copia os arquivos do módulo e baixa dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código para dentro da imagem
COPY . .

# Compila o binário
RUN go build -o main .

# Etapa 2 - Runtime (imagem leve)
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/main .

# Porta padrão usada no main.go
EXPOSE 8080

CMD ["./main"]
