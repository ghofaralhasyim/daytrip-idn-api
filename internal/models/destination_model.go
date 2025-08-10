package models

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type Destination struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
}

func ToDestinationModel(
	e entities.DestinationEntity,
) Destination {
	return Destination{
		Id:        e.Id,
		Name:      e.Name,
		Image:     e.Image,
		CreatedAt: e.CreatedAt,
	}
}
