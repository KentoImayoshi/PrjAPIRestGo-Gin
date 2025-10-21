# PrjAPIRestGo-Gin

API REST em **Go** usando **Gin** + **GORM** com **PostgreSQL**.  
Suporta execução **local** ou via **Docker Compose**, utilizando variáveis de ambiente (sem credenciais no código).

## 🧰 Stack
- Go (Gin, GORM)
- PostgreSQL
- Docker & Docker Compose
- Variáveis de ambiente via '.env' (dev) ou ambiente (prod)

---

## 🚀 Começando

### 1) Pré-requisitos
- Go 1.21+ (recomendado)
- Docker & Docker Compose

### 2) Clone o projeto

git clone https://github.com/KentoImayoshi/PrjAPIRestGo-Gin.git
cd PrjAPIRestGo-Gin

3) Crie seu .env (não comite)
Crie um arquivo .env na raiz com os valores do seu ambiente:

Arquivo .env

# Postgres
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_DB=root
POSTGRES_PORT=5432
POSTGRES_HOST=postgres

# App (conexão do GORM)
DB_HOST=localhost          # use 'localhost' se rodar a API fora do Docker
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=root
DB_SSLMODE=disable
DB_TIMEZONE=America/Sao_Paulo
(Opcional) DATABASE_URL=postgres://root:root@localhost:5432/root?sslmode=disable

🧪 Rodando
Opção A — Postgres no Docker + API local
Suba só o banco:

docker compose up -d postgres

Rode a API local (usa DB_HOST=localhost):

go run main.go

Opção B — Tudo no Docker
O docker-compose.yml já injeta DB_HOST=postgres para o container da API:

docker compose up --build

A API ficará acessível em http://localhost:8080 (se você expôs a porta no compose/main).


🗄️ Banco de Dados & Migrações
Na inicialização, o GORM executa AutoMigrate para o modelo Aluno.
Isso cria/atualiza a tabela automaticamente.

Modelo (referência típica do curso/gin):

type Aluno struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Nome string `json:"nome"`
    CPF  string `json:"cpf"`
    RG   string `json:"rg"`
}

Se seu modelo tiver campos diferentes, ajuste aqui e nos exemplos abaixo.

📚 Endpoints (CRUD Alunos)
Base URL: http://localhost:8080

GET /alunos
Lista todos os alunos.
curl -X GET http://localhost:8080/alunos

GET /alunos/{id}
Busca um aluno por ID.
curl -X GET http://localhost:8080/alunos/1

POST /alunos
Cria um aluno.
curl -X POST http://localhost:8080/alunos \
  -H "Content-Type: application/json" \
  -d '{"nome":"João da Silva","cpf":"12345678900","rg":"1234567"}'

PUT /alunos/{id}
Atualiza um aluno.
curl -X PUT http://localhost:8080/alunos/1 \
  -H "Content-Type: application/json" \
  -d '{"nome":"João Silva","cpf":"12345678900","rg":"7654321"}'

DELETE /alunos/{id}
Remove um aluno.
curl -X DELETE http://localhost:8080/alunos/1

Rotas reais podem variar conforme seu main.go/controllers. Se quiser, mando a seção de rotas exatamente como estão no seu código.

⚙️ Configuração
Variáveis de ambiente usadas
DATABASE_URL (opcional, tem prioridade se definido)

DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE, DB_TIMEZONE

Dica de segurança
Não comite o .env.

Para produção, defina as variáveis no ambiente do servidor, ou use Docker secrets.

🧩 Estrutura (sugestão)
.
├─ database/
│  └─ db.go              # Conexão com o Postgres (GORM + AutoMigrate)
├─ models/
│  └─ aluno.go           # Model Aluno
├─ controllers/
│  └─ alunos.go          # Handlers (CRUD)
├─ routes/
│  └─ routes.go          # Registro de rotas Gin
├─ main.go               # Boot da aplicação
├─ docker-compose.yml
├─ Dockerfile
├─ .env                  # (ignorado no Git)
└─ .gitignore


🧯 Troubleshooting
Erro “lookup postgres: no such host” rodando local:
Use DB_HOST=localhost no .env quando a API roda fora do Docker.
postgres é o nome do serviço dentro da rede do Docker.

Conectou mas deu erro de autenticação:
Confira POSTGRES_USER/POSTGRES_PASSWORD/POSTGRES_DB no container e DB_* da API.

Banco “sumiu” após recriar container:
Monte um volume (ex.: postgres-data:/var/lib/postgresql/data) — já previsto no compose.

📜 Licença
Este projeto é distribuído sob a licença MIT. Veja LICENSE (se aplicável).
