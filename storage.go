package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=bank password=0123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
            last_name VARCHAR(255) NOT NULL,
            account_number BIGINT UNIQUE NOT NULL,
            balance FLOAT DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	query := `INSERT INTO accounts (first_name, last_name, account_number, balance, created_at)
		VALUES ($1, $2, $3, $4, $5)`

	result, err := s.db.Exec(
		query,
		acc.FirstName,
		acc.LastName,
		acc.AccountNumber,
		acc.Balance,
		acc.CreatedAt,
	)
	fmt.Printf("+%v\n", result)
	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	query := `SELECT id, first_name, last_name, account_number, balance, created_at FROM accounts WHERE id = $1`
	row, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		return scanIntoAccount(row)
	}

	return nil, fmt.Errorf("accont with id %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`SELECT id, first_name, last_name, account_number, balance, created_at FROM accounts`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil

}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}
