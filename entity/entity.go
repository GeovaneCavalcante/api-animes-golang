package entity

import "github.com/google/uuid"

// ID entity ID
type ID = uuid.UUID

type IDHash = string

// NewID create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

// StringToID convert a string to an entity id
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return id, err
}
