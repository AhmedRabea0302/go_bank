# Go Bank Application

## Overview

This project is a backend application built with **Go**, simulating basic banking operations. It provides API endpoints to manage accounts, perform money transfers between accounts, and track transaction history. The application uses **PostgreSQL** as the database and implements robust data handling with tools like **GORM** and **SQLC** for efficient database interactions.

## Features

- **Account Management**: Create and manage bank accounts with attributes such as owner name, balance, and currency.
- **Transactions**: Transfer money between accounts and maintain transaction history.
- **Authentication**: (Optional) Secure API endpoints using JWT for stateless authentication.
- **Database Interactions**: CRUD operations on the database using GORM and SQLC.
  
## Prerequisites

- **Go** version 1.16 or higher.
- **PostgreSQL** installed locally or via Docker.
- **Docker** (optional for database setup).

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/AhmedRabea0302/go_bank.git
cd go_bank
```

## 2. Install Dependencies

```bash

go mod tidy

```

## 3. Database Setup
### Using Docker (Recommended)
**To quickly spin up a PostgreSQL instance using Docker**

```bash

docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

```

## 4. Seeding Database

```bash

./bin/gobank --seed

```

## 5. Start the Application

```bash

make run

```
