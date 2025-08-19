package models

import "time"

type InvitationTemplateAsset struct {
	Id         int64     `db:"id" json:"id"`
	TemplateId int64     `db:"template_id" json:"template_id"`
	AssetType  string    `db:"asset_type" json:"asset_type"`
	AssetUrl   string    `db:"asset_url" json:"asset_url"`
	SortOrder  int       `db:"sort_order" json:"sort_order"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
