package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var JwtSecret []byte

func LoadENV() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el .env")
	}

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))

	return nil
}
