package main

//go:generate go-swagger generate client -f api.yml

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/t-muehlberger/rio-tachograph-downloader/models"
)

func main() {
	rio, cfg, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	var totalCount, downloadedCount int

	files, errs := rio.GetFileMetadata()
	for file := range files {
		totalCount++

		p := getFilePath(cfg.targetDir, file)

		// check if file exists
		_, err := os.Stat(p)
		if err == nil {
			continue
		} else if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("failed to stat file: %v", err)
		}

		// download file
		size, err := rio.DownloadFile(file.FileID, p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s (%d bytes)\n", file.FileName, size)

		// simple rate limiting
		time.Sleep(time.Duration(1000/cfg.filesPerSec) * time.Millisecond)

		downloadedCount++
	}

	for error := range errs {
		log.Fatal(error)
	}

	fmt.Printf("Finished - Downloaded %d files, skipped %d existing files\n", downloadedCount, totalCount-downloadedCount)
}

func setup() (*rio, config, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, cfg, err
	}

	httpClient := &http.Client{}

	authenticator := NewAuthenticator(cfg.toeknUrl, cfg.clientID, cfg.clientSecret, cfg.integrationID, httpClient)

	rio := NewRio(httpClient, authenticator, cfg)

	return rio, cfg, nil
}

func getFilePath(targetDir string, fileMetadata *models.FileMetadataModel) string {
	timeCreated := time.Time(fileMetadata.TimeCreated)
	return path.Join(
		targetDir,
		fileMetadata.FileType,
		strconv.Itoa(timeCreated.Local().Year()),
		fmt.Sprintf("%02d", timeCreated.Local().Month()),
		fileMetadata.FileName)
}
