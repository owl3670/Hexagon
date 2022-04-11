package redis

import (
	"Hexagon/internal/user/core/domains/entity"
	smsError "Hexagon/internal/user/core/domains/errors/sms"
	"Hexagon/internal/user/core/ports"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

func GetSMSRepository(client *redis.Client) ports.SMSRepository {
	return &smsRepository{client: client}
}

type smsRepository struct {
	client *redis.Client
}

func (s *smsRepository) GetByToken(token string) (*entity.SMS, error) {
	ctx := context.Background()

	var sms entity.SMS
	err := s.client.Get(ctx, token).Scan(&sms)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, smsError.SMSTokenNotFound
		}
		return nil, err
	}
	return &sms, nil
}

func (s *smsRepository) Save(sms entity.SMS) error {
	ctx := context.Background()

	err := s.client.Set(ctx, sms.Token, sms, time.Minute*30).Err()
	if err != nil {
		return err
	}

	return nil
}
