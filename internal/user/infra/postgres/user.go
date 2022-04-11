package postgres

import (
	"Hexagon/internal/user/core/domains/entity"
	userError "Hexagon/internal/user/core/domains/errors/user"
	"Hexagon/internal/user/core/ports"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func GetUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) get(columnName string, param interface{}) (*entity.User, error) {
	var user entity.User

	query := fmt.Sprintf(`SELECT * FROM users WHERE %s = $1`, columnName)
	err := r.db.QueryRow(query, param).Scan(
		&user.ID,
		&user.Email,
		&user.Nickname,
		&user.Password,
		&user.Name,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, userError.UserNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r *userRepository) Get(id int64) (*entity.User, error) {
	return r.get("id", id)
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	return r.get("email", email)
}

func (r *userRepository) GetByNickname(nickname string) (*entity.User, error) {
	return r.get("nickname", nickname)
}

func (r *userRepository) GetByPhoneNumber(phoneNumber string) (*entity.User, error) {
	return r.get("phone_number", phoneNumber)
}

func (r *userRepository) Save(user entity.User) error {
	query := `INSERT INTO users (email, nickname, name, password, phone_number)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	args := []interface{}{
		user.Email,
		user.Nickname,
		user.Name,
		user.Password,
		user.PhoneNumber,
	}

	err := r.db.QueryRow(query, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "violates unique"):
			{
				errText := err.Error()
				switch {
				case strings.Contains(errText, "email"):
					return userError.UserEmailAlreadyExists
				case strings.Contains(errText, "nickname"):
					return userError.UserNicknameAlreadyExists
				case strings.Contains(errText, "phone_number"):
					return userError.UserPhoneNumberAlreadyExists
				default:
					return err
				}
			}
		default:
			return err
		}
	}

	return nil
}

func (r *userRepository) Update(user entity.User) error {
	query := `UPDATE users SET password = $1
	WHERE id = $2
	RETURNING id`

	args := []interface{}{
		user.Password,
		user.ID,
	}

	err := r.db.QueryRow(query, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return userError.UserNotFound
		default:
			return err
		}
	}

	return nil
}
