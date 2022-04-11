package dto

type UserLogin struct {
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
