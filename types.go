package main

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRespone struct {
	AccountNumber int64  `json:"account_number"`
	Token         string `json:"token"`
}

type LoginRequest struct {
	AccountNumber int64  `json:"account_number"`
	Password      string `json:"password"`
}

type TrandferRequest struct {
	ToAccount int `json:"account_number"`
	Amount    int `json:"amount"`
}

type CrerateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type Account struct {
	ID                int       `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	AccountNumber     int64     `json:"account_number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewAccount(fisrtName, lastName, password string) (*Account, error) {
	encPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         fisrtName,
		LastName:          lastName,
		AccountNumber:     int64(rand.Intn(1000000)),
		EncryptedPassword: string(encPassword),
		CreatedAt:         time.Now().UTC(),
	}, nil
}

func (a *Account) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(password)) == nil
}
