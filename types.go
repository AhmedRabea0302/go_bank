package main

import (
	"math/rand"
	"time"
)

type CrerateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	AccountNumber int64     `json:"account_number"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewAccount(fisrtName, lastName string) *Account {
	return &Account{
		FirstName:     fisrtName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(1000000)),
		CreatedAt:     time.Now().UTC(),
	}
}
