package spotify

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *outbound) GetTokenDetails() (string, string, error) {

	if o.AccessToken == "" || time.Now().After(o.ExpiredAt) {
		if err := o.generateToken(); err != nil {
			return "", "", err
		}
	}

	return o.AccessToken, o.TypeToken, nil
}

func (o *outbound) generateToken() error {
	body := url.Values{}
	body.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, o.cfg.Spotify.TokenURL, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var tokenResponse SpotifyTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return err
	}

	o.AccessToken = tokenResponse.AccessToken
	o.TypeToken = tokenResponse.TokenType
	o.ExpiredAt = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	return nil
}
