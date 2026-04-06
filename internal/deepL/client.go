package deepl

import (
	"net/http"
	"time"

	"babelbridge/internal/config"
)

type DeepLClient struct {
	authKey    string
	baseURL    string
	httpClient *http.Client
}

func NewDeepLClient(cfg config.DeepLConfig) *DeepLClient {
	return &DeepLClient{
		authKey:    cfg.AuthKey,
		baseURL:    cfg.BaseURL,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}
}
