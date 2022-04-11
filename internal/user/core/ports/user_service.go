package ports

import (
	"Hexagon/common/helper"
	"Hexagon/internal/auth"
	"Hexagon/internal/user/core/dto"
)

type UserService interface {
	SignUp(payload dto.UserSignUp) error
	Login(payload dto.UserLogin) (*auth.Token, error)
	GetUser(id int64) (helper.Envelope, error)
	ResetPassword(payload dto.UserResetPassword) error
	GenerateSMSCode(payload dto.UserGenerateSMSCode) (helper.Envelope, error)
	ConfirmSMSCode(payload dto.UserConfirmSMSCode) error
}
