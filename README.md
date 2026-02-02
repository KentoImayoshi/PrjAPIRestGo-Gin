# PrjAPIRestGo-Gin

A REST API built with **Go** using **Gin + GORM** and **PostgreSQL**.

- ✅ **Students** CRUD
- ✅ Run locally (Go) with Postgres on Docker
- ✅ Run everything with Docker Compose *(optional)*
- ✅ Environment-based configuration via `.env`

---

## Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- Docker / Docker Compose

---

## Getting Started

### Option A (recommended) — Postgres on Docker + API locally

## 1) Start the database with Docker:

docker compose up -d

## 2) Create a .env file at the project root (see example below)

## 3) Run the API:
go run main.go
### API: http://localhost:8080

### Option B — Everything with Docker (API + Postgres)
#### Use this option if you have a Dockerfile at the project root (we can add it next).

docker compose up --build
API: http://localhost:8080

### Environment Variables (.env)
Create a .env file at the project root:

* Postgres (container)
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_DB=root
POSTGRES_PORT=5432
POSTGRES_HOST=postgres

* App (GORM)
* If the API runs locally (outside Docker), use localhost:
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=root
DB_SSLMODE=disable
DB_TIMEZONE=America/Sao_Paulo
Quick tip:

DB_HOST=localhost → API running locally

DB_HOST=postgres → API running inside Docker Compose

## API Endpoints (Students CRUD)
Base URL: http://localhost:8080

### List students
curl -X GET http://localhost:8080/alunos

### Get student by ID
curl -X GET http://localhost:8080/alunos/1

### Create student
curl -X POST http://localhost:8080/alunos \
  -H "Content-Type: application/json" \
  -d '{"nome":"John Doe","cpf":"12345678900","rg":"1234567"}'
  
### Update student
curl -X PUT http://localhost:8080/alunos/1 \
  -H "Content-Type: application/json" \
  -d '{"nome":"John Updated","cpf":"12345678900","rg":"7654321"}'
  
### Delete student
curl -X DELETE http://localhost:8080/alunos/1


## Database
This project uses GORM and runs AutoMigrate at startup to create/update the Student table based on the model.

## 🧩 Project Structure (suggested)
```text
.
├── database/
│   └── db.go              # Postgres connection (GORM + AutoMigrate)
├── models/
│   └── aluno.go           # Student model (Aluno)
├── controllers/
│   └── alunos.go          # Handlers (CRUD)
├── routes/
│   └── routes.go          # Gin routes registration
├── main.go                # App bootstrap
├── docker-compose.yml
├── .env                   # (gitignored)
└── .gitignore
```

## 🧯 Troubleshooting

- **Error `lookup postgres: no such host` when running locally**
  - If the API is running **outside Docker**, set `DB_HOST=localhost` in your `.env`.
  - `postgres` is the service name/hostname **inside** the Docker Compose network.

- **Connected, but authentication failed**
  - Double-check `POSTGRES_USER`, `POSTGRES_PASSWORD`, and `POSTGRES_DB` in the container and the `DB_*` variables used by the API.

- **Database “disappeared” after recreating the container**
  - Make sure you’re using a volume (e.g. `postgres-data:/var/lib/postgresql/data`) — usually already defined in `docker-compose.yml`.

---

## 📜 License

This project is licensed under the **MIT License**. See `LICENSE`.
