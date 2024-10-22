package main

import "math/rand"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber int64  `json:"account_number"`
	Balance       int64  `json:"balance"`
}

func NewAccount(fisrtName, lastName string) *Account {
	return &Account{
		ID:            rand.Intn(10000),
		FirstName:     fisrtName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(1000000)),
	}
}
