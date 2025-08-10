package repositories

import (
	"context"
	"database/sql"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils"
)

type (
	BannerRepository interface {
		InsertBanner(ctx context.Context, data entities.BannerEntity) (entities.BannerEntity, error)
		GetBanners(ctx context.Context) ([]entities.BannerEntity, error)
	}

	bannerRepository struct {
		db *sql.DB
	}
)

func NewBannerRepository(db *sql.DB) BannerRepository {
	return &bannerRepository{
		db: db,
	}
}

func (r *bannerRepository) InsertBanner(
	ctx context.Context, data entities.BannerEntity,
) (entities.BannerEntity, error) {

	model := models.ToBannerModel(data)

	query := `
		INSERT INTO banners
			(
				desktop_image, mobile_image, cta,
				cta_url, title, description
			)
		VALUES
			(
				$1, $2, $3,
				$4, $5, $6
			)
		RETURNING id;
	`

	err := r.db.QueryRowContext(ctx,
		query, model.DesktopImage, model.MobileImage, model.Cta,
		model.CtaUrl, model.Title, model.Description,
	).Scan(&model.Id)
	if err != nil {
		return data, err
	}

	data.Id = model.Id

	return data, nil
}

func (r *bannerRepository) GetBanners(ctx context.Context) ([]entities.BannerEntity, error) {

	column := utils.GenerateSelectColumns[models.Banner](nil)

	query := `SELECT ` + column + " FROM banners;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := utils.ScanRowsToStructs[models.Banner](rows)
	if err != nil {
		return nil, err
	}

	banners := make([]entities.BannerEntity, 0)
	for _, item := range results {
		entity := entities.MakeBannerEntity(
			item.Id, item.DesktopImage, item.MobileImage, item.Cta,
			item.Title, item.CtaUrl, item.Description,
		)
		banners = append(banners, *entity)
	}

	return banners, nil
}
