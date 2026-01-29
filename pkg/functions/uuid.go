package functions

import (
	"github.com/google/uuid"
)

type UUIDFuncs struct{}

func (UUIDFuncs) IsValid(in string) bool {
	_, err := uuid.Parse(in)

	return err == nil
}

func (UUIDFuncs) Nil() string {
	return uuid.Nil.String()
}

func (UUIDFuncs) V1() (string, error) {
	return UUIDToString(uuid.NewUUID())
}

func (UUIDFuncs) V4() (string, error) {
	return UUIDToString(uuid.NewRandom())
}

func (UUIDFuncs) V6() (string, error) {
	return UUIDToString(uuid.NewV6())
}

func (UUIDFuncs) V7() (string, error) {
	return UUIDToString(uuid.NewV7())
}

func UUIDToString(in uuid.UUID, err error) (string, error) {
	if err != nil {
		return "", err
	}

	return in.String(), err
}
