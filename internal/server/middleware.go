package server

import (
	"Hexagon/common/errors"
	"fmt"
	"net/http"
)

func (s *Server) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.logger.Println("Recover panic")
				s.logger.Println(err)
				w.Header().Set("Connection", "close")
				errors.InternalServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
