package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type (
	InvitationResponseRepository interface {
		Create(ctx context.Context, response *entities.InvitationResponseEntity) (int64, error)
		GetInvitationResponse(ctx context.Context) ([]entities.InvitationResponseEntity, error)
		GetInvitationResponseBySlug(ctx context.Context, slug string) ([]entities.InvitationResponseEntity, error)
	}

	invitationResponseRepository struct {
		db *sql.DB
	}
)

func NewInvitationResponseRepository(
	db *sql.DB,
) InvitationResponseRepository {
	return &invitationResponseRepository{
		db: db,
	}
}

func (r *invitationResponseRepository) Create(
	ctx context.Context, response *entities.InvitationResponseEntity,
) (int64, error) {
	dataModel := models.ToInvitationResponseModel(*response)

	query := `
		INSERT INTO invitation_response (
			invitation_id, name, is_attending, message, created_at
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int64
	err := r.db.QueryRowContext(ctx,
		query,
		dataModel.InvitationId,
		dataModel.Name,
		dataModel.IsAttending,
		dataModel.Message,
		time.Now(),
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to insert invitation response: %w", err)
	}

	return id, nil
}

func (r *invitationResponseRepository) GetInvitationResponse(ctx context.Context) ([]entities.InvitationResponseEntity, error) {
	column := helpers.GenerateSelectColumns[models.InvitationResponse](nil)

	query := `SELECT ` + column + " FROM invitation_response;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := helpers.ScanRowsToStructs[models.InvitationResponse](rows)
	if err != nil {
		return nil, err
	}

	response := make([]entities.InvitationResponseEntity, 0)
	for _, item := range results {
		entity := entities.MakeInvitationResponseEntity(
			item.Id, item.InvitationId, item.Name, item.IsAttending,
			item.Message, item.CreatedAt,
		)
		response = append(response, *entity)
	}

	return response, nil
}

func (r *invitationResponseRepository) GetInvitationResponseBySlug(ctx context.Context, slug string) ([]entities.InvitationResponseEntity, error) {
	alias := "r"
	column := helpers.GenerateSelectColumns[models.InvitationResponse](&alias)

	query := `SELECT ` + column + " FROM invitation_response r LEFT JOIN invitations i ON i.id = r.invitation_id WHERE i.slug = $1;"

	rows, err := r.db.QueryContext(ctx, query, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	defer rows.Close()

	results, err := helpers.ScanRowsToStructs[models.InvitationResponse](rows)
	if err != nil {
		return nil, err
	}

	response := make([]entities.InvitationResponseEntity, 0)
	for _, item := range results {
		entity := entities.MakeInvitationResponseEntity(
			item.Id, item.InvitationId, item.Name, item.IsAttending,
			item.Message, item.CreatedAt,
		)
		response = append(response, *entity)
	}

	return response, nil
}
