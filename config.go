package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/t-muehlberger/rio-tachograph-downloader/client"
)

const (
	defaultTokenUrl    = "https://auth.iam.rio.cloud/oauth/token"
	defaultTargetDir   = "."
	defaultFilesPerSec = 4
)

const (
	envApiBaseUrl    = "API_BASE_URL"
	envTokenUrl      = "API_TOKEN_URL"
	envClientID      = "API_CLIENT_ID"
	envClientSecret  = "API_CLIENT_SECRET"
	envIntegrationID = "API_INTEGRATION_ID"
	envTargetDir     = "TARGET_DIR"
	envFilesPerSec   = "FILES_PER_SEC"
)

type config struct {
	apiBaseUrl    string
	toeknUrl      string
	clientID      string
	clientSecret  string
	integrationID string
	targetDir     string
	filesPerSec   int
}

func getConfig() (config, error) {
	c := config{
		apiBaseUrl:    os.Getenv(envApiBaseUrl),
		toeknUrl:      os.Getenv(envTokenUrl),
		clientID:      os.Getenv(envClientID),
		clientSecret:  os.Getenv(envClientSecret),
		integrationID: os.Getenv(envIntegrationID),
		targetDir:     os.Getenv(envTargetDir),
	}

	filesPerSecStr := os.Getenv(envFilesPerSec)
	if filesPerSecStr == "" {
		c.filesPerSec = defaultFilesPerSec
	} else {
		filesPerSec, err := strconv.Atoi(filesPerSecStr)
		if err != nil {
			return c, err
		}
		c.filesPerSec = filesPerSec
	}

	if c.apiBaseUrl == "" {
		c.apiBaseUrl = client.DefaultSchemes[0] + "://" + client.DefaultHost + client.DefaultBasePath
	}
	if c.toeknUrl == "" {
		c.toeknUrl = defaultTokenUrl
	}
	if c.targetDir == "" {
		c.targetDir = defaultTargetDir
	}

	if c.clientID == "" {
		return c, fmt.Errorf("variable '%s' is not set", envClientID)
	}
	if c.clientSecret == "" {
		return c, fmt.Errorf("variable '%s' is not set", envClientSecret)
	}
	if c.integrationID == "" {
		return c, fmt.Errorf("variable '%s' is not set", envIntegrationID)
	}

	return c, nil
}
