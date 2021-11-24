package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	grantType          = "partner_integration"
	scope              = "tachograph-partner.read"
	expirySafetyMargin = 10 * time.Second
)

type authenticator struct {
	tokenUrl      string
	clientID      string
	clientSecret  string
	integrationID string
	httpClient    *http.Client
	mu            sync.Mutex
	token         string
	expiresAt     time.Time
}

func NewAuthenticator(tokenUrl, clientID, clientSecret, integrationID string, httpClient *http.Client) *authenticator {
	return &authenticator{
		tokenUrl:      tokenUrl,
		clientID:      clientID,
		clientSecret:  clientSecret,
		integrationID: integrationID,
		httpClient:    httpClient,
	}
}

func (a *authenticator) GetOrCreateToken() (string, error) {
	if a.checkTokenValid() {
		return a.token, nil
	}

	// optimistic locking to prevent multiple concurrent token creations
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.checkTokenValid() {
		return a.token, nil
	}

	token, expiryTime, err := a.createToken()
	if err != nil {
		return "", err
	}
	a.token = token
	a.expiresAt = expiryTime

	return a.token, nil
}

func (a *authenticator) checkTokenValid() bool {
	if a.token == "" || time.Until(a.expiresAt) < expirySafetyMargin {
		return false
	}
	return true
}

func (a *authenticator) createToken() (string, time.Time, error) {
	data := url.Values{}
	data.Set("client_id", a.clientID)
	data.Set("client_secret", a.clientSecret)
	data.Set("integration_id", a.integrationID)
	data.Set("grant_type", grantType)
	data.Set("scope", scope)

	req, err := http.NewRequest(http.MethodPost, a.tokenUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return "", time.Now(), err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", time.Now(), err
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", time.Now(), err
	}

	expiryTime := time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	return tokenResponse.AccessToken, expiryTime, nil
}
