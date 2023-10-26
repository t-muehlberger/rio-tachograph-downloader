// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostFilesReader is a Reader for the PostFiles structure.
type PostFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFilesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostFilesCreated creates a PostFilesCreated with default headers values
func NewPostFilesCreated() *PostFilesCreated {
	return &PostFilesCreated{}
}

/*
PostFilesCreated describes a response with status code 201, with default header values.

Success
*/
type PostFilesCreated struct {
}

// IsSuccess returns true when this post files created response has a 2xx status code
func (o *PostFilesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post files created response has a 3xx status code
func (o *PostFilesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post files created response has a 4xx status code
func (o *PostFilesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post files created response has a 5xx status code
func (o *PostFilesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post files created response a status code equal to that given
func (o *PostFilesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostFilesCreated) Error() string {
	return fmt.Sprintf("[POST /files][%d] postFilesCreated ", 201)
}

func (o *PostFilesCreated) String() string {
	return fmt.Sprintf("[POST /files][%d] postFilesCreated ", 201)
}

func (o *PostFilesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
