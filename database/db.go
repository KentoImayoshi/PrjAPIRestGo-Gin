package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kentoimayoshi/PrjAPIRestGo-Gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	_ = godotenv.Load()

	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		open(dsn)
		return
	}

	host := getEnv("DB_HOST", "localhost")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")
	timezone := getEnv("DB_TIMEZONE", "America/Sao_Paulo")

	if user == "" || name == "" {
		log.Panic("Variáveis de ambiente do banco ausentes (DB_USER/DB_NAME).")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, pass, name, port, sslmode, timezone,
	)
	open(dsn)
}

func open(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panic("Erro no AutoMigrate")
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
