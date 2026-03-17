package entity

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewUUID() ID {
	return ID(uuid.New())
}

func ParseID(str string) (ID, error) {
	id, err := uuid.Parse(str)
	return id, err
}
