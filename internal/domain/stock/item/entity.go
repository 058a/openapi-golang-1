package item

import "github.com/google/uuid"

type Id uuid.UUID

func (e Id) ToUuid() uuid.UUID {
	return uuid.UUID(e)
}

func (e Id) ToString() string {
	return uuid.UUID(e).String()
}
