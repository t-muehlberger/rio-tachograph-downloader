// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFiles(params *GetFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesOK, error)

	GetFilesFileID(params *GetFilesFileIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesFileIDOK, error)

	PostFiles(params *PostFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFilesCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetFiles gets paged file informations or zip of files based on content header results are sorted by file name ascending

  When sending Accept-Header='application/json' API will respond with a description in JSON for the files matching query parameters.
When sending Accept-Header='application/zip' API will respond with a compressed file in ZIP format containing all the tachograph files itself matching request paramters.

# Examples:
## Get all files

`GET /files`

## Get all files between two dates with paging information. `From` and `to` refering to FileMetadataModel::time_created.

`GET /files?from=2018-07-01T08%3A42%3A05.346Z&to=2018-05-28T08%3A42%3A05.346Z&offset=10&limit=10`

## Get Files from driver between two dates. `From` and `to` refering to FileMetadataModel::time_created.

`GET /files?file_type=driver&from=2018-07-01T08%3A42%3A05.346Z&to=2018-05-28T08%3A42%3A05.346Z`

## Get Files relating to specific driver with paging information

`GET /files?offset=0&limit=10&driver_id=7b290aff-6eab-47a3-9b61-e9f6c9dfc906`

*/
func (a *Client) GetFiles(params *GetFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFiles",
		Method:             "GET",
		PathPattern:        "/files",
		ProducesMediaTypes: []string{"application/json", "application/zip"},
		ConsumesMediaTypes: []string{"text/plain; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFilesFileID gets a file itself or file metadata via file id based on content type header

  # Example:

## Get file via file ID 148

`GET /files/148`

*/
func (a *Client) GetFilesFileID(params *GetFilesFileIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFilesFileIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesFileIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFilesFileID",
		Method:             "GET",
		PathPattern:        "/files/{file-id}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"text/plain; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFilesFileIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFilesFileIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFilesFileID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostFiles uploads a file

  # Example:

## Upload a file
`POST /files?fleet_id=d304220d-430a-42fc-939a-b01c50ceef04&file_name=upload.ddd`

Body must contain file in binary format

*/
func (a *Client) PostFiles(params *PostFilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFilesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostFiles",
		Method:             "POST",
		PathPattern:        "/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostFilesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostFiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
