package main

import "math/rand"

type Account struct {
	ID            int
	FirstName     string
	LastName      string
	AccountNumber int64
	Balance       int64
}

func NewAccount(fisrtName, lastName string) *Account {
	return &Account{
		ID:            rand.Intn(10000),
		FirstName:     fisrtName,
		LastName:      lastName,
		AccountNumber: int64(rand.Intn(1000000)),
	}
}
