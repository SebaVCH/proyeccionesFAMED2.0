package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func StartRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error al conectar con Redis: %v", err)
		return err
	}

	log.Println("Conectado exitosamente con Redis")
	return nil
}

func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Printf("Error al conectar con Redis: %v", err)
	}
	log.Println("Conexion con Redis cerrada exitosamente")
}
