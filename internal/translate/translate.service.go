package translate

type ITranslateService interface {
	Translate()
}

type TranslateService struct {
	repo *TranslateRepository
}

func NewTranslateService(repo *TranslateRepository) *TranslateService {
	return &TranslateService{
		repo: repo,
	}
}

func (ts *TranslateService) Translate() {
	// http бросок
}
