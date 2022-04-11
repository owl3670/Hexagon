package server

import (
	"Hexagon/common/errors"
	"Hexagon/internal/auth"
	userHttp "Hexagon/internal/user/handlers/http"
	"Hexagon/internal/user/infra/postgres"
	"Hexagon/internal/user/infra/redis"
	"Hexagon/internal/user/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *Server) GetHandlers() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(errors.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(errors.MethodNotAllowedResponse)

	// AuthServer
	authServer := auth.GetAuthServer(s.config.JWTSecret)

	// User
	userRepo := postgres.GetUserRepository(s.db)
	smsRepo := redis.GetSMSRepository(s.redisClient)
	userService := services.GetUserService(userRepo, smsRepo, authServer)
	userHandler := userHttp.GetUserHandler(userService)
	router.HandlerFunc(http.MethodPost, "/v1/users", userHandler.SignUpHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", userHandler.LoginHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/me", authServer.UseTokenValidation(userHandler.MeHandler))
	router.HandlerFunc(http.MethodPost, "/v1/users/sms", userHandler.SMSHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/sms/confirm", userHandler.SMSConfirmationHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/password", userHandler.PasswordChangeHandler)

	return s.recoverPanic(router)
}
