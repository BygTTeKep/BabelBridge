package company

import "babelbridge/internal/company/services"

type ICompanyServices interface {
	services.ICreateCompanyService
}

type CompanyServices struct {
	services.ICreateCompanyService
}

func NewCompanyService(repo ICompanyRepository) *CompanyServices {
	return &CompanyServices{
		ICreateCompanyService: services.NewCreateCompanyService(repo),
	}
}
