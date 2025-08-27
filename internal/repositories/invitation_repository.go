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
		Create(ctx context.Context, invitation *entities.InvitationEntity) (int64, error)
		GetInvitations(ctx context.Context) ([]entities.InvitationEntity, error)
		GetBySlug(ctx context.Context, slug string) (*entities.InvitationEntity, error)
		GetById(ctx context.Context, id int) (*entities.InvitationEntity, error)
		Update(ctx context.Context, invitation *entities.InvitationEntity) error
		DeleteInvitation(ctx context.Context, tx *sql.Tx, id int) error
		DeleteResponseInvitation(ctx context.Context, tx *sql.Tx, id int) error
		BeginTx() (*sql.Tx, error)
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

func (r *invitationRepository) Create(ctx context.Context, invitation *entities.InvitationEntity) (int64, error) {
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
			created_at,
			image,
			image1,
			"keyPass",
			birthday_val
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15
		) RETURNING id
	`

	var invitationID int64
	err = tx.QueryRowContext(ctx,
		query,
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
		dataModel.Image,
		dataModel.Image1,
		dataModel.KeyPass,
		dataModel.BirthdayVal,
	).Scan(&invitationID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert invitation: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return invitationID, nil
}

func (r *invitationRepository) Update(ctx context.Context, invitation *entities.InvitationEntity) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dataModel := models.ToInvitationModel(*invitation)

	query := `
		UPDATE invitations
		SET 
			slug = $1,
			title = $2,
			description = $3,
			template_id = $4,
			start_date = $5,
			end_date = $6,	
			maps_url = $7,
			address = $8,
			location = $9,
			dress_code = $10,
			image = $11,
			image1 = $12,
			"keyPass" = $13,
			birthday_val = $14
		WHERE id = $15
	`

	_, err = tx.ExecContext(ctx,
		query,
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
		dataModel.Image,
		dataModel.Image1,
		dataModel.KeyPass,
		dataModel.BirthdayVal,
		dataModel.Id, // WHERE id
	)
	if err != nil {
		return fmt.Errorf("failed to update invitation: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *invitationRepository) GetBySlug(
	ctx context.Context, slug string,
) (*entities.InvitationEntity, error) {

	columns := []string{
		"id", "title", "slug", "description", "template_id",
		"start_date", "end_date", "maps_url", "address", "location",
		"dress_code", "created_at", "image", "image1", "keyPass", "birthday_val",
	}

	query := `
		SELECT 
			id, title, slug, description, template_id,
			start_date, end_date, maps_url, address, location,
			dress_code, created_at, image, image1, "keyPass", birthday_val
		FROM invitations
		WHERE slug = $1;
	`

	row := r.db.QueryRowContext(ctx, query, slug)

	result, err := helpers.ScanRowToStruct[models.Invitation](row, columns)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	entity := entities.MakeInvitationEntity(
		result.Id,
		result.Slug,
		result.Title,
		&result.Description,
		&result.TemplateId,
		&result.StartDate,
		&result.EndDate,
		&result.MapsUrl,
		&result.Address,
		&result.Location,
		&result.DressCode,
		result.CreatedAt,
		&result.Image,
		&result.Image1,
		result.KeyPass,
		&result.BirthdayVal,
	)

	return &entity, nil
}

func (r *invitationRepository) GetById(
	ctx context.Context, id int,
) (*entities.InvitationEntity, error) {

	columns := []string{
		"id", "title", "slug", "description", "template_id",
		"start_date", "end_date", "maps_url", "address", "location",
		"dress_code", "created_at", "image", "image1", "keyPass", "birthday_val",
	}

	query := `
		SELECT 
			id, title, slug, description, template_id,
			start_date, end_date, maps_url, address, location,
			dress_code, created_at, image, image1, "keyPass", birthday_val
		FROM invitations
		WHERE id = $1;
	`

	row := r.db.QueryRowContext(ctx, query, id)

	result, err := helpers.ScanRowToStruct[models.Invitation](row, columns)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	entity := entities.MakeInvitationEntity(
		result.Id,
		result.Slug,
		result.Title,
		&result.Description,
		&result.TemplateId,
		&result.StartDate,
		&result.EndDate,
		&result.MapsUrl,
		&result.Address,
		&result.Location,
		&result.DressCode,
		result.CreatedAt,
		&result.Image,
		&result.Image1,
		result.KeyPass,
		&result.BirthdayVal,
	)

	return &entity, nil
}

func (r *invitationRepository) GetInvitations(ctx context.Context) (
	[]entities.InvitationEntity, error,
) {
	query := `
		SELECT 
			id, title, slug, description, template_id,
			start_date, end_date, maps_url, address, location,
			dress_code, created_at, image, image1, "keyPass", birthday_val
		FROM invitations;
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := helpers.ScanRowsToStructs[models.Invitation](rows)
	if err != nil {
		return nil, err
	}

	invitations := make([]entities.InvitationEntity, 0)
	for _, item := range result {
		entity := entities.MakeInvitationEntity(
			item.Id, item.Slug, item.Title,
			&item.Description, &item.TemplateId,
			&item.StartDate, &item.EndDate,
			&item.MapsUrl, &item.Address,
			&item.Location, &item.DressCode,
			item.CreatedAt, &item.Image, &item.Image1,
			item.KeyPass, &item.BirthdayVal,
		)
		invitations = append(invitations, entity)
	}

	return invitations, nil
}

func (r *invitationRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *invitationRepository) DeleteResponseInvitation(ctx context.Context, tx *sql.Tx, id int) error {
	query := `
		DELETE FROM invitation_response WHERE id = $1;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *invitationRepository) DeleteInvitation(ctx context.Context, tx *sql.Tx, id int) error {
	query := `
		DELETE FROM invitations WHERE id = $1;
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
