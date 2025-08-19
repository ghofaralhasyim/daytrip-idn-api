package models

import (
	"database/sql"
	"time"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type Invitation struct {
	Id          int64          `db:"id"`
	UserId      int64          `db:"user_id"`
	Title       string         `db:"title"`
	Slug        string         `db:"slug"`
	Description sql.NullString `db:"description"`
	TemplateId  sql.NullInt64  `db:"template_id"`
	EventTypeId sql.NullInt64  `db:"event_type_id"`
	StartDate   sql.NullTime   `db:"start_date"`
	EndDate     sql.NullTime   `db:"end_date"`
	MapsUrl     sql.NullString `db:"maps_url"`
	Address     sql.NullString `db:"address"`
	Location    sql.NullString `db:"location"`
	DressCode   sql.NullString `db:"dress_code"`
	CreatedAt   time.Time      `db:"created_at"`
}

func ToInvitationModel(e entities.InvitationEntity) Invitation {
	return Invitation{
		Id:          e.Id,
		Slug:        e.Slug,
		UserId:      e.UserId,
		Title:       e.Title,
		Description: helpers.NewNullString(e.Description),
		TemplateId:  helpers.NewNullInt64(e.TemplateId),
		StartDate:   helpers.NewNullTime(e.StartDate),
		EndDate:     helpers.NewNullTime(e.EndDate),
		MapsUrl:     helpers.NewNullString(e.MapsUrl),
		Address:     helpers.NewNullString(e.Address),
		Location:    helpers.NewNullString(e.Location),
		DressCode:   helpers.NewNullString(e.DressCode),
		CreatedAt:   e.CreatedAt,
	}
}
