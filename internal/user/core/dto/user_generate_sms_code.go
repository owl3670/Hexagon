package dto

type UserGenerateSMSCode struct {
	ConfirmationType string `json:"confirmation_type"`
	PhoneNumber      string `json:"phone_number"`
}
