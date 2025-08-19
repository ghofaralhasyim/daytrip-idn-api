package models

import "github.com/daytrip-idn-api/internal/entities"

type Activity struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Image string `db:"image"`
}

func ToActivityModel(
	e entities.ActivityEntity,
) Activity {
	return Activity{
		Id:    e.Id,
		Name:  e.Name,
		Image: e.Image,
	}
}
