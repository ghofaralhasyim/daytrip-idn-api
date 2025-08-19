package repositories

import (
	"context"
	"database/sql"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type (
	ActivityRepository interface {
		GetActivities(ctx context.Context) ([]entities.ActivityEntity, error)
	}

	activityRepository struct {
		db *sql.DB
	}
)

func NewActivityRepository(db *sql.DB) ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (r *activityRepository) GetActivities(ctx context.Context) (
	[]entities.ActivityEntity, error,
) {
	column := helpers.GenerateSelectColumns[models.Activity](nil)

	query := `SELECT ` + column + " FROM activities;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := helpers.ScanRowsToStructs[models.Activity](rows)
	if err != nil {
		return nil, err
	}

	activities := make([]entities.ActivityEntity, 0)
	for _, item := range results {
		entity := entities.MakeActivityEntity(
			item.Id, item.Name, item.Image,
		)
		activities = append(activities, *entity)
	}

	return activities, nil
}
