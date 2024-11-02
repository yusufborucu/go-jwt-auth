# Go JWT Authentication

This project is a JWT authentication example built with Go and MySQL.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Usage](#usage)

## Features

- Routing requests with gorilla/mux
- JWT authentication
- MySQL integration using Docker
- ORM support with GORM
- Validation
- Bcrypt hashing

## Technologies Used

- **Go**: Programming language for backend development
- **MySQL**: Database management
- **Docker**: Containerization for database
- **GORM**: ORM (Object-Relational Mapping) for Go
- **gorilla/mux**: Routing requests

## Setup

### 1. Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)

### 2. Clone the Repository

```bash
git clone https://github.com/yusufborucu/go-jwt-auth.git
cd go-jwt-auth
```

### 3. Set Up Environment Variables

Add an .env file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=testuser
DB_PASSWORD=testpassword
DB_NAME=testdb

JWT_SECRET=my_secret_key
```

### 4. Start MySQL with Docker

Run the following command to start the MySQL container:

```bash
docker-compose up -d
```

### 5. Install Required Go Packages
```bash
go mod tidy
```

### 6. Run the Application
```bash
go run main.go
```

The API should now be running at http://localhost:8080.

## Usage

You can test the API’s core features using the following example cURL commands.

### Register User

```bash
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "johndoe@example.com", "password": "123456"}'
```

### Get Token

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email": "johndoe@example.com", "password": "123456"}'
```

### Get Profile Infos with Token

```bash
curl http://localhost:8080/profile -H "Authorization: Bearer token"
```