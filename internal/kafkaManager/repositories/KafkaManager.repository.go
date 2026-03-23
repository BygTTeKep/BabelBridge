package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"babelbridge/internal/database/repositories"
)

type Topic struct {
	// Company    string `json:"company" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Partitions int    `json:"partitions" binding:"required"`
}

type ITopicRepository interface {
	Save(ctx context.Context, t Topic, companyToken string) error
}

type TopicRepository struct {
	db     *sql.DB
	logger *logrus.Entry
}

func NewTopicRepository(db *sql.DB, logger *logrus.Logger) *TopicRepository {
	return &TopicRepository{db: db, logger: logger.WithField("repository", "TopicRepository")}
}

func (tr *TopicRepository) Save(ctx context.Context, t Topic, companyToken string) error {
	var companyID int
	query := fmt.Sprintf(`
		INSERT INTO %s(company_id, name, partitions) VALUES ($1, $2, $3)`,
		repositories.TopicTable,
	)
	getCopanyByIDQuery := fmt.Sprintf(`
		SELECT id FROM %s WHERE token = $1
	`, repositories.CompanyTable)
	err := tr.db.QueryRowContext(ctx, getCopanyByIDQuery, companyToken).Scan(&companyID)
	if err != nil {
		tr.logger.Errorf("error to get company: %v", err)
		return err
	}
	_, err = tr.db.ExecContext(ctx, query, companyID, t.Name, t.Partitions)
	if err != nil {
		tr.logger.Errorf("error to create topic for company: %v", err)
		return err
	}
	return nil
}
