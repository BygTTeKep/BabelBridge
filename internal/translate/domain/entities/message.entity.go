package entities

// MessageEntity сущность для перевода
//
// messageID - id сообщения в мессенджере
//
// message - само сообщенние
//
// translateFrom - с какого языка переводить
//
// translateTo - на какой язык переводить
//
// topic - топик куда писать ответ
type MessageEntity struct {
	messageID     string `validate:"required"`
	message       string `validate:"required"`
	translateFrom string `validate:"required"`
	translateTo   string `validate:"required"`
	topic         string `validate:"required"`
}
