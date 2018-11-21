package util

import "github.com/google/uuid"

type TransID struct {
	uuid  uuid.UUID
}

func Get() *TransID {
	id := new(TransID)
	id.uuid = uuid.New()
	return id
}
