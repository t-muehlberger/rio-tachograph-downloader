package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/t-muehlberger/rio-tachograph-downloader/client"
	"github.com/t-muehlberger/rio-tachograph-downloader/client/operations"
	"github.com/t-muehlberger/rio-tachograph-downloader/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	metadataPageSize = 100
)

type rio struct {
	httpClient    *http.Client
	authenticator *authenticator
	apiClient     *client.TachographFileArchive
	config        config
}

func NewRio(httpClient *http.Client, auth *authenticator, cfg config) *rio {
	r := &rio{
		httpClient:    httpClient,
		authenticator: auth,
		config:        cfg,
	}

	transport := httptransport.NewWithClient(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes, httpClient)

	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		token, err := auth.GetOrCreateToken()
		if err != nil {
			return fmt.Errorf("failed to get token for request: %v", err)
		}
		r.SetHeaderParam("Authorization", "Bearer "+token)
		return nil
	})

	r.apiClient = client.New(transport, strfmt.Default)

	return r
}

func (r *rio) GetFileMetadata() (<-chan *models.FileMetadataModel, <-chan error) {
	fileInfos := make(chan *models.FileMetadataModel)
	error := make(chan error, 1)

	go func() {
		defer close(fileInfos)
		defer close(error)

		var offset int32 = 0
		var limit int32 = metadataPageSize

		hasMore := true
		for hasMore {
			params := operations.NewGetFilesParams()
			params.SetOffset(&offset)
			params.SetLimit(&limit)

			resp, err := r.apiClient.Operations.GetFiles(params, nil, func(op *runtime.ClientOperation) {
				op.ConsumesMediaTypes = []string{"application/json"}
			})
			if err != nil {
				error <- err
				return
			}
			files := resp.Payload

			for _, file := range files.Items {
				fileInfos <- file
			}

			hasMore = int(files.Pagination.Offset)+len(files.Items) < int(files.TotalCount)
			offset += int32(len(files.Items))
		}
	}()

	return fileInfos, error
}

func (r *rio) DownloadFile(id int32, filePath string) (int64, error) {
	url := fmt.Sprintf("%s/files/%d", r.config.apiBaseUrl, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}

	token, err := r.authenticator.GetOrCreateToken()
	if err != nil {
		return 0, fmt.Errorf("failed to get token: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/octet-stream")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to get file: %v", err)
	}
	defer resp.Body.Close()

	dir := path.Dir(filePath)
	os.MkdirAll(dir, 0755)

	f, err := os.Create(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	size, err := io.Copy(f, resp.Body)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}

	return size, nil
}
