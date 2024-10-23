package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, password string) *Account {
	acc, err := NewAccount(fname, lname, password)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Account created:", acc.AccountNumber)
	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Ahmed", "Abo Elfadle", "hunt45")
}

func main() {
	seed := flag.Bool("seed", false, "Seed Database")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding database...")
		seedAccounts(store)
	}

	server := NewAPIServer(":8080", store)
	server.Run()
}
