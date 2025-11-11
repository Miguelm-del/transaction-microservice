package handler

import (
	"net/http"
	"transaction-microservice/src/message"
	"transaction-microservice/src/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandlerInterface interface {
	Validate(c *gin.Context)
}

type TransactionHandler struct {
}

func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{}
}

func (th *TransactionHandler) Validate(c *gin.Context) {
	var transactionDTO model.Transaction

	if err := c.ShouldBindJSON(&transactionDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": message.ReasonInvalidPayload})
		return
	}

	if _, err := uuid.Parse(transactionDTO.PayerID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "denied", "reason": message.ReasonInvalidPayerUUID})
		return
	}

	if _, err := uuid.Parse(transactionDTO.PayeeID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "denied", "reason": message.ReasonInvalidPayeeUUID})
		return
	}

	if transactionDTO.PayerID == transactionDTO.PayeeID {
		c.JSON(http.StatusBadRequest, gin.H{"status": "denied", "reason": message.ReasonSameIDs})
		return
	}

	if len(transactionDTO.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "denied", "reason": message.ReasonEmptyItems})
		return
	}

	total := calcTotalByItems(transactionDTO.Items)

	if total != transactionDTO.TotalValue {
		c.JSON(http.StatusBadRequest, gin.H{"status": "denied", "reason": message.ReasonTotalMismatch})
		return
	}

	if hasMixedItemTypes(transactionDTO.Items) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "denied",
			"reason": message.ReasonMixedItemTypes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "approved"})
}

func calcTotalByItems(items []model.Item) float64 {
	var total float64

	for _, item := range items {
		total += float64(item.Quantity) * item.UnitPrice
	}

	return total
}

func hasMixedItemTypes(items []model.Item) bool {
	if items == nil || len(items) == 0 {
		return false
	}

	firstType := items[0].Type
	for _, item := range items[1:] {
		if item.Type != firstType {
			return true
		}
	}
	return false
}
