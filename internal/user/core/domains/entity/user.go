package entity

import (
	userError "Hexagon/internal/user/core/domains/errors/user"
	"database/sql"
	"net/mail"
	"time"
)

type User struct {
	ID          int    `db:"id"`
	Email       string `db:"email"`
	Nickname    string `db:"nickname"`
	Password    string `db:"password"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) CreateUser(email, nickname, password, name, phoneNumber string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return userError.EmailInvalid
	}

	if nickname == "" {
		return userError.NicknameIsEmpty
	}

	if password == "" {
		return userError.PasswordIsEmpty
	}

	if name == "" {
		return userError.NameIsEmpty
	}

	if phoneNumber == "" {
		return userError.PhoneNumberIsEmpty
	}

	u.Email = email
	u.Nickname = nickname
	u.Password = password
	u.Name = name
	u.PhoneNumber = phoneNumber

	return nil
}

func (u *User) ChangePassword(password string) error {
	u.Password = password
	
	return nil
}

func (u *User) ConfirmPassword(password string) error {
	if u.Password != password {
		return userError.ConfirmPasswordIsInvalid
	}

	return nil
}
