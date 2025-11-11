package message

const (
	ReasonInvalidPayerUUID = "PayerID não é um UUID válido."
	ReasonInvalidPayeeUUID = "PayeeID não é um UUID válido."
	ReasonSameIDs          = "PayerID e PayeeID não podem ser iguais."
	ReasonEmptyItems       = "A lista de itens não pode estar vazia."
	ReasonMixedItemTypes   = "Não é permitido misturar itens físicos e digitais."
	ReasonTotalMismatch    = "A soma dos itens não bate com o valor total."
	ReasonInvalidPayload   = "Payload inválido."
)
