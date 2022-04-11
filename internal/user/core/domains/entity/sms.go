package entity

import (
	"Hexagon/internal/user/core/domains/errors/sms"
	"encoding/json"
)

type SMS struct {
	Token       string `json:"token"`
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
	Confirmed   bool   `json:"confirmed"`
}

func NewSMS() *SMS {
	return &SMS{}
}

func (s *SMS) CreateSMS(token, phoneNumber, code string) error {
	s.Token = token
	s.PhoneNumber = phoneNumber
	s.Code = code
	s.Confirmed = false

	return nil
}

func (s *SMS) ConfirmSMS(code string) error {
	if s.Confirmed {
		return sms.SMSIsAlreadyConfirmed
	}

	if s.Code != code {
		return sms.SMSCodeIsInvalid
	}
	s.Confirmed = true

	return nil
}

func (s *SMS) CheckIsConfirmed() error {
	if !s.Confirmed {
		return sms.SMSIsNotConfirmed
	}

	return nil
}

func (s *SMS) CheckPhoneNumber(phoneNumber string) error {
	if s.PhoneNumber != phoneNumber {
		return sms.PhoneNumberIsMismatch
	}

	return nil
}

func (s SMS) MarshalBinary() ([]byte, error) {
	bytes, err := json.Marshal(s)
	return bytes, err
}

func (s *SMS) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
