package item

import "github.com/google/uuid"

type Id uuid.UUID

func (v Id) ToUuid() uuid.UUID {
	return uuid.UUID(v)
}

func (v Id) ToString() string {
	return uuid.UUID(v).String()
}
