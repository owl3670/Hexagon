package errors

import (
	"Hexagon/common/helper"
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := helper.Envelope{"error": message}

	err := helper.WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusNotFound, "The requested resource could not be found")
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusBadRequest, "The requested method is not supported for this resource")
}

func InternalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err)
	ErrorResponse(w, r, http.StatusInternalServerError, "internal server error")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func ForbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusForbidden, err.Error())
}
