package main

import (
	"Hexagon/config"
	"Hexagon/internal/db/postgres"
	"Hexagon/internal/db/redis"
	"Hexagon/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal(err)
	}

	var cfg = config.Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      4000,
		Postgres: config.PostgresConfig{
			DSN:         os.Getenv("POSTGRES_DSN"),
			MaxIdleConn: 25,
			MaxOpenConn: 25,
			MaxIdleTime: "15m",
		},
		Redis: config.RedisConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	db, err := postgres.OpenDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	redis := redis.OpenDB(cfg)

	server := server.GetServer(cfg, db, redis, logger)
	if err = server.Run(); err != nil {
		logger.Fatal(err)
	}
}
