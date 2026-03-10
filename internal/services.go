package internal

import (
	"babelbridge/internal/company"
	kafkamanager "babelbridge/internal/kafkaManager"
)

type Services struct {
	KafkaManagerService kafkamanager.IKafkaManagerServices
	CompanyService      company.ICompanyServices
}

func NewServices(kms kafkamanager.IKafkaManagerServices, cs company.ICompanyServices) *Services {
	return &Services{
		KafkaManagerService: kms,
		CompanyService:      cs,
	}
}
