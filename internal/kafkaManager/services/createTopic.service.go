package services

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"babelbridge/internal/kafkamanager/repositories"
)

type ICreateTopicService interface {
	CreateTopic(ctx context.Context, topic repositories.Topic, token string) error
}

type CreateTopicService struct {
	repo       repositories.ITopicRepository
	kafkaAdmin *kafka.Conn
	logger     *logrus.Entry
}

func NewCreateTopicService(
	repo repositories.ITopicRepository,
	kafkaAdmin *kafka.Conn,
	logger *logrus.Logger,
) *CreateTopicService {
	return &CreateTopicService{
		repo:       repo,
		kafkaAdmin: kafkaAdmin,
		logger:     logger.WithField("service", "CreateTopicService"),
	}
}

func (cts *CreateTopicService) CreateTopic(ctx context.Context, topic repositories.Topic, token string) error {
	// topicConfig := kafka.TopicConfig{Topic: topic.Name, NumPartitions: topic.Partitions, ReplicationFactor: 1}
	// err := cts.kafkaAdmin.CreateTopics(topicConfig)
	// if err != nil {
	// 	cts.logger.Errorf("error to create topic for company %s:%v", "", err)
	// 	return err
	// }
	err := cts.repo.Save(ctx, topic, token)

	return err
}
