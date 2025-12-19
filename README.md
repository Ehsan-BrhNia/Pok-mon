# Pokémon Battle 

This is a Go project that provides a Pokémon Battle API using the **Gin framework** and **GORM** for database interactions. The project is fully **containerized with Docker** and orchestrated with **Docker Compose**, including a PostgreSQL database.

---

## Features

- List available Pokémon heroes (`GET /heroes`)
- Select heroes for a battle (`POST /select`)
- Perform a battle simulation (`GET /battle`)
- Fully containerized for easy deployment

---

## Technologies Used

- [Go](https://golang.org/) 1.25
- [Gin](https://github.com/gin-gonic/gin) Web Framework
- [GORM](https://gorm.io/) ORM
- [PostgreSQL](https://www.postgresql.org/) Database
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## Project Structure
pokemon/
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── battleService/
├── database/
└── README.md



- `main.go` – Entry point of the application  
- `battleService/` – API handlers for battle logic  
- `database/` – Database initialization and connection logic  

---

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed
- [Docker Compose](https://docs.docker.com/compose/install/) installed

---

## Getting Started

1. **Clone the repository**

```bash
git clone https://github.com/Ehsan-BrhNia/Pok-mon.git
cd Pok-mon

2. docker-compose up --build

