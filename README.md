# Transaction Microservice

Microserviço desenvolvido em **Go (Golang)** utilizando o framework **Gin**, com o objetivo de validar transações financeiras conforme regras de negócio.

## Execução

### Clonar o repositório
```bash
git clone https://github.com/miguelmarcio/dafiti-transaction-microservice.git
cd transaction-microservice
```

### Instalar dependências
```bash
go mod tidy
```

### Executar o servidor
```bash
go run main.go
```

O servidor iniciará em: `http://localhost:8080`

### Endpoint principal

**POST** => `/validate`

Exemplo de Request:

```json
{
  "totalValue": 180.50,
  "payerID": "c7a2f0a8-5a25-4c3e-9c3d-1a2b3f4e5d6c",
  "payeeID": "f5b4e3d2-c1a0-9b8c-7f6e-5d4c3b2a1f0e",
  "items": [
    {
      "name": "Produto Físico 1",
      "quantity": 2,
      "unitPrice": 75.25,
      "type": "physical"
    },
    {
      "name": "Produto Físico 2",
      "quantity": 1,
      "unitPrice": 30.00,
      "type": "physical"
    }
  ]
}
```

Response Body (JSON):

- Sucesso (200 OK): `{"status": "approved"}`
- Rejeição (400 Bad Request): `{"status": "denied", "reason": "Motivo da rejeição."}`
- Erro de Payload (422 Unprocessable Entity): `{"error": "Payload inválido."}`
