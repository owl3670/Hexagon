package ports

import (
	"Hexagon/internal/user/core/domains/entity"
)

type UserRepository interface {
	Get(id int64) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByNickname(nickname string) (*entity.User, error)
	GetByPhoneNumber(phoneNumber string) (*entity.User, error)
	Save(entity.User) error
	Update(entity.User) error
}
