package kafkamanager

import (
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"babelbridge/internal/kafkaManager/services"
)

type IKafkaManagerServices interface {
	services.ICreateTopicService
}

type kafkaManagerServices struct {
	services.ICreateTopicService
}

func NewKafkaManagerServices(repo IKafkaManagerRepositories, kafka *kafka.Conn, logger *logrus.Logger) *kafkaManagerServices {
	return &kafkaManagerServices{
		ICreateTopicService: services.NewCreateTopicService(repo, kafka, logger),
	}
}
