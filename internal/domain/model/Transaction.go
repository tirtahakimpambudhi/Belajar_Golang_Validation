package model

import "time"

type Transaction struct {
	TransactionID string     `json:"transactionId" validate:"required,uuid"`
	Amount        float64    `json:"amount" validate:"required,gtefield=MinimumAmount,ltefield=MaximumAmount"`
	MaximumAmount float64    `json:"maximumAmount" validate:"required"`
	MinimumAmount float64    `json:"minimumAmount" validate:"required"`
	Description   string     `json:"description"`
	CreatedAt     *time.Time `json:"createdAt"`
	UpeatedAt     *time.Time `json:"upeatedAt"`
}

func NewTransaction(transactionID string, amount float64, maximumAmount float64, minimumAmount float64, description string, createdAt *time.Time, upeatedAt *time.Time) *Transaction {
	return &Transaction{TransactionID: transactionID, Amount: amount, MaximumAmount: maximumAmount, MinimumAmount: minimumAmount, Description: description, CreatedAt: createdAt, UpeatedAt: upeatedAt}
}
