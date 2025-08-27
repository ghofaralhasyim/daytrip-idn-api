package models

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type Invitation struct {
	Id          int64     `db:"id"`
	Title       string    `db:"title"`
	Slug        string    `db:"slug"`
	Description string    `db:"description"`
	TemplateId  int64     `db:"template_id"`
	EventTypeId int64     `db:"event_type_id"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	MapsUrl     string    `db:"maps_url"`
	Address     string    `db:"address"`
	Location    string    `db:"location"`
	DressCode   string    `db:"dress_code"`
	CreatedAt   time.Time `db:"created_at"`
	Image       string    `db:"image"`
	Image1      string    `db:"image1"`
	KeyPass     string    `db:"keyPass"`
	BirthdayVal int64     `db:"birthday_val"`
}

func ToInvitationModel(e entities.InvitationEntity) Invitation {
	return Invitation{
		Id:          e.Id,
		Slug:        e.Slug,
		Title:       e.Title,
		Description: *e.Description,
		TemplateId:  *e.TemplateId,
		StartDate:   *e.StartDate,
		EndDate:     *e.EndDate,
		MapsUrl:     *e.MapsUrl,
		Address:     *e.Address,
		Location:    *e.Location,
		DressCode:   *e.DressCode,
		CreatedAt:   e.CreatedAt,
		Image:       *e.Image,
		Image1:      *e.Image1,
		KeyPass:     e.KeyPass,
		BirthdayVal: *e.BirthdayVal,
	}
}
