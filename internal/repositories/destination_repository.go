package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils"
)

type (
	DestinationRepository interface {
		InsertDestination(ctx context.Context, data entities.DestinationEntity) (entities.DestinationEntity, error)
		GetDestinations(ctx context.Context) ([]entities.DestinationEntity, error)
	}

	destinationRepository struct {
		db *sql.DB
	}
)

func NewDestinationRepository(db *sql.DB) DestinationRepository {
	return &destinationRepository{
		db: db,
	}
}

func (r *destinationRepository) InsertDestination(
	ctx context.Context, data entities.DestinationEntity,
) (entities.DestinationEntity, error) {

	model := models.ToDestinationModel(data)
	query := `
		INSERT INTO destinations
			(name, image)
		VALUES
			($1, $2)
		RETURNING id, created_at
	`

	err := r.db.QueryRowContext(ctx, query, model.Name, model.Image).
		Scan(&model.Id, &model.CreatedAt)
	if err != nil {
		return data, err
	}

	data.Id = model.Id
	layout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(layout, model.CreatedAt.String())
	if err != nil {
		return data, err
	}

	data.CreatedAt = parsedTime

	return data, nil
}

func (r *destinationRepository) GetDestinations(ctx context.Context) ([]entities.DestinationEntity, error) {

	column := utils.GenerateSelectColumns[models.Destination](nil)

	query := `SELECT ` + column + " FROM destinations"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := utils.ScanRowsToStructs[models.Destination](rows)
	if err != nil {
		return nil, err
	}

	destinations := make([]entities.DestinationEntity, 0)
	for _, item := range results {
		entity := entities.MakeDestinationEntity(
			item.Id, item.Name, item.Image, item.CreatedAt,
		)
		destinations = append(destinations, *entity)
	}

	return destinations, nil
}
