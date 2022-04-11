package services

import (
	"Hexagon/common/helper"
	"Hexagon/internal/auth"
	"Hexagon/internal/user/core/domains/entity"
	smsError "Hexagon/internal/user/core/domains/errors/sms"
	userError "Hexagon/internal/user/core/domains/errors/user"
	"Hexagon/internal/user/core/dto"
	"Hexagon/internal/user/core/ports"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type userService struct {
	userRepository ports.UserRepository
	smsRepository  ports.SMSRepository
	authServer     *auth.AuthServer
}

func GetUserService(userRepo ports.UserRepository, smsRepo ports.SMSRepository, authServ *auth.AuthServer) ports.UserService {
	return &userService{
		userRepository: userRepo,
		smsRepository:  smsRepo,
		authServer:     authServ,
	}
}

func (u *userService) SignUp(payload dto.UserSignUp) error {
	sms, err := u.smsRepository.GetByToken(payload.SMSToken)
	if err != nil {
		return err
	}

	err = sms.CheckPhoneNumber(payload.PhoneNumber)
	if err != nil {
		return err
	}

	err = sms.CheckIsConfirmed()
	if err != nil {
		return err
	}

	user := entity.NewUser()

	encryptPassword := helper.EncryptSHA256(payload.Password)
	err = user.CreateUser(payload.Email, payload.Nickname, encryptPassword, payload.Name, payload.PhoneNumber)
	if err != nil {
		return err
	}

	encryptConfirmPassword := helper.EncryptSHA256(payload.Password)
	err = user.ConfirmPassword(encryptConfirmPassword)
	if err != nil {
		return err
	}

	err = u.userRepository.Save(*user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Login(payload dto.UserLogin) (*auth.Token, error) {
	var user *entity.User
	var err error
	if payload.Email != "" {
		user, err = u.userRepository.GetByEmail(payload.Email)
		if err != nil {
			return nil, err
		}
	} else if payload.Nickname != "" {
		user, err = u.userRepository.GetByEmail(payload.Email)
		if err != nil {
			return nil, err
		}
	} else if payload.PhoneNumber != "" {
		user, err = u.userRepository.GetByEmail(payload.Email)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, userError.UserIdentificationParamIsEmpty
	}

	encryptPassword := helper.EncryptSHA256(payload.Password)
	err = user.ConfirmPassword(encryptPassword)
	if err != nil {
		return nil, err
	}

	var token *auth.Token
	token, err = u.authServer.CreateToken(*user)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (u *userService) GetUser(id int64) (helper.Envelope, error) {
	user, err := u.userRepository.Get(id)
	if err != nil {
		return nil, err
	}

	data := helper.Envelope{
		"email":        user.Email,
		"nickname":     user.Nickname,
		"name":         user.Name,
		"phone_number": user.PhoneNumber,
	}

	return data, nil
}

func (u *userService) ResetPassword(payload dto.UserResetPassword) error {
	sms, err := u.smsRepository.GetByToken(payload.SMSToken)
	if err != nil {
		return err
	}

	err = sms.CheckPhoneNumber(payload.PhoneNumber)
	if err != nil {
		return err
	}

	err = sms.CheckIsConfirmed()
	if err != nil {
		return err
	}

	var user *entity.User
	user, err = u.userRepository.GetByPhoneNumber(sms.PhoneNumber)
	if err != nil {
		return err
	}

	encryptPassword := helper.EncryptSHA256(payload.NewPassword)
	err = user.ChangePassword(encryptPassword)
	if err != nil {
		return err
	}

	encryptConfirmPassword := helper.EncryptSHA256(payload.ConfirmPassword)
	err = user.ConfirmPassword(encryptConfirmPassword)
	if err != nil {
		return err
	}

	err = u.userRepository.Update(*user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GenerateSMSCode(payload dto.UserGenerateSMSCode) (helper.Envelope, error) {
	if payload.ConfirmationType == "reset_password" {
		_, err := u.userRepository.GetByPhoneNumber(payload.PhoneNumber)
		if err != nil {
			return nil, err
		}
	} else if payload.ConfirmationType == "sign_up" {
		_, err := u.userRepository.GetByPhoneNumber(payload.PhoneNumber)
		if err != nil {
			if !errors.Is(err, userError.UserNotFound) {
				return nil, err
			}
		} else {
			return nil, userError.UserPhoneNumberAlreadyExists
		}
	} else {
		return nil, smsError.SMSConfirmationTypeIsInvalid
	}

	sms := entity.NewSMS()
	uuid := helper.NewUUID()
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	codeNum := r.Intn(900000)
	codeNum += 100000
	code := fmt.Sprintf("%06d", codeNum)
	err := sms.CreateSMS(uuid.String(), payload.PhoneNumber, code)
	if err != nil {
		return nil, err
	}

	err = u.smsRepository.Save(*sms)
	if err != nil {
		return nil, err
	}

	// SMS 실제 발송 대신 response 에 code 포함 (TEST 편의 목적)
	return helper.Envelope{"token": sms.Token, "code": sms.Code}, nil
}

func (u *userService) ConfirmSMSCode(payload dto.UserConfirmSMSCode) error {
	sms, err := u.smsRepository.GetByToken(payload.Token)
	if err != nil {
		return err
	}

	err = sms.ConfirmSMS(payload.Code)
	if err != nil {
		return err
	}

	err = u.smsRepository.Save(*sms)
	if err != nil {
		return err
	}

	return nil
}
