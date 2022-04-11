package dto

type UserConfirmSMSCode struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}
