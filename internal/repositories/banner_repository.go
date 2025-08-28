package repositories

import (
	"context"
	"database/sql"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type (
	BannerRepository interface {
		InsertBanner(ctx context.Context, data *entities.BannerEntity) (entities.BannerEntity, error)
		GetBanners(ctx context.Context) ([]entities.BannerEntity, error)
		UpdateBanner(ctx context.Context, data *entities.BannerEntity) (entities.BannerEntity, error)
		GetBannerById(ctx context.Context, id int) (*entities.BannerEntity, error)
		DeleteBanner(ctx context.Context, id int) error
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
	ctx context.Context, data *entities.BannerEntity,
) (entities.BannerEntity, error) {

	model := models.ToBannerModel(*data)

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
		return *data, err
	}

	data.Id = model.Id

	return *data, nil
}

func (r *bannerRepository) UpdateBanner(
	ctx context.Context, data *entities.BannerEntity,
) (entities.BannerEntity, error) {

	model := models.ToBannerModel(*data)

	query := `
		UPDATE banners
		SET
			desktop_image = $1,
			mobile_image  = $2,
			cta           = $3,
			cta_url       = $4,
			title         = $5,
			description   = $6
		WHERE id = $7;
	`

	_, err := r.db.ExecContext(ctx,
		query,
		model.DesktopImage,
		model.MobileImage,
		model.Cta,
		model.CtaUrl,
		model.Title,
		model.Description,
		model.Id,
	)

	if err != nil {
		return *data, err
	}

	data.Id = model.Id

	return *data, nil
}

func (r *bannerRepository) GetBanners(ctx context.Context) ([]entities.BannerEntity, error) {

	column := helpers.GenerateSelectColumns[models.Banner](nil)

	query := `SELECT ` + column + " FROM banners ORDER BY id;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := helpers.ScanRowsToStructs[models.Banner](rows)
	if err != nil {
		return nil, err
	}

	banners := make([]entities.BannerEntity, 0)
	for _, item := range results {
		entity := entities.MakeBannerEntity(
			item.Id, item.DesktopImage, item.MobileImage, item.Cta,
			item.CtaUrl, item.Title, item.Description,
		)
		banners = append(banners, *entity)
	}

	return banners, nil
}

func (r *bannerRepository) DeleteBanner(ctx context.Context, id int) error {
	query := `DELETE FROM banners WHERE id = $1;`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *bannerRepository) GetBannerById(ctx context.Context, id int) (*entities.BannerEntity, error) {
	column := helpers.GenerateSelectColumns[models.Banner](nil)

	query := `SELECT ` + column + " FROM banners WHERE id = $1;"

	row := r.db.QueryRowContext(ctx, query, id)

	result, err := helpers.ScanRowToStruct[models.Banner](row, []string{"id", "desktop_image", "mobile_image", "cta", "cta_url", "title", "description"})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	entity := entities.MakeBannerEntity(
		result.Id, result.DesktopImage, result.MobileImage, result.Cta, result.CtaUrl, result.Title, result.Description,
	)

	return entity, nil
}
