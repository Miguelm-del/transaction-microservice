package model

type TransactionStatus struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}
