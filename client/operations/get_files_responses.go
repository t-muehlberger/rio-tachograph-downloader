// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t-muehlberger/rio-tachograph-downloader/models"
)

// GetFilesReader is a Reader for the GetFiles structure.
type GetFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilesOK creates a GetFilesOK with default headers values
func NewGetFilesOK() *GetFilesOK {
	return &GetFilesOK{}
}

/* GetFilesOK describes a response with status code 200, with default header values.

Success
*/
type GetFilesOK struct {
	Payload *models.FileMetadataResponse
}

func (o *GetFilesOK) Error() string {
	return fmt.Sprintf("[GET /files][%d] getFilesOK  %+v", 200, o.Payload)
}
func (o *GetFilesOK) GetPayload() *models.FileMetadataResponse {
	return o.Payload
}

func (o *GetFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
