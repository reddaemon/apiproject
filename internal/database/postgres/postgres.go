package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type PsqlRepository struct {
	*sqlx.DB //nolint
	logger   *zap.Logger
}

func NewPsqlRepository(DB *sqlx.DB, logger *zap.Logger) PsqlRepository { // nolint
	return PsqlRepository{DB: DB, logger: logger}
}

// Insert to DB method
func (p PsqlRepository) InsertToDB(ctx context.Context,
	title string, link string, published string) error {
	sugar := p.logger.Sugar()
	query := `INSERT INTO rss (title,link,date) VALUES ($1,$2,$3)`
	_, err := p.DB.ExecContext(ctx, query, title, link, published)
	if err != nil {
		sugar.Errorf("Cannot insert query %v", err)
		return err
	}
	return nil
}
