package main

import (
	"log"
	"transaction-microservice/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	transactionHandler := handler.NewTransactionHandler()

	r.POST("/validate", transactionHandler.Validate)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
