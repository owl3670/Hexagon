package http

import (
	httpError "Hexagon/common/errors"
	"Hexagon/common/helper"
	"Hexagon/internal/auth"
	smsError "Hexagon/internal/user/core/domains/errors/sms"
	userError "Hexagon/internal/user/core/domains/errors/user"
	"Hexagon/internal/user/core/dto"
	"Hexagon/internal/user/core/ports"
	"errors"
	"net/http"
)

type UserHandler struct {
	userService ports.UserService
}

func GetUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserSignUp
	err := helper.ReadJSON(w, r, &payload)
	if err != nil {
		httpError.BadRequestResponse(w, r, err)
		return
	}

	err = h.userService.SignUp(payload)
	if err != nil {
		var userError userError.UserError
		var smsError smsError.SMSTokenError

		switch {
		case errors.As(err, &userError):
			httpError.BadRequestResponse(w, r, err)
		case errors.As(err, &smsError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	env := helper.Envelope{"email": payload.Email, "nickname": payload.Nickname, "name": payload.Name, "phone_number": payload.PhoneNumber}
	err = helper.WriteJSON(w, http.StatusCreated, env, nil)
	if err != nil {
		httpError.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserLogin
	err := helper.ReadJSON(w, r, &payload)
	if err != nil {
		httpError.BadRequestResponse(w, r, err)
		return
	}

	var token *auth.Token
	token, err = h.userService.Login(payload)

	if err != nil {
		var userError userError.UserError
		switch {
		case errors.As(err, &userError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	env := helper.Envelope{"jwt_token": token}
	err = helper.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		httpError.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *UserHandler) MeHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int64)

	if !ok {
		httpError.ErrorResponse(w, r, http.StatusBadRequest, "user is invalid")
		return
	}

	userData, err := h.userService.GetUser(userId)
	if err != nil {
		var userError userError.UserError
		switch {
		case errors.As(err, &userError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, userData, nil)
	if err != nil {
		httpError.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *UserHandler) SMSHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserGenerateSMSCode
	err := helper.ReadJSON(w, r, &payload)
	if err != nil {
		httpError.BadRequestResponse(w, r, err)
		return
	}

	var data helper.Envelope
	data, err = h.userService.GenerateSMSCode(payload)
	if err != nil {
		var userError userError.UserError
		var smsTokenError smsError.SMSTokenError
		switch {
		case errors.As(err, &userError):
			httpError.BadRequestResponse(w, r, err)
		case errors.As(err, &smsTokenError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	err = helper.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		httpError.InternalServerErrorResponse(w, r, err)
		return
	}
}

func (h *UserHandler) SMSConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserConfirmSMSCode
	err := helper.ReadJSON(w, r, &payload)
	if err != nil {
		httpError.BadRequestResponse(w, r, err)
		return
	}

	err = h.userService.ConfirmSMSCode(payload)
	if err != nil {
		var smsTokenError smsError.SMSTokenError
		switch {
		case errors.As(err, &smsTokenError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) PasswordChangeHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserResetPassword
	err := helper.ReadJSON(w, r, &payload)
	if err != nil {
		httpError.BadRequestResponse(w, r, err)
		return
	}

	err = h.userService.ResetPassword(payload)
	if err != nil {
		var smsTokenError smsError.SMSTokenError
		switch {
		case errors.As(err, &smsTokenError):
			httpError.BadRequestResponse(w, r, err)
		default:
			httpError.InternalServerErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
