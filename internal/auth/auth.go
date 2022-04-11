package auth

import (
	"Hexagon/common/errors"
	"Hexagon/internal/user/core/domains/entity"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"time"
)

const (
	BearerSchema string = "BEARER "
)

type AuthServer struct {
	Secret string
}

type Token struct {
	AccessToken string `json:"access_token""`
}

func GetAuthServer(secret string) *AuthServer {
	return &AuthServer{secret}
}

func (a *AuthServer) CreateToken(user entity.User) (*Token, error) {
	var err error

	claims := jwt.MapClaims{}
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken := &Token{}
	jwtToken.AccessToken, err = token.SignedString([]byte(a.Secret))
	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
}

func (a *AuthServer) ValidateToken(accessToken string) (int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.Secret), nil
	})
	if err != nil {
		return -1, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return -1, fmt.Errorf("invalid token")
	}

	userId := int64(payload["sub"].(float64))
	return userId, nil
}

func (a *AuthServer) GetTokenFromRequestHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization header required")
	}

	bearerLength := len(BearerSchema)
	if len(authHeader) <= bearerLength || strings.ToUpper(authHeader[0:bearerLength]) != BearerSchema {
		return "", fmt.Errorf("Authorization requires valid Bearer scheme")
	}

	return authHeader[bearerLength:], nil
}

func (a *AuthServer) UseTokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := a.GetTokenFromRequestHeader(r)
		if err != nil {
			errors.ForbiddenResponse(w, r, err)
			return
		}

		var userId int64
		userId, err = a.ValidateToken(token)
		if err != nil {
			errors.ForbiddenResponse(w, r, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
