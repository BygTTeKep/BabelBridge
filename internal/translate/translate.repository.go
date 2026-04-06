package translate

type ITranslateRepository interface {
	Translate()
}

type TranslateRepository struct{}

func NewTranslateRepository() *TranslateRepository {
	return &TranslateRepository{}
}

func (tr *TranslateRepository) Translate() {}
