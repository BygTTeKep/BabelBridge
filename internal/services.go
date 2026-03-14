package internal

import (
	"babelbridge/internal/company"
	kafkamanager "babelbridge/internal/kafkaManager"
)

type Services struct {
	KafkaManagerService kafkamanager.IKafkaManagerServices
	CompanyService      company.ICompanyService
}

func NewServices(kms kafkamanager.IKafkaManagerServices, cs company.ICompanyService) *Services {
	return &Services{
		KafkaManagerService: kms,
		CompanyService:      cs,
	}
}
