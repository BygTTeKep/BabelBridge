package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"babelbridge/internal/translate/domain/entities"
)

type DeepLService struct {
	client *DeepLClient
}

func NewDeepLService(client *DeepLClient) *DeepLService {
	return &DeepLService{client: client}
}

func (dls *DeepLService) Translate(message *entities.MessageEntity) {
	jsonBody, err := json.Marshal(&message)
	if err != nil { // TODO
	}
	contentType := "application/json"
	body := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequest("POST", dls.client.baseURL, body)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", dls.client.authKey)
	if err != nil {
		// TODO
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO
	}
	fmt.Println(string(responseBody))
}
