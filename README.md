# Ticket Management System

A RESTful Ticket Management System built with **Go**, **Gin**, **GORM**, and **SQLite**. The application provides JWT-based authentication and allows authenticated users to create and manage their own support tickets.

## Features

- User Registration
- User Login with JWT Authentication
- Create Ticket
- Get All Tickets (User Specific)
- Get Ticket By ID 
- Update Ticket Status
- Ticket Ownership Validation
- Status Transition Validation
- Docker Support

---

## Tech Stack

- Go
- Gin
- GORM
- SQLite
- JWT
- bcrypt
- Docker

---

## Project Structure

```
.
├── cmd
│   └── main.go
├── internal
│   ├── config
│   ├── dto
│   ├── handlers
│   ├── middleware
│   ├── models
│   ├── repository
│   ├── routes
│   ├── services
│   └── utils
├── Dockerfile
├── .env.example
├── go.mod
└── README.md
```

---

## Environment Variables

Create a `.env` file in the project root.

Example:

```env
PORT=8080

JWT_SECRET=your_jwt_secret
```

Default PORT = 8080

---

## Running Locally

### Install dependencies

```bash
go mod download
```

### Start the application

```bash
go run cmd/main.go
```

The server starts on

```
http://localhost:8080
```

---

## Running with Docker

### Build

```bash
docker build -t ticket-system .
```

### Run

```bash
docker run -p 8080:8080 ticket-system
```

---

## API Endpoints

### Health Check

```
GET /health
```

---

### Authentication

#### Register

```
POST /auth/register
```

Request

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

---

#### Login

```
POST /auth/login
```

Request

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

Response

```json
{
  "token": "<jwt_token>"
}
```

---

### Tickets

All ticket endpoints require

```
Authorization: Bearer <jwt_token>
```

---

#### Create Ticket

```
POST /tickets
```

Request

```json
{
  "title": "Payment Issue",
  "description": "Unable to complete payment."
}
```

---

#### Get All Tickets

```
GET /tickets
```

---

#### Get Ticket By ID

```
GET /tickets/{id}
```

---

#### Update Ticket Status

```
PATCH /tickets/{id}/status
```

Request

```json
{
  "status": "in_progress"
}
```

Allowed status transitions:

```
open
    ↓
in_progress
    ↓
closed
```

Closed tickets cannot be reopened.

---

## Authentication

Protected endpoints require a JWT in the Authorization header.

```
Authorization: Bearer <jwt_token>
```

---

## Database

This project uses **SQLite** as the persistent data store and **GORM** as the ORM.

---

## Docker

The project includes a multi-stage Dockerfile for building a lightweight production image.

Build:

```bash
docker build -t ticket-system .
```

Run:

```bash
docker run --env-file .env -p 8080:8080 ticket-system
```

---

## Health Check

```
GET /health
```

Response

```json
{
    "status": "ok"
}
```

---

## Author

**Divyshekhar Sinha**