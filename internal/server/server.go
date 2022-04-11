package server

import (
	"Hexagon/config"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config      config.Config
	db          *sql.DB
	redisClient *redis.Client
	logger      *log.Logger
}

func GetServer(cfg config.Config, db *sql.DB, redisClient *redis.Client, logger *log.Logger) *Server {
	return &Server{
		config:      cfg,
		db:          db,
		redisClient: redisClient,
		logger:      logger,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.config.Port),
		Handler:      s.GetHandlers(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	s.logger.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	s.logger.Fatal(err)
	return err
}
