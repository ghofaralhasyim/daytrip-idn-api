package models

type InvitationAsset struct {
	Id           int64  `db:"id"`
	InvitationId int64  `db:"invitation_Id"`
	AssetUrl     string `db:"asset_url"`
	SortOrder    int64  `db:"sort_order"`
}
