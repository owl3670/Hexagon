package helper

import "github.com/google/uuid"

type UUID struct {
	value uuid.UUID
}

func NewUUID() *UUID {
	return &UUID{
		value: uuid.New(),
	}
}

func (uuid *UUID) String() string {
	return uuid.value.String()
}

func TryParse(str string) (*UUID, error) {
	uuid, err := uuid.Parse(str)
	if err != nil {
		return nil, err
	}

	return &UUID{uuid}, nil
}
