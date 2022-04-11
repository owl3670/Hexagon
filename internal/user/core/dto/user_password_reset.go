package dto

type UserResetPassword struct {
	PhoneNumber     string `json:"phone_number"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	SMSToken        string `json:"sms_token"`
}
