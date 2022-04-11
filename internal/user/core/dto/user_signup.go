package dto

type UserSignUp struct {
	Email           string `json:"email"`
	Nickname        string `json:"nickname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	PhoneNumber     string `json:"phone_number"`
	SMSToken        string `json:"sms_token"`
}
