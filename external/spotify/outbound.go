package spotify

import (
	"catalog-music/internal/configs"
	"catalog-music/pkg/httpclient"
	"time"
)

type outbound struct {
	cfg         *configs.Config
	client      httpclient.HTTPClient
	AccessToken string
	TypeToken   string
	ExpiredAt   time.Time
}

func NewOutbound(cfg *configs.Config, client httpclient.HTTPClient) *outbound {
	return &outbound{
		cfg:    cfg,
		client: client,
	}
}
