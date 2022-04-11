package ports

import "Hexagon/internal/user/core/domains/entity"

type SMSRepository interface {
	GetByToken(token string) (*entity.SMS, error)
	Save(sms entity.SMS) error
}
