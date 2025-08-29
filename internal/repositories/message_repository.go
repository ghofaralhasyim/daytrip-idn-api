package repositories

import (
	"context"
	"database/sql"

	"github.com/daytrip-idn-api/internal/entities"
	"github.com/daytrip-idn-api/internal/models"
	"github.com/daytrip-idn-api/pkg/utils/helpers"
)

type (
	MessageRepository interface {
		InsertMessage(ctx context.Context, data entities.MessageEntity) (entities.MessageEntity, error)
		GetMessages(ctx context.Context) ([]entities.MessageEntity, error)
		DeleteMessage(ctx context.Context, messageId int64) error
	}

	messageRepository struct {
		db *sql.DB
	}
)

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) InsertMessage(
	ctx context.Context, data entities.MessageEntity,
) (entities.MessageEntity, error) {
	model := models.ToMessageModel(data)

	query := `
		INSERT INTO messages
			(phone, email, package_name, message, name)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id, created_at;
	`

	err := r.db.QueryRowContext(
		ctx, query,
		model.Phone, model.Email, model.PackageName, model.Message, model.Name,
	).Scan(&model.Id, &model.CreatedAt)
	if err != nil {
		return data, err
	}

	data.Id = model.Id
	data.CreatedAt = model.CreatedAt

	return data, nil
}

func (r *messageRepository) GetMessages(
	ctx context.Context,
) ([]entities.MessageEntity, error) {

	column := helpers.GenerateSelectColumns[models.Message](nil)

	query := `SELECT ` + column + " FROM messages;"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	results, err := helpers.ScanRowsToStructs[models.Message](rows)
	if err != nil {
		return nil, err
	}

	messages := make([]entities.MessageEntity, 0)
	for _, item := range results {
		entity := entities.MakeMessageEntity(
			item.Id, item.Phone, item.Email, item.PackageName, item.Message, item.CreatedAt, item.Name,
		)
		messages = append(messages, *entity)
	}

	return messages, nil
}

func (r *messageRepository) DeleteMessage(
	ctx context.Context, messageId int64,
) error {
	query := `
		DELETE FROM messages WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, messageId)
	if err != nil {
		return err
	}

	return nil
}
