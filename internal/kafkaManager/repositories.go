package kafkamanager

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	"babelbridge/internal/kafkamanager/repositories"
)

type IKafkaManagerRepositories interface {
	repositories.ITopicRepository
}

type KafkaManagerRepositories struct {
	repositories.ITopicRepository
}

func NewKafkaManagerRepositories(db *sql.DB, logger *logrus.Logger) *KafkaManagerRepositories {
	return &KafkaManagerRepositories{
		ITopicRepository: repositories.NewTopicRepository(db, logger),
	}
}
