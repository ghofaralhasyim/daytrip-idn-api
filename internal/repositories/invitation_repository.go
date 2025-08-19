package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type (
	InvitationRepository interface {
		Create(ctx context.Context, invitation *entities.InvitationEntity, assets []entities.InvitationAssetEntity) (int64, error)
	}
	invitationRepository struct {
		db *sql.DB
	}
)

func NewInvitationRepository(db *sql.DB) InvitationRepository {
	return &invitationRepository{
		db: db,
	}
}

func (r *invitationRepository) Create(ctx context.Context, invitation *entities.InvitationEntity, assets []entities.InvitationAssetEntity) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dataModel := models.ToInvitationModel(*invitation)

	query := `
		INSERT INTO invitations (
			user_id,
			slug,
			title,
			description,
			template_id,
			start_date,
			end_date,
			maps_url,
			address,
			location,
			dress_code,
			created_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
		) RETURNING id
	`

	var invitationID int64
	err = tx.QueryRowContext(ctx,
		query,
		dataModel.UserId,
		dataModel.Slug,
		dataModel.Title,
		dataModel.Description,
		dataModel.TemplateId,
		dataModel.StartDate,
		dataModel.EndDate,
		dataModel.MapsUrl,
		dataModel.Address,
		dataModel.Location,
		dataModel.DressCode,
		time.Now(),
	).Scan(&invitationID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert invitation: %w", err)
	}

	if len(assets) > 0 {
		assetQuery := `
			INSERT INTO invitation_assets (invitation_id, asset_url, sort_order, created_at)
			VALUES ($1, $2, $3, $4)
		`

		for _, asset := range assets {
			_, err = tx.ExecContext(ctx,
				assetQuery,
				invitationID,
				asset.AssetUrl,
				asset.SortOrder,
				time.Now(),
			)
			if err != nil {
				return 0, fmt.Errorf("failed to insert asset: %w", err)
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return invitationID, nil
}

func (r *invitationRepository) GetBySlug(
	slug string,
) (*entities.InvitationEntity, error) {
	query := `
		SELECT 
			id, user_id, title, slug, description, template_id,
			start_date, end_date, maps_url, address, location,
			dress_code, created_at
		FROM invitations
		WHERE slug = $1
	`

	row := r.db.QueryRow(query, slug)

	result, err := helpers.ScanRowToStruct[models.Invitation](row, nil)
	if err != nil {
		return nil, err
	}

	entity := entities.MakeInvitationEntity(
		result.Id,
		result.UserId,
		result.Slug,
		result.Title,
		helpers.NullString(result.Description),
		helpers.NullInt64(result.TemplateId),
		helpers.NullTime(result.StartDate),
		helpers.NullTime(result.EndDate),
		helpers.NullString(result.MapsUrl),
		helpers.NullString(result.Address),
		helpers.NullString(result.Location),
		helpers.NullString(result.DressCode),
		result.CreatedAt,
	)

	return &entity, nil
}
