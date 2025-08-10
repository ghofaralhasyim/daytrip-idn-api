package entities

import (
	"time"
)

type DestinationEntity struct {
	Id        int64
	Name      string
	Image     string
	CreatedAt time.Time
}

func MakeDestinationEntity(
	id int64,
	name string,
	image string,
	createdAt time.Time,
) *DestinationEntity {
	return &DestinationEntity{
		Id:        id,
		Name:      name,
		Image:     image,
		CreatedAt: createdAt,
	}
}
